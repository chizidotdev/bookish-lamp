package api

import (
	"encoding/gob"
	"log"

	"github.com/chizidotdev/copia/auth"
	db "github.com/chizidotdev/copia/db/sqlc"
	"github.com/chizidotdev/copia/middleware"
	"github.com/gin-gonic/gin"

	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
)

// Server serves HTTP requests for the service
type Server struct {
	store  db.Store
	router *gin.Engine
}

// NewServer creates a new HTTP server and setup routing
func NewServer(store db.Store) *Server {
	server := &Server{store: store}
	router := gin.Default()

	auth, err := auth.New()
	if err != nil {
		log.Fatal("Cannot create authenticator:", err)
	}

	gob.Register(map[string]interface{}{})
	cookieStore := cookie.NewStore([]byte("secret"))
	router.Use(sessions.Sessions("auth-session", cookieStore))
	// config := cors.DefaultConfig()
	// config.AllowOrigins = []string{"http://google.com"}
	router.Use(cors.Default())

	router.GET("/login", server.login(auth))
	router.GET("/callback", server.callback(auth))
	router.GET("/logout", server.logout)
	router.GET("/user", server.getUser)

	router.POST("/items", server.createItem)
	router.GET("/items", middleware.IsAuthenticated, server.listItems)
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
