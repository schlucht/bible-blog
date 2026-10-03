package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
	"github.com/schlucht/bibleblog/cmd/web/config"
	"github.com/schlucht/bibleblog/cmd/web/handlers"
)

func routes(app *config.AppConfig) http.Handler {
	mux := chi.NewRouter()
	h := handlers.NewHandler(app)

	mux.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300,
	}))

	if !app.MySqlLog {

		mux.Get("/", h.Home)
	}

	mux.NotFound(h.NotFound)

	fileServer := http.FileServer(http.Dir("./assets"))
	mux.Handle("/assets/*", http.StripPrefix("/assets", fileServer))

	return mux
}
