package persister

import (
	"database/sql"
	//	"github.com/jmoiron/sqlx"
	"go_angular_starter/model/entity"
	"log"
	"reflect"
)

type UserPersister struct {
	db     DB
	common *commonPersister
}

func NewUserPersister(db DB) *UserPersister {
	entityType := reflect.TypeOf(&entity.User{}).Elem()
	return &UserPersister{db, &commonPersister{db, entityType}}
}

func (self UserPersister) GetAll() ([]*entity.User, error) {
	users := []*entity.User{}

	err := self.common.getAll(&users, "")

	return users, err
}

func (self UserPersister) GetById(id int64) (*entity.User, error) {
	user := &entity.User{}

	err := self.common.getById(user, id)

	return user, err
}

func (self UserPersister) Insert(user *entity.User) error {
	return self.common.insert(user)
}

func (self UserPersister) Update(user *entity.User) error {
	return self.common.update(user)
}

func (self UserPersister) Delete(id int64) error {
	return self.common.delete(id)
}

func (self UserPersister) GetByEmailAndHash(email string, hash string) (*entity.User, error) {
	log.Printf("email %s, hash %s", email, hash)

	query := "SELECT * FROM users WHERE email = $1 AND hash = $2"

	user := &entity.User{}
	err := self.db.Get(user, query, email, hash)

	if err != nil {
		return nil, err
	}

	return user, nil
}

func (self UserPersister) GetByApiKey(apiKey string) (*entity.User, error) {
	query := "SELECT u.* FROM users AS u INNER JOIN sessions AS s ON u.id = s.user_id WHERE s.api_key = $1"

	user := &entity.User{}
	err := self.db.Get(user, query, apiKey)

	if err == sql.ErrNoRows {
		return nil, nil
	}

	if err != nil {
		return nil, err
	}

	return user, nil
}
