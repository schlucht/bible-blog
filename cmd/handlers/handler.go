package handlers

import (
	"net/http"

	"github.com/a-h/templ"
	"github.com/schlucht/bibleblog/cmd/application"
	"github.com/schlucht/bibleblog/frontend"
	"github.com/schlucht/bibleblog/frontend/page"
)

var App *application.Application

type Handler struct {
	App *application.Application
}

func NewHandler(a *application.Application) *Handler {
	return &Handler{
		App: a,
	}
}

func (h *Handler) Home(w http.ResponseWriter, r *http.Request) {
	templ.Handler(frontend.Base("Meine Seite", page.HomePage())).ServeHTTP(w, r)
}

func (h *Handler) NotFound(w http.ResponseWriter, r *http.Request) {
	templ.Handler(frontend.NotFound()).ServeHTTP(w, r)
}

