package handler

import (
	"github.com/sonereker/kubbe/model"
	"net/http"

	"github.com/jinzhu/gorm"
)

func GetHomePage(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	var places []model.Place
	db.Preload("Contents").Find(&places)

	RenderTemplate(w, "home", places)
}
