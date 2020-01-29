package app

import (
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"os"
)

type App struct {
	Router *mux.Router
}

func (a *App) Initialize(config *config.Config) {
	/*dbInfo := fmt.Sprintf("postgres://%s:%s@%s/%s?sslmode=%s",
		config.DB.Username,
		config.DB.Password,
		config.DB.Host,
		config.DB.Name,
		config.DB.SSLMode,
	)

	db, err := gorm.Open(config.DB.Dialect, dbInfo)
	if err != nil {
		log.Fatal("Could not connect database: ", err)
	}

	a.DB = DBMigrate(db)
	a.Router = mux.NewRouter()
	a.setRouters()*/
}

func (a *App) setRouters() {
	/*c := a.Router.PathPrefix("/catalogs").Subrouter()
	c.HandleFunc("", handler.ValidateToken(a.GetAllCatalogs)).Methods("GET")
	c.HandleFunc("", handler.ValidateToken(a.CreateCatalog)).Methods("POST")
	c.HandleFunc("/{id}", handler.ValidateToken(a.GetCatalog)).Methods("GET")
	c.HandleFunc("/{id}", handler.ValidateToken(a.UpdateCatalog)).Methods("PUT")
	c.HandleFunc("/{id}", handler.ValidateToken(a.DeleteCatalog)).Methods("DELETE")

	c.HandleFunc("/{id}/assets", handler.ValidateToken(a.GetAllAssets)).Methods("GET")
	c.HandleFunc("/{id}/assets/new", handler.ValidateToken(a.CreateAsset)).Methods("POST")
	c.HandleFunc("/{id}/assets/{assetId}", handler.ValidateToken(a.GetAsset)).Methods("GET")
	c.HandleFunc("/{id}/assets/{assetId}", handler.ValidateToken(a.UpdateAsset)).Methods("PUT")
	c.HandleFunc("/{id}/assets/{assetId}", handler.ValidateToken(a.DeleteAsset)).Methods("DELETE")*/
}

func (a *App) Run(host string) {
	headers := handlers.AllowedHeaders([]string{"X-Requested-With", "Accept", "Content-Type", "Content-Length", "Accept-Encoding", "X-CSRF-Token", "Authorization"})
	origins := handlers.AllowedOrigins([]string{"*"})
	methods := handlers.AllowedMethods([]string{"GET", "HEAD", "POST", "PUT", "OPTIONS", "DELETE"})

	log.Fatal(http.ListenAndServe(host, handlers.LoggingHandler(os.Stdout, handlers.CORS(headers, origins, methods)(a.Router))))
}
