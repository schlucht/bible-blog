package routes

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/schlucht/bibleblog/cmd/application"
	"github.com/schlucht/bibleblog/cmd/handlers"
)

func Routes(app *application.Application) http.Handler {

	mux := chi.NewRouter()
	h := handlers.NewHandler(app)
	mux.Use(middleware.Logger)

	mux.Get("/", h.Home)

	mux.NotFound(h.NotFound)

	fileServer := http.FileServer(http.Dir("./assets"))
	mux.Handle("/assets/*", http.StripPrefix("/assets", fileServer))

	return mux
}
