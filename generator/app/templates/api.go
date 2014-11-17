package api

import (
	"github.com/cjdell/go_angular_starter/model/persister"
	"github.com/cjdell/go_angular_starter/services"
	"github.com/ant0ine/go-json-rest/rest"
	"github.com/jmoiron/sqlx"
	"log"
	"net/http"
)

type (
	<%= entityNameSingularCamelCase %>Api struct{}

	<%= entityNameSingularCamelCase %>Request  struct{ *rest.Request }
	<%= entityNameSingularCamelCase %>Response struct{ rest.ResponseWriter }

	<%= entityNameSingularCamelCase %>Action func(*services.<%= entityNameSingularPascalCase %>Service, <%= entityNameSingularCamelCase %>Response, <%= entityNameSingularCamelCase %>Request) error
)

func New<%= entityNameSingularPascalCase %>Api(db *sqlx.DB) http.Handler {
	api := <%= entityNameSingularCamelCase %>Api{}

	handler := &rest.ResourceHandler{}

	wrap := func(action <%= entityNameSingularCamelCase %>Action, db *sqlx.DB, requireUser bool) rest.HandlerFunc {
		return transactionWrap(db, func(w rest.ResponseWriter, r *rest.Request, db persister.DB) error {
			user := GetUser(r.Request)

			if requireUser && user == nil {
				return AuthError{}
			}

			service := services.New<%= entityNameSingularPascalCase %>Service(db, user)

			return action(service, <%= entityNameSingularCamelCase %>Response{w}, <%= entityNameSingularCamelCase %>Request{r})
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

func (<%= entityNameSingularCamelCase %>Api) getAll(service *services.<%= entityNameSingularPascalCase %>Service, res <%= entityNameSingularCamelCase %>Response, req <%= entityNameSingularCamelCase %>Request) error {
	<%= entityNamePluralCamelCase %>, err := service.GetAll(nil)

	if err != nil {
		return err
	}

	return res.Write<%= entityNamePluralPascalCase %>(<%= entityNamePluralCamelCase %>)
}

func (<%= entityNameSingularCamelCase %>Api) getOne(service *services.<%= entityNameSingularPascalCase %>Service, res <%= entityNameSingularCamelCase %>Response, req <%= entityNameSingularCamelCase %>Request) error {
	<%= entityNameSingularCamelCase %>, err := service.GetOne(*req.Id())

	if err != nil {
		return err
	}

	return res.Write<%= entityNameSingularPascalCase %>(<%= entityNameSingularCamelCase %>)
}

func (<%= entityNameSingularCamelCase %>Api) post(service *services.<%= entityNameSingularPascalCase %>Service, res <%= entityNameSingularCamelCase %>Response, req <%= entityNameSingularCamelCase %>Request) error {
	var err error
	var <%= entityNameSingularCamelCase %> *services.<%= entityNameSingularPascalCase %>Info

	if <%= entityNameSingularCamelCase %>, err = service.Insert(req.<%= entityNameSingularPascalCase %>()); err != nil {
		return err
	}

	return res.Write<%= entityNameSingularPascalCase %>(<%= entityNameSingularCamelCase %>)
}

func (<%= entityNameSingularCamelCase %>Api) put(service *services.<%= entityNameSingularPascalCase %>Service, res <%= entityNameSingularCamelCase %>Response, req <%= entityNameSingularCamelCase %>Request) error {
	var err error
	var <%= entityNameSingularCamelCase %> *services.<%= entityNameSingularPascalCase %>Info

	if <%= entityNameSingularCamelCase %>, err = service.Update(req.<%= entityNameSingularPascalCase %>()); err != nil {
		return err
	}

	return res.Write<%= entityNameSingularPascalCase %>(<%= entityNameSingularCamelCase %>)
}

func (<%= entityNameSingularCamelCase %>Api) delete(service *services.<%= entityNameSingularPascalCase %>Service, res <%= entityNameSingularCamelCase %>Response, req <%= entityNameSingularCamelCase %>Request) error {
	if err := service.Delete(*req.Id()); err != nil {
		return err
	}

	res.WriteHeader(http.StatusOK)

	return nil
}

func (req <%= entityNameSingularCamelCase %>Request) Id() *int64 {
	return pathInt(req.Request, "id")
}

func (req <%= entityNameSingularCamelCase %>Request) <%= entityNameSingularPascalCase %>() *services.<%= entityNameSingularPascalCase %>Save {
	<%= entityNameSingularCamelCase %> := &services.<%= entityNameSingularPascalCase %>Save{}

	if err := req.DecodeJsonPayload(&<%= entityNameSingularCamelCase %>); err != nil {
		panic(err)
	}

	if id := req.Id(); id != nil {
		<%= entityNameSingularCamelCase %>.Id = *id
	}

	return <%= entityNameSingularCamelCase %>
}

func (res <%= entityNameSingularCamelCase %>Response) Write<%= entityNamePluralPascalCase %>(<%= entityNamePluralCamelCase %> []*services.<%= entityNameSingularPascalCase %>Info) error {
	return res.WriteJson(<%= entityNamePluralCamelCase %>)
}

func (res <%= entityNameSingularCamelCase %>Response) Write<%= entityNameSingularPascalCase %>(<%= entityNameSingularCamelCase %> *services.<%= entityNameSingularPascalCase %>Info) error {
	return res.WriteJson(<%= entityNameSingularCamelCase %>)
}
