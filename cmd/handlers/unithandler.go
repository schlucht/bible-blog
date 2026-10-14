package handlers

import (
	"net/http"

	"github.com/schlucht/bibleblog/pkg/repositories"
)

func (h *Handler) AllUnits(w http.ResponseWriter, r *http.Request) {
	repo := repositories.NewRepo(h.App.DB)
	// err = repo.NewUnit()
	units, err := repo.AllUnits()
	if err != nil {
		h.App.ErrorLog.Println("Fehler in der DB")
	}
	for _, unit := range units {
		h.App.InfoLog.Println(unit.UnitName)
	}
}
