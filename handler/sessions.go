package handler

import (
	"github.com/jinzhu/gorm"
	"github.com/sonereker/kubbe/config"
	"net/http"
)

func GetLoginPage(db *gorm.DB, c *config.Config, w http.ResponseWriter, r *http.Request) {
	RenderTemplate(w, "sessions", "sessions/new", PageData{c.App.Title, nil})
}
