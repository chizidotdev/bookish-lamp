package handlers

import (
	"encoding/json"
	"net/http"
	"github.com/go-chi/chi"
)

func (h handler) DeleteItem(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	for i, item := range items {
		if id == item.ID {
			items = append(items[:i], items[i+1:]...)

			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode("Item deleted successfully")
			return
		}
	}

	http.NotFound(w, r)
}
