package api

import (
	"time"

	db "github.com/chizidotdev/copia/db/sqlc"
	"github.com/gin-gonic/gin"

	"github.com/gin-contrib/cors"
)

// Server serves HTTP requests for the service
type Server struct {
	store  *db.Store
	router *gin.Engine
}

// NewServer creates a new HTTP server and setup routing
func NewServer(store *db.Store) *Server {
	server := &Server{store: store}
	router := gin.Default()

	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000"},
		AllowMethods:     []string{"PUT", "POST", "GET", "DELETE"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	// User Authentication
	router.POST("/signup", server.signup)
	router.POST("/login", server.login)
	router.GET("/logout", server.logout)
	router.GET("/users", server.listUsers)
	router.GET("/validateToken", server.isAuthenticated, server.validateToken)

	// Items
	items := router.Group("/items")
	items.Use(server.isAuthenticated)
	{
		items.POST("/", server.createItem)
		items.GET("/", server.listItems)
		items.GET("/:id", server.getItem)
		items.PUT("/:id", server.updateItem)
		items.DELETE("/:id", server.deleteItem)

		// Sales
		sales := items.Group("/:id/sales")
		{
			sales.POST("/", server.createSale)
			sales.GET("/", server.listSales)
			sales.GET("/:saleID", server.getSale)
			sales.PUT("/:saleID", server.updateSale)
			sales.DELETE("/:saleID", server.deleteSale)
		}
	}

	// Inventory
	router.GET("/inventory", server.isAuthenticated, server.getInventoryStats)

	server.router = router
	return server
}

// Start runs the HTTP server on a specific address.
func (server *Server) Start(address string) error {
	return server.router.Run(address)
}
