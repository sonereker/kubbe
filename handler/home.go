package handler

import (
	"github.com/sonereker/kubbe/model"
	"html/template"
	"net/http"

	"github.com/jinzhu/gorm"
)

var tmpl *template.Template

func init() {
	// Parse one-time to reduce response time for home page
	tmpl = template.Must(template.ParseFiles("templates/layouts/base.html", "templates/home/index.html"))
}

func GetHomePage(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	var places []model.Place
	db.Preload("Contents").Find(&places)

	tmpl.ExecuteTemplate(w, "base", places)
}
