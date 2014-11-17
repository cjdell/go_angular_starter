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
	authApi struct{}

	authRequest  struct{ *rest.Request }
	authResponse struct{ rest.ResponseWriter }

	authAction func(*services.AuthService, authResponse, authRequest) error
)

func NewAuthApi(db *sqlx.DB) http.Handler {
	api := authApi{}

	handler := &rest.ResourceHandler{}

	wrap := func(action authAction, db *sqlx.DB) rest.HandlerFunc {
		return transactionWrap(db, func(w rest.ResponseWriter, r *rest.Request, db persister.DB) error {
			service := services.NewAuthService(db)

			return action(service, authResponse{w}, authRequest{r})
		})
	}

	err := handler.SetRoutes(
		&rest.Route{"POST", "/sign-in", wrap(authAction(api.signIn), db)},
		&rest.Route{"POST", "/sign-up", wrap(authAction(api.signUp), db)},
	)

	if err != nil {
		log.Fatal(err)
	}

	return handler
}

func (authApi) signIn(service *services.AuthService, res authResponse, req authRequest) error {
	sireq := req.SignInRequest()

	result, err := service.SignIn(sireq.Email, sireq.Password)

	if err != nil {
		return err
	}

	return res.WriteJson(result)
}

func (authApi) signUp(service *services.AuthService, res authResponse, req authRequest) error {
	sireq := req.SignInRequest()

	_, err := service.RegisterUser(sireq.Email, sireq.Password)

	return err
}

type signInRequest struct {
	Email    string
	Password string
}

func (req authRequest) SignInRequest() *signInRequest {
	sireq := &signInRequest{}

	if err := req.DecodeJsonPayload(&sireq); err != nil {
		panic(err)
	}

	return sireq
}
