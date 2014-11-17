package persister

import (
	"github.com/cjdell/go_angular_starter/model/entity"
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

func (self UserPersister) GetAll(limit *Limit) ([]*entity.User, error) {
	users := []*entity.User{}
	return users, self.common.getAll(&users, limit, "", nil)
}

func (self UserPersister) GetById(id int64) (*entity.User, error) {
	user := &entity.User{}
	return user, self.common.getOne(user, "WHERE id = :id", NewQueryParametersWithId(id))
}

func (self UserPersister) Insert(user *entity.User) (int64, error) {
	return self.common.insert(user)
}

func (self UserPersister) Update(user *entity.User) error {
	return self.common.update(user)
}

func (self UserPersister) Delete(id int64) error {
	return self.common.delete(id)
}

func (self UserPersister) GetByEmailAndHash(email string, hash string) (*entity.User, error) {
	user := &entity.User{}

	params := QueryParameters{}
	params["email"] = email
	params["hash"] = hash

	return user, self.common.getOne(user, "WHERE email = :email AND hash = :hash", params)
}

func (self UserPersister) GetByApiKey(apiKey string) (*entity.User, error) {
	user := &entity.User{}

	params := QueryParameters{}
	params["api_key"] = apiKey

	return user, self.common.getOne(user, "INNER JOIN sessions AS s ON e.id = s.user_id WHERE s.api_key = :api_key", params)
}
