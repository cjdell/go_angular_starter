package handlers

import (
	"github.com/jmoiron/sqlx"
	"github.com/cjdell/go_angular_starter/api"
	"html/template"
	"net/http"
)

func HomeHandler(db *sqlx.DB) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		productService := api.NewProductApi(db)

		args := &api.ProductGetAllArgs{}
		reply := &api.ProductGetAllReply{}

		productService.GetAll(r, args, reply)

		p := &struct {
			Title    string
			Products []*api.ProductInfo
		}{Title: "Welcome", Products: reply.Products}

		t := template.Must(template.New("layout").ParseFiles("web/templates/layouts/main.html", "web/templates/pages/home.html"))

		t.Execute(w, p)
	}
}
