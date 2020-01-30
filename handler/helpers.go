package handler

import (
	"html/template"
	"log"
	"net/http"
)

// RenderTemplate renders HTML template with the name using provided data
func RenderTemplate(w http.ResponseWriter, name string, data interface{}) {
	tmpl, err := template.ParseFiles("templates/"+name+".html", "templates/layouts/manager.html")
	if err != nil {
		log.Fatal(err)
	}
	_ = tmpl.ExecuteTemplate(w, "manager", data)
}
