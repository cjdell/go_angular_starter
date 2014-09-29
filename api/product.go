package api

import (
	"errors"
	"github.com/jmoiron/sqlx"
	"go_angular_starter/model/entity"
	"go_angular_starter/model/persister"
	"go_angular_starter/services"
	"net/http"
)

type ProductApi struct {
	db        *sqlx.DB
	persister *persister.ProductPersister
}

type ProductInfo struct {
	*entity.Product
	Images     []*services.Image
	Categories []*entity.Category
}

func NewProductApi(db *sqlx.DB) *ProductApi {
	return &ProductApi{db, persister.NewProductPersister(db)}
}

func (self *ProductApi) getProductInfo(product *entity.Product) (*ProductInfo, error) {
	images, _ := services.GetImages(product)

	categoryPersister := persister.NewCategoryPersister(self.db)

	categories := make([]*entity.Category, len(product.GetCategoryIds()))

	for i, cid := range product.GetCategoryIds() {
		categories[i], _ = categoryPersister.GetById(cid)
	}

	return &ProductInfo{product, images, categories}, nil
}

type ProductGetAllArgs struct {
}

type ProductGetAllReply struct {
	Products []*ProductInfo
}

func (self *ProductApi) GetAll(r *http.Request, args *ProductGetAllArgs, reply *ProductGetAllReply) error {
	products, err := self.persister.GetAll()

	if err != nil {
		return err
	}

	productInfos := make([]*ProductInfo, len(products), len(products))

	for i, product := range products {
		productInfos[i], _ = self.getProductInfo(product)
	}

	reply.Products = productInfos

	return nil
}

type ProductGetOneArgs struct {
	Id int64
}

type ProductGetOneReply struct {
	Product *ProductInfo
}

func (self *ProductApi) GetOne(r *http.Request, args *ProductGetOneArgs, reply *ProductGetOneReply) error {
	product, err := self.persister.GetById(args.Id)

	if err != nil {
		return err
	}

	if product == nil {
		return errors.New("Product not found")
	}

	reply.Product, _ = self.getProductInfo(product)

	return nil
}

type ProductInsertArgs struct {
	Product *entity.Product
}

type ProductInsertReply struct {
	Product *entity.Product
}

func (self *ProductApi) Insert(r *http.Request, args *ProductInsertArgs, reply *ProductInsertReply) error {
	err := self.persister.Insert(args.Product)

	if err != nil {
		return err
	}

	reply.Product = args.Product

	return nil
}

type ProductUpdateArgs struct {
	Product *entity.Product

	Extra struct {
		NewImageFileName string
		//NewCategoryIds   []int64
	}
}

type ProductUpdateReply struct {
	Product *ProductInfo
}

func (self *ProductApi) Update(r *http.Request, args *ProductUpdateArgs, reply *ProductUpdateReply) error {
	//product, err := self.persister.GetById(args.Product.Id)

	//product.Update(args.Product)

	//if args.Extra.NewCategoryIds != nil {
	//	product.SetCategoryIds(args.Extra.NewCategoryIds)
	//}

	err := self.persister.Update(args.Product)

	if err != nil {
		return err
	}

	if args.Extra.NewImageFileName != "" {
		services.AssignImage(args.Product, args.Extra.NewImageFileName, "primary_image")
	}

	reply.Product, _ = self.getProductInfo(args.Product)

	return nil
}

type ProductDeleteArgs struct {
	Id int64
}

type ProductDeleteReply struct {
}

func (self *ProductApi) Delete(r *http.Request, args *ProductDeleteArgs, reply *ProductDeleteReply) error {
	return self.persister.Delete(args.Id)
}
