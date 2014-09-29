package entity

type Category struct {
	__table struct{} `db:"categories"`

	Id       int64  `db:"id"`
	ParentId int64  `db:"parent_id"`
	Name     string `db:"name"`
	FqName   string `db:"fq_name"`
}

func (Category) GetTypeName() string {
	return "Category"
}

func (self *Category) GetId() int64 {
	return self.Id
}

func (self *Category) SetId(id int64) {
	self.Id = id
}
