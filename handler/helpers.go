package handler

import (
	"html/template"
	"log"
	"net/http"
)

type PageData struct {
	AppTitle string
	Data     interface{}
}

// RenderTemplate renders HTML template with the name using provided layout and data
func RenderTemplate(w http.ResponseWriter, layout string, name string, data interface{}) {
	tmpl, err := template.ParseFiles("templates/"+name+".html", "templates/layouts/"+layout+".html")
	if err != nil {
		log.Fatal(err)
	}

	_ = tmpl.ExecuteTemplate(w, layout, data)
}

// RenderError renders an error page using provided layout and errorCode
func RenderError(w http.ResponseWriter, layout string, errorCode int) {
	tmpl, err := template.ParseFiles("templates/error.html", "templates/layouts/"+layout+".html")
	if err != nil {
		log.Fatal(err)
	}

	_ = tmpl.ExecuteTemplate(w, layout, errorCode)
}
