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
	Config *config.Config
	Router *mux.Router
	DB     *gorm.DB
}

func (a *App) Initialize(config *config.Config) {
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
	a.setRouters()
}

func (a *App) setRouters() {
	a.Router.HandleFunc("/", a.GetHomePage).Methods("GET")
	a.Router.PathPrefix("/styles/").Handler(http.StripPrefix("/styles/",
		http.FileServer(http.Dir("templates/assets/styles/"))))

	c := a.Router.PathPrefix("/manage").Subrouter()
	c.HandleFunc("/places/new", a.GetNewPlacePage).Methods("GET")
	c.HandleFunc("/places", a.CreatePlace).Methods("POST")
}

func (a *App) Run(host string) {
	origins := handlers.AllowedOrigins([]string{"*"})
	methods := handlers.AllowedMethods([]string{"GET", "HEAD", "POST", "PUT", "OPTIONS", "DELETE"})

	fmt.Printf("Running Kubbe on %s \n", host)
	log.Fatal(http.ListenAndServe(host, handlers.LoggingHandler(os.Stdout, handlers.CORS(origins,
		methods)(a.Router))))
}

// DBMigrate creates and migrates the tables also creates relationships if necessary
func DBMigrate(db *gorm.DB) *gorm.DB {
	db.AutoMigrate(&model.Author{}, &model.Place{}, &model.Content{})
	return db
}

func (a *App) GetHomePage(w http.ResponseWriter, r *http.Request) {
	handler.GetHomePage(a.DB, a.Config, w, r)
}

func (a *App) GetNewPlacePage(w http.ResponseWriter, r *http.Request) {
	handler.GetNewPlacePage(a.DB, w, r)
}

func (a *App) CreatePlace(w http.ResponseWriter, r *http.Request) {
	handler.CreatePlace(a.DB, w, r)
}
