package services

import (
	"github.com/cjdell/go_angular_starter/model/entity"
	"github.com/cjdell/go_angular_starter/model/persister"
)

type ProductService struct {
	db        persister.DB
	persister *persister.ProductPersister
	user      *entity.User
}

// A struct composed of the record coupled with additional computed information
type ProductInfo struct {
	*entity.Product

	Images     []*Image
	Categories []*entity.Category
}

// A struct to hold additional modifications that don't fit into the record data structure
type ProductSave struct {
	*entity.Product

	Changes struct {
		Fields []string

		NewImageFileName string
		NewImageHandle   string

		ImageChanges map[string]struct {
			Desc string
		}
	}
}

func NewProductService(db persister.DB, user *entity.User) *ProductService {
	return &ProductService{db, persister.NewProductPersister(db), user}
}

func (self *ProductService) GetAll(categoryId *int64, limit *persister.Limit) ([]*ProductInfo, error) {
	var err error
	var products []*entity.Product

	filter := persister.ProductFilter{}

	filter.CategoryId = categoryId

	if products, err = self.persister.GetByFilter(filter, limit); err != nil {
		return nil, err
	}

	productInfos := make([]*ProductInfo, len(products), len(products))

	for i, product := range products {
		productInfos[i], _ = self.productInfo(product)
	}

	return productInfos, nil
}

func (self *ProductService) GetOne(id int64) (*ProductInfo, error) {
	var err error
	var product *entity.Product

	if product, err = self.persister.GetById(id); err != nil {
		return nil, err
	}

	return self.productInfo(product)
}

func (self *ProductService) GetByHandle(handle string) (*ProductInfo, error) {
	var err error
	var product *entity.Product

	if product, err = self.persister.GetByHandle(handle); err != nil {
		return nil, err
	}

	return self.productInfo(product)
}

func (self *ProductService) Insert(productSave *ProductSave) (*ProductInfo, error) {
	var err error

	product := productSave.Product

	if err = self.beforeSave(product, productSave); err != nil {
		return nil, err
	}

	if product.Id, err = self.persister.Insert(product); err != nil {
		return nil, err
	}

	if err = self.afterSave(product, productSave); err != nil {
		return nil, err
	}

	return self.GetOne(product.Id)
}

func (self *ProductService) Update(productSave *ProductSave) (*ProductInfo, error) {
	var err error
	var product *entity.Product

	if product, err = self.persister.GetById(productSave.Id); err != nil {
		return nil, err
	}

	if err = self.beforeSave(product, productSave); err != nil {
		return nil, err
	}

	if err = self.persister.Update(product); err != nil {
		return nil, err
	}

	if err = self.afterSave(product, productSave); err != nil {
		return nil, err
	}

	return self.GetOne(product.Id)
}

func (self *ProductService) Delete(id int64) error {
	return self.persister.Delete(id)
}

// Wrap Product into ProductInfo - Add computed properties here
func (self *ProductService) productInfo(product *entity.Product) (*ProductInfo, error) {
	images, err := GetImages(product)

	if err != nil {
		return nil, err
	}

	categoryPersister := persister.NewCategoryPersister(self.db)

	categories := make([]*entity.Category, len(product.GetCategoryIds()))

	for i, cid := range product.GetCategoryIds() {
		categories[i], _ = categoryPersister.GetById(cid)
	}

	return &ProductInfo{product, images, categories}, nil
}

// Handle ProductSave - before saving to the database
func (self *ProductService) beforeSave(product *entity.Product, productSave *ProductSave) error {
	if err := product.Merge(productSave.Product, productSave.Changes.Fields); err != nil {
		return err
	}

	product.GenerateHandle()

	return nil
}

// Handle ProductSave - after saving to the database
func (self *ProductService) afterSave(product *entity.Product, productSave *ProductSave) error {
	if productSave.Changes.NewImageFileName != "" {
		AssignImage(product, productSave.Changes.NewImageFileName, productSave.Changes.NewImageHandle, "")
	}

	if productSave.Changes.ImageChanges != nil {
		for handle, changes := range productSave.Changes.ImageChanges {
			err := SetImageDescription(product, handle, changes.Desc)

			if err != nil {
				return err
			}
		}
	}

	return nil
}

type NotOwnedError struct {
}

func (NotOwnedError) Error() string { return "You do not own this item" }
