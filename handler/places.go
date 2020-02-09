package handler

import (
	"github.com/gorilla/mux"
	"github.com/gosimple/slug"
	"github.com/jinzhu/gorm"
	"github.com/sonereker/kubbe/model"
	"net/http"
	"strconv"
)

// ShowPlace renders place page with base layout
func (p *Page) ShowPlace(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	var query *gorm.DB
	place := model.Place{}

	if isInteger(id) {
		query = p.DB.Preload("Contents").First(&place, id)
	} else {
		query = p.DB.Preload("Contents").Table("places").Joins("left join contents on contents."+
			"place_id = places.id").Where("contents.slug = ? and contents.status = ?", id, model.Published).First(&place)
	}

	if err := query.Error; err != nil {
		RenderError(w, Base, http.StatusNotFound)
		return
	}

	RenderTemplate(w, Base, "places/show", PageData{p.Config.Title, place})
}

func (p *Page) NewPlace(w http.ResponseWriter, r *http.Request) {
	RenderTemplate(w, Management, "places/new", nil)
}

func (p *Page) CreatePlace(w http.ResponseWriter, r *http.Request) {
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

	p.DB.Create(&place)

	http.Redirect(w, r, "/", http.StatusMovedPermanently)
}

// EditPlace renders edit place page
func (p *Page) EditPlace(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	var query *gorm.DB
	place := model.Place{}

	query = p.DB.Preload("Contents").Table("places").Joins("left join contents on contents."+
		"place_id = places.id").Where("contents.id = ? and contents.status = ?", id, model.Published).First(&place)

	if err := query.Error; err != nil {
		RenderError(w, Base, http.StatusNotFound)
		return
	}

	RenderTemplate(w, Base, "places/edit", PageData{p.Config.Title, place})
}

func isInteger(s string) bool {
	_, err := strconv.Atoi(s)
	return err == nil
}
