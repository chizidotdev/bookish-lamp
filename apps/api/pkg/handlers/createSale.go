package handlers

import (
	"copia/api/apps/api/pkg/models"
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func (h handler) CreateSale(w http.ResponseWriter, r *http.Request) {
	itemID := chi.URLParam(r, "id")

	var sale models.Sale
	defer r.Body.Close()

	if err := json.NewDecoder(r.Body).Decode(&sale); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var item models.Item
	if err := h.DB.First(&item, "id = ?", itemID).Error; err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if item.Quantity == 0 {
		http.Error(w, "Item is out of stock", http.StatusBadRequest)
		return
	}

	if item.Quantity < sale.QuantitySold {
		http.Error(w, "Not enough items in stock", http.StatusBadRequest)
		return
	}

	item.Quantity -= sale.QuantitySold

	tx := h.DB.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	// Create a new sales history record
	sale.ItemID = item.ID
	if err := h.DB.Create(&sale).Error; err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		tx.Rollback()
		return
	}

	// Save the item with the new Quantity
	if err := h.DB.Save(&item).Error; err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		tx.Rollback()
		return
	}

	if err := tx.Commit().Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(sale)
}
