package api

import (
	"github.com/ant0ine/go-json-rest/rest"
	"github.com/cjdell/go_angular_starter/model/entity"
	"github.com/cjdell/go_angular_starter/model/persister"
	"github.com/gorilla/context"
	"github.com/jmoiron/sqlx"
	"log"
	"net/http"
	"strconv"
	"time"
)

type ContextKey int

const (
	UserKey ContextKey = 1
)

func GetUser(r *http.Request) *entity.User {
	return context.Get(r, UserKey).(*entity.User)
}

func SetUser(r *http.Request, user *entity.User) {
	context.Set(r, UserKey, user)
}

func pathInt(r *rest.Request, key string) *int64 {
	v := r.PathParam(key)

	if v == "" {
		return nil
	}

	id, err := strconv.ParseInt(v, 10, 64)

	if err != nil {
		panic(err)
	}

	return &id
}

func queryInt(r *rest.Request, key string) *int64 {
	q := r.URL.Query()

	v := q[key]

	if len(v) == 0 {
		return nil
	}

	id, err := strconv.ParseInt(v[0], 10, 64)

	if err != nil {
		panic(err)
	}

	return &id
}

type AuthError struct {
}

func (AuthError) Error() string { return "Authentication required" }

type HandlerFuncWithError func(w rest.ResponseWriter, r *rest.Request, db persister.DB) error

// Wraps a handler func that returns an error into a new handler that runs inside a transaction
func transactionWrap(db *sqlx.DB, f HandlerFuncWithError) rest.HandlerFunc {
	return func(w rest.ResponseWriter, r *rest.Request) {
		var err error

		time.Sleep(200 * time.Millisecond)

		tx, err := db.Beginx()

		defer func() {
			if err != nil && tx != nil {
				tx.Rollback()
				log.Println("Transaction rolled back")
				return
			}

			err = tx.Commit()

			if err != nil {
				panic(err)
			} else {
				log.Println("Transaction commited")
			}
		}()

		if err != nil {
			rest.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		err = f(w, r, tx)

		if err != nil {
			rest.Error(w, err.Error(), http.StatusInternalServerError)
		}
	}
}
