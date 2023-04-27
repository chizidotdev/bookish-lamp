package main

import (
	"log"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"

	_ "github.com/lib/pq"
)

func main() {
	// items = []Item{
	// 	{ID: "1", Name: "Perfume 1", BuyingPrice: 100, SellingPrice: 200},
	// 	{ID: "2", Name: "Perfume 2", BuyingPrice: 200, SellingPrice: 300},
	// }

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
