package api

import (
	db "github.com/chizidotdev/copia/db/sqlc"
	"github.com/gin-gonic/gin"
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

	router.POST("/items", server.createItem)
	router.GET("/items", server.listItems)
	router.GET("/items/:id", server.getItem)

	router.PATCH("/items", server.updateItem)
	router.DELETE("/items/:id", server.deleteItem)

	server.router = router
	return server
}

// Start runs the HTTP server on a specific address.
func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
