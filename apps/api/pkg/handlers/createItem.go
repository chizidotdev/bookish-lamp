package handlers

import (
	"encoding/json"
	"net/http"
)

func createItem(w http.ResponseWriter, r *http.Request) {
	var item Item
	err := json.NewDecoder(r.Body).Decode(&item)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	items = append(items, item)

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(item)
}
