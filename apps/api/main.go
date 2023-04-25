package main

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
)

type Item struct {
	ID           string  `json:"id"`
	Name         string  `json:"name"`
	BuyingPrice  float64 `json:"buying_price"`
	SellingPrice float64 `json:"selling_price"`
}

type Sale struct {
	ID        string  `json:"id"`
	ItemID    string  `json:"item_id"`
	Quantity  int     `json:"quantity"`
	Total     float64 `json:"total"`
	CreatedAt string  `json:"created_at"`
}

var items []Item
var sales []Sale

func main() {
	items = []Item{
		{ID: "1", Name: "Perfume 1", BuyingPrice: 100, SellingPrice: 200},
		{ID: "2", Name: "Perfume 2", BuyingPrice: 200, SellingPrice: 300},
	}

	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Use(middleware.Timeout(60 * time.Second))

	corsMiddleware := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		AllowCredentials: true,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
	})
	r.Use(corsMiddleware.Handler)

	r.Route("/api/v1/items", func(r chi.Router) {
		r.Get("/", getItems)
		r.Post("/", createItem)

		r.Route("/{id}", func(r chi.Router) {
			// r.Get("/", getItemById)
			r.Put("/", updateItem)
			r.Delete("/", deleteItem)
		})
	})

	r.Route("/api/v1/sales", func(r chi.Router) {
		r.Get("/", getSales)
		// r.Post("/", createSale)
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

func deleteItem(w http.ResponseWriter, r *http.Request) {
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

func getSales(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	jsonBytes, err := json.Marshal(sales)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	_, _ = w.Write(jsonBytes)
}
