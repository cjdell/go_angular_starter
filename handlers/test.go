package handlers

import (
	"net/http"
)

func (self AppHandlers) TestHandler() AppHandler {
	return func(w http.ResponseWriter, r *http.Request) error {
		w.Write([]byte("Hello"))

		return nil
	}
}
