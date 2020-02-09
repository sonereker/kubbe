package handler

import (
	"github.com/jinzhu/gorm"
	"github.com/sonereker/kubbe/config"
	"html/template"
	"log"
	"net/http"
)

const (
	Base       Layout = "base"
	Auth       Layout = "auth"
	Management Layout = "manage"
)

type (
	Page struct {
		DB     *gorm.DB
		Config *config.Config
	}

	PageData struct {
		AppTitle string
		Data     interface{}
	}

	Layout string
)

// RenderTemplate renders HTML template with the name using provided layout and data
func RenderTemplate(w http.ResponseWriter, l Layout, name string, data interface{}) {
	tmpl, err := template.ParseFiles("templates/"+name+".html", "templates/layouts/"+string(l)+".html")
	if err != nil {
		log.Fatal(err)
	}

	_ = tmpl.ExecuteTemplate(w, string(l), data)
}

// RenderError renders an error page using provided layout and errorCode
func RenderError(w http.ResponseWriter, l Layout, errorCode int) {
	tmpl, err := template.ParseFiles("templates/error.html", "templates/layouts/"+string(l)+".html")
	if err != nil {
		log.Fatal(err)
	}

	_ = tmpl.ExecuteTemplate(w, string(l), errorCode)
}
