package api

import (
	"github.com/ant0ine/go-json-rest/rest"
	"github.com/cjdell/go_angular_starter/model/persister"
	"github.com/cjdell/go_angular_starter/services"
	"github.com/jmoiron/sqlx"
	"log"
	"net/http"
)

type (
	productApi struct{}

	productRequest  struct{ *rest.Request }
	productResponse struct{ rest.ResponseWriter }

	productAction func(*services.ProductService, productResponse, productRequest) error
)

func NewProductApi(db *sqlx.DB) http.Handler {
	api := productApi{}

	handler := &rest.ResourceHandler{}

	wrap := func(action productAction, db *sqlx.DB, requireUser bool) rest.HandlerFunc {
		return transactionWrap(db, func(w rest.ResponseWriter, r *rest.Request, db persister.DB) error {
			user := GetUser(r.Request)

			if requireUser && user == nil {
				return AuthError{}
			}

			service := services.NewProductService(db, user)

			return action(service, productResponse{w}, productRequest{r})
		})
	}

	err := handler.SetRoutes(
		&rest.Route{"GET", "/", wrap(api.getAll, db, false)},
		&rest.Route{"GET", "/:id", wrap(api.getOne, db, false)},
		&rest.Route{"POST", "/", wrap(api.post, db, true)},
		&rest.Route{"PUT", "/:id", wrap(api.put, db, true)},
		&rest.Route{"DELETE", "/:id", wrap(api.delete, db, true)},
	)

	if err != nil {
		log.Fatal(err)
	}

	return handler
}

func (productApi) getAll(service *services.ProductService, res productResponse, req productRequest) error {
	products, err := service.GetAll(req.CategoryId(), nil)

	if err != nil {
		return err
	}

	return res.WriteProducts(products)
}

func (productApi) getOne(service *services.ProductService, res productResponse, req productRequest) error {
	product, err := service.GetOne(*req.Id())

	if err != nil {
		return err
	}

	return res.WriteProduct(product)
}

func (productApi) post(service *services.ProductService, res productResponse, req productRequest) error {
	var err error
	var product *services.ProductInfo

	if product, err = service.Insert(req.Product()); err != nil {
		return err
	}

	return res.WriteProduct(product)
}

func (productApi) put(service *services.ProductService, res productResponse, req productRequest) error {
	var err error
	var product *services.ProductInfo

	if product, err = service.Update(req.Product()); err != nil {
		return err
	}

	return res.WriteProduct(product)
}

func (productApi) delete(service *services.ProductService, res productResponse, req productRequest) error {
	if err := service.Delete(*req.Id()); err != nil {
		return err
	}

	res.WriteHeader(http.StatusOK)

	return nil
}

func (req productRequest) Id() *int64 {
	return pathInt(req.Request, "id")
}

func (req productRequest) CategoryId() *int64 {
	return pathInt(req.Request, "category_id")
}

func (req productRequest) Product() *services.ProductSave {
	product := &services.ProductSave{}

	if err := req.DecodeJsonPayload(&product); err != nil {
		panic(err)
	}

	if id := req.Id(); id != nil {
		product.Id = *id
	}

	return product
}

func (res productResponse) WriteProducts(products []*services.ProductInfo) error {
	return res.WriteJson(products)
}

func (res productResponse) WriteProduct(product *services.ProductInfo) error {
	return res.WriteJson(product)
}
