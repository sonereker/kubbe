package handler

import (
	"github.com/gorilla/mux"
	"github.com/gosimple/slug"
	"github.com/jinzhu/gorm"
	"github.com/sonereker/kubbe/config"
	"github.com/sonereker/kubbe/model"
	"net/http"
	"strconv"
)

// GetShowPlacePage renders place page with base layout
func GetShowPlacePage(db *gorm.DB, c *config.Config, w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	var query *gorm.DB
	place := model.Place{}

	if isInteger(id) {
		query = db.Preload("Contents").First(&place, id)
	} else {
		query = db.Preload("Contents").Table("places").Joins("left join contents on contents."+
			"place_id = places.id").Where("contents.slug = ? and contents.status = ?", id, model.Published).First(&place)
	}

	if err := query.Error; err != nil {
		RenderError(w, "base", http.StatusNotFound)
		return
	}

	RenderTemplate(w, "base", "places/show", PageData{c.App.Title, place})
}

func GetNewPlacePage(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	RenderTemplate(w, "manager", "places/new", nil)
}

func CreatePlace(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	_ = r.ParseForm()

	content := model.Content{}
	content.Title = r.Form.Get("title")
	content.Slug = slug.Make(content.Title)
	content.Description = r.Form.Get("description")
	content.Status = model.Draft

	place := model.Place{}
	place.Contents = append(place.Contents, content)
	place.Lat = r.Form.Get("lat")
	place.Lon = r.Form.Get("lon")

	db.Create(&place)

	http.Redirect(w, r, "/", http.StatusMovedPermanently)
}

// GetEditPlacePage renders edit place page
func GetEditPlacePage(db *gorm.DB, c *config.Config, w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	var query *gorm.DB
	place := model.Place{}

	query = db.Preload("Contents").Table("places").Joins("left join contents on contents."+
		"place_id = places.id").Where("contents.id = ? and contents.status = ?", id, model.Published).First(&place)

	if err := query.Error; err != nil {
		RenderError(w, "base", http.StatusNotFound)
		return
	}

	RenderTemplate(w, "base", "places/edit", PageData{c.App.Title, place})
}

func isInteger(s string) bool {
	_, err := strconv.Atoi(s)
	return err == nil
}
