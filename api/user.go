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
	userApi struct{}

	userRequest  struct{ *rest.Request }
	userResponse struct{ rest.ResponseWriter }

	userAction func(*services.UserService, userResponse, userRequest) error
)

func NewUserApi(db *sqlx.DB) http.Handler {
	api := userApi{}

	handler := &rest.ResourceHandler{}

	wrap := func(action userAction, db *sqlx.DB, requireUser bool) rest.HandlerFunc {
		return transactionWrap(db, func(w rest.ResponseWriter, r *rest.Request, db persister.DB) error {
			user := GetUser(r.Request)

			if requireUser && user == nil {
				return AuthError{}
			}

			service := services.NewUserService(db, user)

			return action(service, userResponse{w}, userRequest{r})
		})
	}

	err := handler.SetRoutes(
		&rest.Route{"GET", "/", wrap(userAction(api.getAll), db, false)},
		&rest.Route{"GET", "/:id", wrap(userAction(api.getOne), db, false)},
		&rest.Route{"POST", "/", wrap(userAction(api.post), db, true)},
		&rest.Route{"PUT", "/:id", wrap(userAction(api.put), db, true)},
		&rest.Route{"DELETE", "/:id", wrap(userAction(api.delete), db, true)},
	)

	if err != nil {
		log.Fatal(err)
	}

	return handler
}

func (userApi) getAll(service *services.UserService, res userResponse, req userRequest) error {
	users, err := service.GetAll(nil)

	if err != nil {
		return err
	}

	return res.WriteUsers(users)
}

func (userApi) getOne(service *services.UserService, res userResponse, req userRequest) error {
	user, err := service.GetOne(*req.Id())

	if err != nil {
		return err
	}

	return res.WriteUser(user)
}

func (userApi) post(service *services.UserService, res userResponse, req userRequest) error {
	var err error
	var user *services.UserInfo

	if user, err = service.Insert(req.User()); err != nil {
		return err
	}

	return res.WriteUser(user)
}

func (userApi) put(service *services.UserService, res userResponse, req userRequest) error {
	var err error
	var user *services.UserInfo

	if user, err = service.Update(req.User()); err != nil {
		return err
	}

	return res.WriteUser(user)
}

func (userApi) delete(service *services.UserService, res userResponse, req userRequest) error {
	if err := service.Delete(*req.Id()); err != nil {
		return err
	}

	res.WriteHeader(http.StatusOK)

	return nil
}

func (req userRequest) Id() *int64 {
	return pathInt(req.Request, "id")
}

func (req userRequest) User() *services.UserChanges {
	user := &services.UserChanges{}

	if err := req.DecodeJsonPayload(&user); err != nil {
		panic(err)
	}

	if id := req.Id(); id != nil {
		user.Id = *id
	}

	return user
}

func (res userResponse) WriteUsers(users []*services.UserInfo) error {
	return res.WriteJson(users)
}

func (res userResponse) WriteUser(user *services.UserInfo) error {
	return res.WriteJson(user)
}
