package api

import (
	"errors"
	"github.com/jmoiron/sqlx"
	"go_angular_starter/model/entity"
	"go_angular_starter/model/persister"
	"go_angular_starter/services"
	"net/http"
)

type CategoryApi struct {
	db        *sqlx.DB
	persister *persister.CategoryPersister
}

type CategoryInfo struct {
	*entity.Category
	Parent *entity.Category
}

func NewCategoryApi(db *sqlx.DB) *CategoryApi {
	return &CategoryApi{db, persister.NewCategoryPersister(db)}
}

type CategoryGetAllArgs struct {
}

type CategoryGetAllReply struct {
	Categories []*entity.Category
}

func (self *CategoryApi) GetAll(r *http.Request, args *CategoryGetAllArgs, reply *CategoryGetAllReply) error {
	Categories, err := self.persister.GetAll()

	if err != nil {
		return err
	}

	reply.Categories = Categories

	return nil
}

type CategoryGetChildrenArgs struct {
	ParentId int64
}

type CategoryGetChildrenReply struct {
	Categories []*entity.Category
}

func (self *CategoryApi) GetChildren(r *http.Request, args *CategoryGetChildrenArgs, reply *CategoryGetChildrenReply) error {
	Categories, err := self.persister.GetChildren(args.ParentId)

	if err != nil {
		return err
	}

	reply.Categories = Categories

	return nil
}

type CategoryGetOneArgs struct {
	Id int64
}

type CategoryGetOneReply struct {
	Category *CategoryInfo
}

func (self *CategoryApi) GetOne(r *http.Request, args *CategoryGetOneArgs, reply *CategoryGetOneReply) error {
	category, err := self.persister.GetById(args.Id)

	if err != nil {
		return err
	}

	if category == nil {
		return errors.New("Category not found")
	}

	var parent *entity.Category = nil

	if category.ParentId != 0 {
		parent, err = self.persister.GetById(category.ParentId)
	}

	reply.Category = &CategoryInfo{category, parent}

	return nil
}

type CategoryInsertArgs struct {
	Category *entity.Category
}

type CategoryInsertReply struct {
	Category *entity.Category
}

func (self *CategoryApi) Insert(r *http.Request, args *CategoryInsertArgs, reply *CategoryInsertReply) error {
	err := self.persister.Insert(args.Category)

	if err != nil {
		return err
	}

	tx, _ := self.db.Beginx()

	err = services.GenerateFullyQualifiedNames(tx, args.Category.Id)

	if err != nil {
		return err
	} else {
		tx.Commit()
	}

	reply.Category = args.Category

	return nil
}

type CategoryUpdateArgs struct {
	Category *entity.Category
}

type CategoryUpdateReply struct {
	Category *CategoryInfo
}

func (self *CategoryApi) Update(r *http.Request, args *CategoryUpdateArgs, reply *CategoryUpdateReply) error {
	err := self.persister.Update(args.Category)

	if err != nil {
		return err
	}

	tx, _ := self.db.Beginx()

	err = services.GenerateFullyQualifiedNames(tx, args.Category.Id)

	if err != nil {
		return err
	} else {
		tx.Commit()
	}

	var parent *entity.Category = nil

	if args.Category.ParentId != 0 {
		parent, err = self.persister.GetById(args.Category.ParentId)
	}

	reply.Category = &CategoryInfo{args.Category, parent}

	return nil
}

type CategoryDeleteArgs struct {
	Id int64
}

type CategoryDeleteReply struct {
}

func (self *CategoryApi) Delete(r *http.Request, args *CategoryDeleteArgs, reply *CategoryDeleteReply) error {
	return self.persister.Delete(args.Id)
}
