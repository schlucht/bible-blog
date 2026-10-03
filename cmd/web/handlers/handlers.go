package handlers

import (
	"net/http"

	"github.com/schlucht/bibleblog/cmd/web/config"
	"github.com/schlucht/bibleblog/cmd/web/frontend"
)

var App *config.AppConfig

type Handler struct {
	App *config.AppConfig
}

func NewHandler(a *config.AppConfig) *Handler {
	return &Handler{
		App: a,
	}
}

func (h *Handler) Home(w http.ResponseWriter, r *http.Request) {
	data := make(map[string]interface{})
	data["test"] = []string{"Hallo Welt", "Hallo Lothar"}
	td := &frontend.TemplateData{
		Data: data,
	}
	if err := frontend.RenderTemplate(w, r, "home", td); err != nil {
		h.App.ErrorLog.Println(err)
	}
}

func (h *Handler) NotFound(w http.ResponseWriter, r *http.Request) {
	if err := frontend.RenderTemplate(w, r, "404", nil); err != nil {
		h.App.ErrorLog.Println(err)
	}
}
