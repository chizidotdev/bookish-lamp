package handlers

import (
	"copia/api/apps/api/pkg/models"
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi"
)

func (h handler) UpdateItem(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	var updatedItem models.Item

	defer r.Body.Close()
	if err := json.NewDecoder(r.Body).Decode(&updatedItem); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var item models.Item
	if err := h.DB.First(&item, id).Error; err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	item = updatedItem
	if err := h.DB.Save(&item).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(item)
}
