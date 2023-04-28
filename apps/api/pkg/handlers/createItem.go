package handlers

import (
	"copia/api/apps/api/pkg/models"
	"encoding/json"
	"net/http"
)

func (h handler) CreateItem(w http.ResponseWriter, r *http.Request) {
	var item models.Item
	defer r.Body.Close()
	if err := json.NewDecoder(r.Body).Decode(&item); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := h.DB.Create(&item).Error; err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(item)
}
