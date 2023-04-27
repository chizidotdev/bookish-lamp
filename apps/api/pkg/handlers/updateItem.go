package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi"
)

func updateItem(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	for _, item := range items {
		if id == item.ID {
			err := json.NewDecoder(r.Body).Decode(&item)
			if err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}

			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(item)
			return
		}
	}

	http.NotFound(w, r)
}
