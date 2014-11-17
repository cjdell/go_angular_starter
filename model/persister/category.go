package persister

import (
	"github.com/cjdell/go_angular_starter/model/entity"
	"reflect"
	"strings"
)

type CategoryPersister struct {
	db     DB
	common *commonPersister
}

func NewCategoryPersister(db DB) *CategoryPersister {
	entityType := reflect.TypeOf(&entity.Category{}).Elem()
	return &CategoryPersister{db, &commonPersister{db, entityType}}
}

func (self CategoryPersister) GetAll(parentId *int64, limit *Limit) ([]*entity.Category, error) {
	categories := []*entity.Category{}

	where := []string{"1 = 1"}
	params := make(QueryParameters)

	if parentId != nil {
		where = append(where, "parent_id = :parent_id")
		params["parent_id"] = parentId
	}

	return categories, self.common.getAll(&categories, limit, "WHERE "+strings.Join(where, " AND "), params)
}

func (self CategoryPersister) GetById(id int64) (*entity.Category, error) {
	category := &entity.Category{}
	return category, self.common.getOne(category, "WHERE id = :id", NewQueryParametersWithId(id))
}

func (self CategoryPersister) Insert(category *entity.Category) (int64, error) {
	return self.common.insert(category)
}

func (self CategoryPersister) Update(category *entity.Category) error {
	return self.common.update(category)
}

func (self CategoryPersister) Delete(id int64) error {
	return self.common.delete(id)
}

func (self CategoryPersister) GetByHandle(handle string) (*entity.Category, error) {
	category := &entity.Category{}

	params := QueryParameters{}
	params["handle"] = handle

	return category, self.common.getOne(category, "WHERE handle = :handle", params)
}
