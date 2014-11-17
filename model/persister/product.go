package persister

import (
	"github.com/cjdell/go_angular_starter/model/entity"
	"reflect"
	"strings"
)

type ProductPersister struct {
	db     DB
	common *commonPersister
}

type ProductFilter struct {
	CategoryId *int64
}

func NewProductPersister(db DB) *ProductPersister {
	entityType := reflect.TypeOf(&entity.Product{}).Elem()
	return &ProductPersister{db, &commonPersister{db, entityType}}
}

func (self ProductPersister) GetAll(limit *Limit) ([]*entity.Product, error) {
	products := []*entity.Product{}
	return products, self.common.getAll(&products, limit, "", nil)
}

func (self ProductPersister) GetByFilter(productFilter ProductFilter, limit *Limit) ([]*entity.Product, error) {
	products := []*entity.Product{}

	where := []string{"1 = 1"}
	params := make(QueryParameters)

	if productFilter.CategoryId != nil {
		where = append(where, "category_ids @> array[(:category_id)::::bigint]")
		params["category_id"] = productFilter.CategoryId
	}

	return products, self.common.getAll(&products, limit, "WHERE "+strings.Join(where, " AND "), params)
}

func (self ProductPersister) GetById(id int64) (*entity.Product, error) {
	product := &entity.Product{}

	return product, self.common.getOne(product, "WHERE id = :id", NewQueryParametersWithId(id))
}

func (self ProductPersister) Insert(product *entity.Product) (int64, error) {
	return self.common.insert(product)
}

func (self ProductPersister) Update(product *entity.Product) error {
	return self.common.update(product)
}

func (self ProductPersister) Delete(id int64) error {
	return self.common.delete(id)
}

func (self ProductPersister) GetByHandle(handle string) (*entity.Product, error) {
	product := &entity.Product{}

	params := QueryParameters{}
	params["handle"] = handle

	return product, self.common.getOne(product, "WHERE handle = :handle", params)
}
