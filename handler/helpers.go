package handler

import (
	"html/template"
	"log"
	"net/http"
)

// RenderTemplate renders HTML template using provided data and template name
func RenderTemplate(w http.ResponseWriter, fileName string, data interface{}) {
	tmpl, err := template.ParseFiles("template/" + fileName + ".html")
	if err != nil {
		log.Fatal(err)
	}
	_ = tmpl.Execute(w, data)
}
