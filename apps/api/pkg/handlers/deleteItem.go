package handlers

import (
	"copia/api/apps/api/pkg/models"
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func (h handler) DeleteItem(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	var item models.Item
	if err := h.DB.First(&item, "id = ?", id).Error; err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	h.DB.Delete(&item)

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode("Item deleted successfully")
}
