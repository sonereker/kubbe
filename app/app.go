package app

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/sonereker/kubbe/config"
	"github.com/sonereker/kubbe/handler"
	"github.com/sonereker/kubbe/model"
)

type App struct {
	Config      *config.Config
	Router      *mux.Router
	DB          *gorm.DB
	PageHandler *handler.Page
}

// Init; initializes a new App instance
func (a *App) Init(config *config.Config) {
	dbInfo := fmt.Sprintf("postgres://%s:%s@%s/%s?sslmode=%s",
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
	a.Config = config
	a.PageHandler = &handler.Page{
		DB:     a.DB,
		Config: a.Config,
	}

	a.setRouters()
}

func (a *App) setRouters() {
	a.Router.HandleFunc("/", a.PageHandler.Home).Methods("GET")

	// Login
	a.Router.HandleFunc("/login", a.PageHandler.Login).Methods("GET")

	// Places
	a.Router.HandleFunc("/{id}", a.PageHandler.ShowPlace).Methods("GET")
	a.Router.HandleFunc("/{id: [a-z]+}", a.PageHandler.ShowPlace).Methods("GET")

	a.Router.HandleFunc("/places/{id}/edit", a.PageHandler.EditPlace).Methods("GET")

	// Statics
	a.Router.PathPrefix("/styles/").Handler(http.StripPrefix("/styles/",
		http.FileServer(http.Dir("templates/assets/styles/"))))

	// Management
	c := a.Router.PathPrefix("/manage").Subrouter()
	c.HandleFunc("/places/new", a.PageHandler.NewPlace).Methods("GET")
	c.HandleFunc("/places", a.PageHandler.CreatePlace).Methods("POST")
}

// Run; sets allowed origins and methods then starts http server
func (a *App) Run(port string) {
	origins := handlers.AllowedOrigins([]string{"*"})
	methods := handlers.AllowedMethods([]string{"GET", "HEAD", "POST", "PUT", "OPTIONS", "DELETE"})

	fmt.Printf("Running Kubbe on %s \n", port)
	log.Fatal(http.ListenAndServe(":"+port, handlers.LoggingHandler(os.Stdout, handlers.CORS(origins,
		methods)(a.Router))))
}

// DBMigrate creates tables, runs migrations also creates relationships if necessary
func DBMigrate(db *gorm.DB) *gorm.DB {
	db.AutoMigrate(&model.User{}, &model.Place{}, &model.Content{})
	return db
}
