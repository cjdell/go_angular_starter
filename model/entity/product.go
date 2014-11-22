package entity

import (
	"github.com/kennygrant/sanitize"
	"strconv"
	"strings"
)

type Product struct {
	__table struct{} `db:"products"`

	Id          int64   `db:"id"`
	CategoryIds string  `db:"category_ids"`
	Name        string  `db:"name"`
	Description string  `db:"description"`
	Price       float64 `db:"price"`
	Handle      string  `db:"handle"`
}

func (Product) GetTypeName() string {
	return "Product"
}

func (self *Product) GetId() int64 {
	return self.Id
}

func (self *Product) Merge(update *Product, fields []string) error {
	self.CategoryIds = update.CategoryIds
	self.Name = update.Name
	self.Description = update.Description
	self.Price = update.Price
	self.Handle = update.Handle

	return nil
}

func (self *Product) GetCategoryIds() []int64 {
	commas := strings.Trim(self.CategoryIds, "{}")

	if len(commas) == 0 {
		return make([]int64, 0)
	}

	idsStr := strings.Split(commas, ",")
	ids := make([]int64, len(idsStr))

	for i, str := range idsStr {
		ids[i], _ = strconv.ParseInt(str, 10, 64)
	}

	return ids
}

func (self *Product) GenerateHandle() {
	self.Handle = sanitize.Path(self.Name)
}
