package services

import (
	"github.com/cjdell/go_angular_starter/model/entity"
	"github.com/cjdell/go_angular_starter/model/persister"
	"github.com/kennygrant/sanitize"
)

type CategoryService struct {
	persister *persister.CategoryPersister
	user      *entity.User
}

// A struct composed of the record coupled with additional computed information
type CategoryInfo struct {
	*entity.Category

	Parent *entity.Category
}

// A struct to hold additional modifications that don't fit into the record data structure
type CategorySave struct {
	*entity.Category

	Changes struct {
		Fields []string
	}
}

func NewCategoryService(db persister.DB, user *entity.User) *CategoryService {
	return &CategoryService{persister.NewCategoryPersister(db), user}
}

func (self *CategoryService) GetAll(parentId *int64, limit *persister.Limit) ([]*CategoryInfo, error) {
	var err error
	var categories []*entity.Category

	if categories, err = self.persister.GetAll(parentId, limit); err != nil {
		return nil, err
	}

	categoryInfos := make([]*CategoryInfo, len(categories), len(categories))

	for i, category := range categories {
		categoryInfos[i], _ = self.categoryInfo(category)
	}

	return categoryInfos, nil
}

func (self *CategoryService) GetOne(id int64) (*CategoryInfo, error) {
	var err error
	var category *entity.Category

	if category, err = self.persister.GetById(id); err != nil {
		return nil, err
	}

	return self.categoryInfo(category)
}

func (self *CategoryService) Insert(categorySave *CategorySave) (*CategoryInfo, error) {
	var err error

	category := categorySave.Category

	if err = self.beforeSave(category, categorySave); err != nil {
		return nil, err
	}

	if category.Id, err = self.persister.Insert(category); err != nil {
		return nil, err
	}

	if err = self.afterSave(category, categorySave); err != nil {
		return nil, err
	}

	return self.GetOne(category.Id)
}

func (self *CategoryService) Update(categorySave *CategorySave) (*CategoryInfo, error) {
	var err error
	var category *entity.Category

	if category, err = self.persister.GetById(categorySave.Id); err != nil {
		return nil, err
	}

	if err = self.beforeSave(category, categorySave); err != nil {
		return nil, err
	}

	if err = self.persister.Update(category); err != nil {
		return nil, err
	}

	if err = self.afterSave(category, categorySave); err != nil {
		return nil, err
	}

	return self.GetOne(category.Id)
}

func (self *CategoryService) Delete(id int64) error {
	return self.persister.Delete(id)
}

// Wrap Category into CategoryInfo - Add computed properties here
func (self *CategoryService) categoryInfo(category *entity.Category) (*CategoryInfo, error) {
	var parent *entity.Category = nil

	if category.ParentId != 0 {
		parent, _ = self.persister.GetById(category.ParentId)
	}

	return &CategoryInfo{category, parent}, nil
}

// Handle CategorySave - before saving to the database
func (self *CategoryService) beforeSave(category *entity.Category, categorySave *CategorySave) error {
	return category.Merge(categorySave.Category, categorySave.Changes.Fields)
}

// Handle CategorySave - after saving to the database
func (self *CategoryService) afterSave(category *entity.Category, categorySave *CategorySave) error {
	return self.generateFullyQualifiedNames(categorySave.Id)
}

func (self *CategoryService) generateFullyQualifiedNames(categoryId int64) error {
	category, err := self.persister.GetById(categoryId)

	if err != nil {
		return err
	}

	var parent *entity.Category = nil

	if category.ParentId != 0 {
		parent, err = self.persister.GetById(category.ParentId)
	}

	return self.generateFullyQualifiedNamesRecursive(category, parent)
}

func (self *CategoryService) generateFullyQualifiedNamesRecursive(category *entity.Category, parent *entity.Category) error {
	if parent != nil {
		category.FqName = parent.FqName + " > " + category.Name
		category.Handle = parent.Handle + "/" + generateHandle(category.Name)
	} else {
		category.FqName = category.Name
		category.Handle = generateHandle(category.Name)
	}

	self.persister.Update(category)

	children, err := self.persister.GetAll(&category.Id, nil)

	if err != nil {
		return err
	}

	for _, child := range children {
		err = self.generateFullyQualifiedNamesRecursive(child, category)

		if err != nil {
			return err
		}
	}

	return nil
}

func generateHandle(name string) string {
	return sanitize.Path(name)
}
