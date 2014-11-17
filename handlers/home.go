package handlers

import (
	"github.com/cjdell/go_angular_starter/model/entity"
	"github.com/cjdell/go_angular_starter/model/persister"
	"github.com/cjdell/go_angular_starter/services"
	"github.com/gorilla/mux"
	"net/http"
)

type homeContent struct {
	title    string
	products []*services.ProductInfo
}

func (self homeContent) Title() string {
	return self.title
}

func (self homeContent) ContentBody() interface{} {
	return self.products
}

func (self AppHandlers) HomeHandler() AppHandler {
	return func(w http.ResponseWriter, r *http.Request) error {
		vars := mux.Vars(r)
		handle := vars["handle"]

		categoryPersister := persister.NewCategoryPersister(self.db)
		productService := services.NewProductService(self.db, nil)

		var (
			title    = ""
			category = (*entity.Category)(nil)
			products = ([]*services.ProductInfo)(nil)
		)

		err := error(nil)

		if handle == "" {
			title = "Home"
			products, err = productService.GetAll(nil, nil)
		} else {
			category, _ = categoryPersister.GetByHandle(handle)
			title = category.FqName
			products, err = productService.GetAll(&category.Id, nil)
		}

		if err != nil {
			return err
		}

		return self.Render(&homeContent{title, products}, []string{"pages/home.html", "partials/product_list.html"}, w)
	}
}
