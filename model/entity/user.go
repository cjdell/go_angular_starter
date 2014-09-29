package entity

type User struct {
	__table struct{} `db:"users"`

	Id    int64  `db:"id"`
	Name  string `db:"name"`
	Email string `db:"email"`
	Hash  string `db:"hash"`
}

func (self User) GetId() int64 {
	return self.Id
}

func (self *User) SetId(id int64) {
	self.Id = id
}
