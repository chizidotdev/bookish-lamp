package main

import (
	"log"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"

	_ "github.com/lib/pq"

	"copia/api/apps/api/pkg/db"
	"copia/api/apps/api/pkg/handlers"
)

func init() {
	if err := godotenv.Load("../../.env"); err != nil {
		log.Fatal(err)
	}
}

func main() {
	DB := db.Init()
	h := handlers.New(DB)
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
		r.Get("/", h.GetAllItems)
		r.Post("/", h.CreateItem)

		r.Route("/{id}", func(r chi.Router) {
			// r.Get("/", getItemById)
			r.Put("/", h.UpdateItem)
			r.Delete("/", h.DeleteItem)
		})
	})

	r.Route("/api/v1/sales", func(r chi.Router) {
		r.Get("/", h.GetAllSales)
		// r.Post("/", createSale)
	})

	log.Println("Server running on port 3333")
	err := http.ListenAndServe(":3333", r)
	log.Fatal(err)
}
