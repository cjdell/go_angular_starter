package handlers

import (
	"github.com/cjdell/go_angular_starter/services"
	"github.com/gorilla/mux"
	"html/template"
	"net/http"
)

type productContent struct {
	title   string
	product *services.ProductInfo
}

func (self productContent) Title() string {
	return self.title
}

func (self productContent) ContentBody() interface{} {
	// To prevent the HTML being escaped by the template engine, we need to add an extra property
	return &struct {
		*services.ProductInfo
		DescriptionHtml interface{}
	}{self.product, template.HTML(self.product.Description)}
}

func (self AppHandlers) ProductHandler() AppHandler {
	return func(w http.ResponseWriter, r *http.Request) error {
		var (
			vars           = mux.Vars(r)
			handle         = vars["handle"]
			productService = services.NewProductService(self.db, nil)
		)

		product, err := productService.GetByHandle(handle)

		if err != nil {
			return err
		}

		return self.Render(&productContent{product.Name, product}, []string{"pages/product.html"}, w)
	}
}
