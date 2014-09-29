package persister

import (
	"github.com/cjdell/go_angular_starter/model/entity"
	"reflect"
)

type CategoryPersister struct {
	db     DB
	common *commonPersister
}

func NewCategoryPersister(db DB) *CategoryPersister {
	entityType := reflect.TypeOf(&entity.Category{}).Elem()
	return &CategoryPersister{db, &commonPersister{db, entityType}}
}

func (self CategoryPersister) GetAll() ([]*entity.Category, error) {
	categories := []*entity.Category{}

	err := self.common.getAll(&categories, "")

	return categories, err
}

func (self CategoryPersister) GetChildren(parentId int64) ([]*entity.Category, error) {
	categories := []*entity.Category{}

	err := self.common.getAll(&categories, "parent_id = $1", parentId)

	return categories, err
}

func (self CategoryPersister) GetById(id int64) (*entity.Category, error) {
	category := &entity.Category{}

	err := self.common.getById(category, id)

	return category, err
}

func (self CategoryPersister) Insert(category *entity.Category) error {
	return self.common.insert(category)
}

func (self CategoryPersister) Update(category *entity.Category) error {
	return self.common.update(category)
}

func (self CategoryPersister) Delete(id int64) error {
	return self.common.delete(id)
}
