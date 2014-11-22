package entity

type User struct {
	__table struct{} `db:"users"`

	Id    int64  `db:"id"`
	Name  string `db:"name"`
	Email string `db:"email"`
	Hash  string `db:"hash" json:"-"`
}

func (self User) GetId() int64 {
	return self.Id
}

func (self *User) Merge(update *User, fields []string) error {
	self.Name = update.Name
	self.Email = update.Email

	return nil
}
