package main

import (
	"fmt"
	"github.com/cjdell/go_angular_starter/api"
	"github.com/cjdell/go_angular_starter/config"
	"github.com/cjdell/go_angular_starter/handlers"
	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"log"
	"net/http"
)

func main() {
	db, err := openDatabase()

	if err != nil {
		log.Fatalf("Could not open database. %s", err)
	}

	err = startServer(db)

	if err != nil {
		log.Fatalf("Could not start HTTP server. %s", err)
	}
}

func openDatabase() (*sqlx.DB, error) {
	return sqlx.Open(config.App.DatabaseDriver, config.App.DatabaseOpen)
}

func startServer(db *sqlx.DB) error {
	http.Handle("/auth/", http.StripPrefix("/auth", api.NewAuthApi(db)))

	// ----------------------------------------------------------------

	apiRouter := mux.NewRouter()

	productApi := http.StripPrefix("/products", api.NewProductApi(db))

	apiRouter.Handle("/products", productApi)
	apiRouter.Handle("/products/{id:[0-9]+}", productApi)

	categoryApi := http.StripPrefix("/categories", api.NewCategoryApi(db))

	apiRouter.Handle("/categories", categoryApi)
	apiRouter.Handle("/categories/{id:[0-9]+}", categoryApi)

	userApi := http.StripPrefix("/users", api.NewUserApi(db))

	apiRouter.Handle("/users", userApi)
	apiRouter.Handle("/users/{id:[0-9]+}", userApi)

	// GENERATOR INJECT

	http.Handle("/api/", http.StripPrefix("/api", handlers.CheckUser(apiRouter, db, false)))

	// ----------------------------------------------------------------

	// Configure handlers
	dynamicHandler := mux.NewRouter()

	appHandlers := handlers.NewAppHandlers(db)

	dynamicHandler.Handle("/", appHandlers.HomeHandler())
	dynamicHandler.Handle("/test", appHandlers.TestHandler())
	dynamicHandler.Handle("/upload", appHandlers.UploadHandler())
	dynamicHandler.Handle("/{handle}", appHandlers.HomeHandler())
	dynamicHandler.Handle("/product/{handle}", appHandlers.ProductHandler())

	http.Handle("/", dynamicHandler)

	// ----------------------------------------------------------------

	assetHandler := http.FileServer(http.Dir(config.App.AssetRoot))
	http.Handle("/assets/", http.StripPrefix("/assets/", assetHandler))

	adminHandler := http.FileServer(http.Dir(config.App.AdminRoot))
	http.Handle("/admin/", http.StripPrefix("/admin/", adminHandler))

	// ----------------------------------------------------------------

	fmt.Printf("Starting server on port %s using Env: %s\n", config.App.ListenAddress, config.App.Env())

	return http.ListenAndServe(config.App.ListenAddress, nil)
}
