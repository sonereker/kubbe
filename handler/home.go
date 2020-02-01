package handler

import (
	"github.com/jinzhu/gorm"
	"github.com/sonereker/kubbe/config"
	"github.com/sonereker/kubbe/model"
	"html/template"
	"net/http"
)

var tmpl *template.Template

func init() {
	// Parse one-time to reduce response time for home page
	tmpl = template.Must(template.ParseFiles("templates/layouts/base.html", "templates/home/index.html"))
}

func GetHomePage(db *gorm.DB, c *config.Config, w http.ResponseWriter, r *http.Request) {
	var places []model.Place
	db.Preload("Contents").Table("places").Joins("left join contents on contents."+
		"place_id = places.id").Where("contents.status = ?", model.Published).Find(&places)

	tmpl.ExecuteTemplate(w, "base", PageData{c.App.Title, places})
}
