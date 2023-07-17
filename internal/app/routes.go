package app

import (
	"github.com/gin-gonic/gin"
)

// createRoutes creates all the routes for the server
func createRoutes(server *Server) {
	// Create root ping route
	server.router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello from Copia!",
		})
	})

	// User/Auth routes
	server.router.POST("/signup", server.createUser)
	server.router.POST("/login", server.login)
	server.router.GET("/logout", server.logout)
	//server.router.GET("/validateToken", server.validateToken)
	server.router.GET("/user", server.getUser)

	// ItemService routes
	itemRoutes := server.router.Group("/items")
	itemRoutes.Use(server.isAuth)
	{
		itemRoutes.POST("", server.createItem)
		itemRoutes.GET("", server.listItems)
		itemRoutes.GET("/:id", server.getItemByID)
		itemRoutes.PUT("/:id", server.updateItem)
		itemRoutes.DELETE("/:id", server.deleteItem)
	}

	// SaleService routes
	saleRoutes := server.router.Group("/sales")
	saleRoutes.Use(server.isAuth)
	{
		saleRoutes.POST("", server.createSale)
		saleRoutes.GET("", server.listSales)
		saleRoutes.GET("/:saleID", server.getSaleByID)
		saleRoutes.PUT("/:saleID", server.updateSale)
		saleRoutes.DELETE("/:saleID", server.deleteSale)
	}

	// DashboardService routes
	server.router.GET("/inventory", server.isAuth, server.getDashboard)
}
