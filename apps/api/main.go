package main

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type Item struct {
	ID           int     `json:"id"`
	Name         string  `json:"name"`
	BuyingPrice  float64 `json:"buying_price"`
	SellingPrice float64 `json:"selling_price"`
}

var items []Item

func main() {
	items = []Item{
		{ID: 1, Name: "Perfume 1", BuyingPrice: 100, SellingPrice: 200},
	}

	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Use(middleware.Timeout(60 * time.Second))

	r.Route("/api/v1/items", func(r chi.Router) {
		r.Get("/", getItems)
		r.Post("/", createItem)

		r.Route("/{id}", func(r chi.Router) {
			r.Get("/", getItemById)
			// r.Put("/", updateItem)
			// r.Delete("/", deleteItem)
		})
	})

	err := http.ListenAndServe(":3333", r)
	log.Fatal(err)
}

func getItems(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	jsonBytes, err := json.Marshal(items)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	_, _ = w.Write(jsonBytes)
}

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

func getItemById(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	for _, item := range items {
		if id == string(item.ID) {
			json.NewEncoder(w).Encode(item)
			return
		}
	}

	w.WriteHeader(http.StatusNotFound)
}
