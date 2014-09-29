package main

import (
	"fmt"
	"github.com/gorilla/rpc"
	"github.com/gorilla/rpc/json"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/cjdell/go_angular_starter/api"
	"github.com/cjdell/go_angular_starter/config"
	"github.com/cjdell/go_angular_starter/handlers"
	"log"
	"net/http"
	"os"
	"path"
)

func main() {
	db, err := openDatabase()

	checkErr(err, "Could not open database. %s")

	startServer(db)
}

func openDatabase() (*sqlx.DB, error) {
	drv, open, err := config.App.DatabaseConfig()

	checkErr(err, "Could not find database config. %s")

	db, err := sqlx.Open(drv, open)

	return db, err
}

func startServer(db *sqlx.DB) {
	fmt.Printf("Using web root dir: %s\n", config.App.WebRoot())
	fmt.Printf("Using ENV: %s\n", config.App.Env())

	staticHandler := http.FileServer(http.Dir(config.App.WebRoot()))

	// Configure JSON-RPC based API (For auth, doesn't require authentication)
	authApiServer := rpc.NewServer()

	authApiServer.RegisterCodec(json.NewCodec(), "application/json")
	authApiServer.RegisterService(api.NewAuthApi(db), "")

	http.Handle("/auth", authApiServer)

	// Configure JSON-RPC based API (For everything else, requires API Key)
	apiServer := rpc.NewServer()

	apiServer.RegisterCodec(json.NewCodec(), "application/json")
	apiServer.RegisterService(api.NewUserApi(db), "")
	apiServer.RegisterService(api.NewProductApi(db), "")
	apiServer.RegisterService(api.NewCategoryApi(db), "")

	// Firewall the API to only authenticated users
	http.Handle("/api", handlers.CheckUser(apiServer, db))

	// Configure handlers
	dynamicHandler := http.NewServeMux()

	dynamicHandler.HandleFunc("/", handlers.HomeHandler(db))
	dynamicHandler.HandleFunc("/test", handlers.TestHandler)
	dynamicHandler.HandleFunc("/upload", handlers.UploadHandler)

	// Merge the two handlers for the root path
	http.Handle("/", MergeHandlers(staticHandler, dynamicHandler))

	err := http.ListenAndServe(":3000", nil)

	checkErr(err, "Could not start HTTP server. %s")
}

func checkErr(err error, msg string) {
	if err != nil {
		log.Fatalf(msg, err)
	}
}

// This function allows handlers to co-exist on the "/" path. URLs that don't match files are assumed to be destined for handlers. TODO: Improve this...
func MergeHandlers(staticHandler http.Handler, dynamicHandler *http.ServeMux) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assetPath := path.Join(config.App.WebRoot(), r.URL.Path)

		log.Printf("Asset path: %s", assetPath)

		if _, err := os.Stat(assetPath); os.IsNotExist(err) || r.URL.Path == "/" {
			log.Printf("DYNAMIC")
			dynamicHandler.ServeHTTP(w, r)
		} else {
			log.Printf("STATIC")
			staticHandler.ServeHTTP(w, r)
		}
	})
}
