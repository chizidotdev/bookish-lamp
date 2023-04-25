package main

import (
	"encoding/json"
	"log"
	"net/http"

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

	r.Use(middleware.Logger)

	r.Get("/items", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		jsonBytes, err := json.Marshal(items)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		_, _ = w.Write(jsonBytes)
	})

	err := http.ListenAndServe(":8000", r)
	log.Fatal(err)
}
