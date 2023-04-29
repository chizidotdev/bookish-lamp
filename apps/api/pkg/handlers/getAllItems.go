package handlers

import (
	"copia/api/apps/api/pkg/models"
	"encoding/json"
	"net/http"
)

func (h handler) GetAllItems(w http.ResponseWriter, r *http.Request) {
	var items []models.Item

	err := h.DB.Order("created_at DESC, id DESC").Find(&items).Error
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(items)
}
