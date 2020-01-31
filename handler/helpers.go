package handler

import (
	"github.com/sonereker/kubbe/config"
	"html/template"
	"log"
	"net/http"
)

type AppI interface {
	Initialize(config *config.Config)
	setRouters()
	Run(host string)
	GetHomePage(w http.ResponseWriter, r *http.Request)
	GetNewPlacePage(w http.ResponseWriter, r *http.Request)
	CreatePlace(w http.ResponseWriter, r *http.Request)
}

type PageData struct {
	AppTitle string
	Contents interface{}
}

// RenderTemplate renders HTML template with the name using provided data
func RenderTemplate(w http.ResponseWriter, name string, data interface{}) {
	tmpl, err := template.ParseFiles("templates/"+name+".html", "templates/layouts/manager.html")
	if err != nil {
		log.Fatal(err)
	}

	_ = tmpl.ExecuteTemplate(w, "manager", data)
}
