package handlers

import (
	"copia/api/apps/api/pkg/models"
	"encoding/json"
	"net/http"
	"github.com/go-chi/chi/v5"
)

func (h handler) GetSalesHistory(w http.ResponseWriter, r *http.Request) {
    itemID := chi.URLParam(r, "id")

    var item models.Item
	if err := h.DB.First(&item, "id = ?", itemID).Error; err != nil {
        http.Error(w, err.Error(), http.StatusNotFound)
        return
    }

	var sales []models.Sale
    if err := h.DB.Where("item_id = ?", item.ID).Find(&sales).Error; err != nil {
        http.Error(w, err.Error(), http.StatusNotFound)
        return
    }

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(sales)
}
