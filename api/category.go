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
	categoryApi struct{}

	categoryRequest  struct{ *rest.Request }
	categoryResponse struct{ rest.ResponseWriter }

	categoryAction func(*services.CategoryService, categoryResponse, categoryRequest) error
)

func NewCategoryApi(db *sqlx.DB) http.Handler {
	api := categoryApi{}

	handler := &rest.ResourceHandler{}

	wrap := func(action categoryAction, db *sqlx.DB, requireUser bool) rest.HandlerFunc {
		return transactionWrap(db, func(w rest.ResponseWriter, r *rest.Request, db persister.DB) error {
			user := GetUser(r.Request)

			if requireUser && user == nil {
				return AuthError{}
			}

			service := services.NewCategoryService(db, user)

			return action(service, categoryResponse{w}, categoryRequest{r})
		})
	}

	err := handler.SetRoutes(
		&rest.Route{"GET", "/", wrap(categoryAction(api.getAll), db, false)},
		&rest.Route{"GET", "/:id", wrap(categoryAction(api.getOne), db, false)},
		&rest.Route{"POST", "/", wrap(categoryAction(api.post), db, true)},
		&rest.Route{"PUT", "/:id", wrap(categoryAction(api.put), db, true)},
		&rest.Route{"DELETE", "/:id", wrap(categoryAction(api.delete), db, true)},
	)

	if err != nil {
		log.Fatal(err)
	}

	return handler
}

func (categoryApi) getAll(service *services.CategoryService, res categoryResponse, req categoryRequest) error {
	categories, err := service.GetAll(req.ParentId(), nil)

	if err != nil {
		return err
	}

	return res.WriteCategories(categories)
}

func (categoryApi) getOne(service *services.CategoryService, res categoryResponse, req categoryRequest) error {
	category, err := service.GetOne(*req.Id())

	if err != nil {
		return err
	}

	return res.WriteCategory(category)
}

func (categoryApi) post(service *services.CategoryService, res categoryResponse, req categoryRequest) error {
	var err error
	var category *services.CategoryInfo

	if category, err = service.Insert(req.Category()); err != nil {
		return err
	}

	return res.WriteCategory(category)
}

func (categoryApi) put(service *services.CategoryService, res categoryResponse, req categoryRequest) error {
	var err error
	var category *services.CategoryInfo

	if category, err = service.Update(req.Category()); err != nil {
		return err
	}

	return res.WriteCategory(category)
}

func (categoryApi) delete(service *services.CategoryService, res categoryResponse, req categoryRequest) error {
	if err := service.Delete(*req.Id()); err != nil {
		return err
	}

	res.WriteHeader(http.StatusOK)

	return nil
}

func (req categoryRequest) Id() *int64 {
	return pathInt(req.Request, "id")
}

func (req categoryRequest) ParentId() *int64 {
	return queryInt(req.Request, "parent_id")
}

func (req categoryRequest) Category() *services.CategoryChanges {
	category := &services.CategoryChanges{}

	if err := req.DecodeJsonPayload(&category); err != nil {
		panic(err)
	}

	if id := req.Id(); id != nil {
		category.Id = *id
	}

	return category
}

func (res categoryResponse) WriteCategories(categories []*services.CategoryInfo) error {
	return res.WriteJson(categories)
}

func (res categoryResponse) WriteCategory(category *services.CategoryInfo) error {
	return res.WriteJson(category)
}
