package handlers

import (
	"fmt"
	"github.com/gorilla/context"
	"github.com/jmoiron/sqlx"
	"github.com/cjdell/go_angular_starter/model/entity"
	"github.com/cjdell/go_angular_starter/model/persister"
	"net/http"
)

type ContextKey int

const (
	Database ContextKey = 0
	UserKey  ContextKey = 1
)

func CheckUser(h http.Handler, db *sqlx.DB) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		apiKey := r.Header.Get("API-Key")

		persister := persister.NewUserPersister(db)
		user, err := persister.GetByApiKey(apiKey)

		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintf(w, err.Error())
			return
		}

		if user == nil {
			w.WriteHeader(http.StatusForbidden)
			fmt.Fprintf(w, "Valid API key is required for this service")
			return
		}

		context.Set(r, UserKey, user)

		h.ServeHTTP(w, r)
	})
}

func GetUser(r *http.Request) *entity.User {
	return context.Get(r, UserKey).(*entity.User)
}
