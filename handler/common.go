package handler

import (
	"fmt"
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
		Config *config.AppConfig
	}

	PageData struct {
		AppTitle string
		Data     interface{}
	}

	Layout string
)

// RenderTemplate renders HTML template with the name using provided layout and data
func RenderTemplate(w http.ResponseWriter, layout Layout, name string, data interface{}) {
	tmpl, err := template.ParseFiles(fmt.Sprintf("templates/%s.html", name),
		fmt.Sprintf("templates/layouts/%s.html", layout))
	if err != nil {
		log.Fatal(err)
	}

	_ = tmpl.ExecuteTemplate(w, string(layout), data)
}

// RenderError renders an error page using provided layout and errorCode
func RenderError(w http.ResponseWriter, layout Layout, errorCode int) {
	tmpl, err := template.ParseFiles("templates/error.html", fmt.Sprintf("templates/layouts/%s.html", layout))
	if err != nil {
		log.Fatal(err)
	}

	_ = tmpl.ExecuteTemplate(w, string(layout), errorCode)
}
