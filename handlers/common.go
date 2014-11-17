package handlers

import (
	"github.com/cjdell/go_angular_starter/config"
	"github.com/cjdell/go_angular_starter/model/entity"
	"github.com/cjdell/go_angular_starter/model/persister"
	"github.com/jmoiron/sqlx"
	"html/template"
	"net/http"
	"path"
)

// Each method on this struct is for a page function (i.e. Home)
type AppHandlers struct {
	db *sqlx.DB
}

func NewAppHandlers(db *sqlx.DB) *AppHandlers {
	return &AppHandlers{db}
}

// Behaviours as a wrapper for handlers that produce an error
type AppHandler func(http.ResponseWriter, *http.Request) error

func (fn AppHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if err := fn(w, r); err != nil {
		http.Error(w, err.Error(), 500)
	}
}

type MainLayoutViewModel struct {
	Title      string
	Categories []*entity.Category
	Content    interface{}
}

type Content interface {
	Title() string
	ContentBody() interface{}
}

func (self AppHandlers) Render(content Content, extraTempSrcs []string, w http.ResponseWriter) error {
	var (
		catPer  = persister.NewCategoryPersister(self.db)
		tp      = config.App.TemplateRoot
		tempSrc = []string{path.Join(tp, "layouts/main.html"), path.Join(tp, "partials/category_menu.html")}
		vm      = &MainLayoutViewModel{}
	)

	var err error

	var zero int64 = 0

	vm.Title = content.Title()
	vm.Categories, err = catPer.GetAll(&zero, nil)
	vm.Content = content.ContentBody()

	if err != nil {
		return err
	}

	for _, t := range extraTempSrcs {
		tempSrc = append(tempSrc, path.Join(tp, t))
	}

	t := template.Must(template.New("layout").ParseFiles(tempSrc...))

	return t.Execute(w, vm)
}

func RenderError(err error, w http.ResponseWriter) {
	w.Write([]byte(err.Error()))
}
