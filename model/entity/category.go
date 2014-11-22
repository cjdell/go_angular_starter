package entity

type Category struct {
	__table struct{} `db:"categories"`

	Id       int64  `db:"id"`
	ParentId int64  `db:"parent_id"`
	Name     string `db:"name"`
	FqName   string `db:"fq_name"`
	Handle   string `db:"handle"`
}

func (Category) GetTypeName() string {
	return "Category"
}

func (self *Category) GetId() int64 {
	return self.Id
}

func (self *Category) Merge(update *Category, fields []string) error {
	self.Name = update.Name

	return nil
}
