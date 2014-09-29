package api

import (
	"errors"
	"github.com/jmoiron/sqlx"
	"go_angular_starter/model/entity"
	"go_angular_starter/model/persister"
	"net/http"
)

type UserApi struct {
	persister *persister.UserPersister
}

func NewUserApi(db *sqlx.DB) *UserApi {
	return &UserApi{persister.NewUserPersister(db)}
}

type UserGetAllArgs struct {
}

type UserGetAllReply struct {
	Users []*entity.User
}

func (self *UserApi) GetAll(r *http.Request, args *UserGetAllArgs, reply *UserGetAllReply) error {
	users, err := self.persister.GetAll()

	if err != nil {
		return err
	}

	reply.Users = users

	return nil
}

type UserGetOneArgs struct {
	Id int64
}

type UserGetOneReply struct {
	User *entity.User
}

func (self *UserApi) GetOne(r *http.Request, args *UserGetOneArgs, reply *UserGetOneReply) error {
	user, err := self.persister.GetById(args.Id)

	if err != nil {
		return err
	}

	if user == nil {
		return errors.New("User not found")
	}

	reply.User = user

	return nil
}

type UserInsertArgs struct {
	User *entity.User
}

type UserInsertReply struct {
	User *entity.User
}

func (self *UserApi) Insert(r *http.Request, args *UserInsertArgs, reply *UserInsertReply) error {
	err := self.persister.Insert(args.User)

	if err != nil {
		return err
	}

	reply.User = args.User

	return nil
}

type UserUpdateArgs struct {
	User *entity.User
}

type UserUpdateReply struct {
	User *entity.User
}

func (self *UserApi) Update(r *http.Request, args *UserUpdateArgs, reply *UserUpdateReply) error {
	err := self.persister.Update(args.User)

	if err != nil {
		return err
	}

	reply.User = args.User

	return nil
}

type UserDeleteArgs struct {
	Id int64
}

type UserDeleteReply struct {
}

func (self *UserApi) Delete(r *http.Request, args *UserDeleteArgs, reply *UserDeleteReply) error {
	return self.persister.Delete(args.Id)
}
