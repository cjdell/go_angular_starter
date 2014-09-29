package entity

import (
	//"log"
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
}

func (Product) GetTypeName() string {
	return "Product"
}

func (self *Product) GetId() int64 {
	return self.Id
}

func (self *Product) SetId(id int64) {
	self.Id = id
}

func (self *Product) Update(update *Product) {
	self.Name = update.Name
	self.Description = update.Description
	self.Price = update.Price
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

//func (self *Product) SetCategoryIds(ids []int64) {
//	idsStr := make([]string, len(ids))

//	for i, id := range ids {
//		idsStr[i] = strconv.FormatInt(id, 10)
//	}

//	self.CategoryIds = "{" + strings.Join(idsStr, ",") + "}"

//	//log.Println(self.CategoryIds)
//}
