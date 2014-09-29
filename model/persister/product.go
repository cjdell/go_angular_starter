package persister

import (
	"github.com/jmoiron/sqlx"
	"github.com/cjdell/go_angular_starter/model/entity"
	"reflect"
)

type ProductPersister struct {
	db     *sqlx.DB
	common *commonPersister
}

func NewProductPersister(db *sqlx.DB) *ProductPersister {
	entityType := reflect.TypeOf(&entity.Product{}).Elem()
	return &ProductPersister{db, &commonPersister{db, entityType}}
}

func (self ProductPersister) GetAll() ([]*entity.Product, error) {
	products := []*entity.Product{}

	err := self.common.getAll(&products, "")

	return products, err
}

func (self ProductPersister) GetById(id int64) (*entity.Product, error) {
	product := &entity.Product{}

	err := self.common.getById(product, id)

	return product, err
}

func (self ProductPersister) Insert(product *entity.Product) error {
	return self.common.insert(product)
}

func (self ProductPersister) Update(product *entity.Product) error {
	return self.common.update(product)
}

func (self ProductPersister) Delete(id int64) error {
	return self.common.delete(id)
}
