package handler

import (
	"github.com/gosimple/slug"
	"github.com/jinzhu/gorm"
	"github.com/sonereker/kubbe/model"
	"net/http"
)

func GetNewPlacePage(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	RenderTemplate(w, "places/new", nil)
}

func CreatePlace(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	_ = r.ParseForm()

	content := model.Content{}
	content.Title = r.Form.Get("title")
	content.Slug = slug.Make(content.Title)
	content.Description = r.Form.Get("description")

	place := model.Place{}
	place.Contents = append(place.Contents, content)
	place.Lat = r.Form.Get("lat")
	place.Lon = r.Form.Get("lon")

	db.Create(&place)

	http.Redirect(w, r, "/", http.StatusMovedPermanently)
}
