package handler

import (
	"net/http"
)

func (p *Page) Login(w http.ResponseWriter, r *http.Request) {
	RenderTemplate(w, Auth, "sessions/new", PageData{p.Config.App.Title, nil})
}
