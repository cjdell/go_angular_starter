package services

import (
	"github.com/cjdell/go_angular_starter/model/entity"
	"github.com/cjdell/go_angular_starter/model/persister"
)

type UserService struct {
	persister *persister.UserPersister
	user      *entity.User
}

// A struct composed of the record coupled with additional computed information
type UserInfo struct {
	*entity.User
}

// A struct to hold additional modifications that don't fit into the record data structure
type UserChanges struct {
	*entity.User

	Changes struct {
		Fields []string
	}
}

func NewUserService(db persister.DB, user *entity.User) *UserService {
	return &UserService{persister.NewUserPersister(db), user}
}

func (self *UserService) GetAll(limit *persister.Limit) ([]*UserInfo, error) {
	var err error
	var users []*entity.User

	if users, err = self.persister.GetAll(limit); err != nil {
		return nil, err
	}

	userInfos := make([]*UserInfo, len(users), len(users))

	for i, user := range users {
		userInfos[i], _ = self.userInfo(user)
	}

	return userInfos, nil
}

func (self *UserService) GetOne(id int64) (*UserInfo, error) {
	var err error
	var user *entity.User

	if user, err = self.persister.GetById(id); err != nil {
		return nil, err
	}

	return self.userInfo(user)
}

func (self *UserService) Insert(userChanges *UserChanges) (*UserInfo, error) {
	var err error
	var id int64

	if id, err = self.persister.Insert(userChanges.User); err != nil {
		return nil, err
	}

	userChanges.Id = id

	if err = self.processChanges(userChanges); err != nil {
		return nil, err
	}

	return self.GetOne(id)
}

func (self *UserService) Update(userChanges *UserChanges) (*UserInfo, error) {
	var err error
	var user *entity.User

	if user, err = self.persister.GetById(userChanges.Id); err != nil {
		return nil, err
	}

	if err = user.Merge(userChanges.User, userChanges.Changes.Fields); err != nil {
		return nil, err
	}

	if err = self.persister.Update(user); err != nil {
		return nil, err
	}

	if err = self.processChanges(userChanges); err != nil {
		return nil, err
	}

	return self.GetOne(user.Id)
}

func (self *UserService) Delete(id int64) error {
	return self.persister.Delete(id)
}

// Wrap User into UserInfo - Add computed properties here
func (self *UserService) userInfo(user *entity.User) (*UserInfo, error) {
	return &UserInfo{user}, nil
}

// Process additional mutations that might exist i.e. an uploaded file
func (self *UserService) processChanges(userChanges *UserChanges) error {
	return nil
}
