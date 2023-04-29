package handlers

import (
	"copia/api/apps/api/pkg/models"
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi"
)

func (h handler) GetItem(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	var item models.Item

	if err := h.DB.First(&item, id).Error; err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(item)
}
