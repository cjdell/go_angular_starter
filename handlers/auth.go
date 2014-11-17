package handlers

import (
	"fmt"
	"github.com/cjdell/go_angular_starter/api"
	"github.com/cjdell/go_angular_starter/model/persister"
	"github.com/jmoiron/sqlx"
	"net/http"
)

func CheckUser(h http.Handler, db *sqlx.DB, userRequired bool) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		apiKey := r.Header.Get("API-Key")

		userPersister := persister.NewUserPersister(db)
		user, err := userPersister.GetByApiKey(apiKey)

		if _, ok := err.(persister.NotFoundError); ok {
			user = nil
		} else if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintf(w, err.Error())
			return
		}

		if userRequired && user == nil {
			w.WriteHeader(http.StatusForbidden)
			fmt.Fprintf(w, "Valid API key is required for this service")
			return
		}

		api.SetUser(r, user)

		h.ServeHTTP(w, r)
	})
}
