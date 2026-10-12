package handlers

import (
	"encoding/json"
	"fmt"
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
	if err := frontend.RenderTemplate(w, r, "home", nil); err != nil {
		h.App.ErrorLog.Println(err)
	}
}

func (h *Handler) Buttons(w http.ResponseWriter, r *http.Request) {
	data := make(map[string]interface{})
	data["test"] = []string{"Hallo Welt", "Hallo Lothar", "Noch ein Button"}
	td := &frontend.TemplateData{
		Data: data,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(201)
	out, err := json.MarshalIndent(td, "", "\t")
	if err != nil {
		fmt.Println(err)
	}
	
	_, err = w.Write(out)
	if (err != nil){
		fmt.Println(err)
	}
}


func (h *Handler) NotFound(w http.ResponseWriter, r *http.Request) {
	if err := frontend.RenderTemplate(w, r, "404", nil); err != nil {
		h.App.ErrorLog.Println(err)
	}
}
