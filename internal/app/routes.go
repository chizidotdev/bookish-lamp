package app

import (
	"github.com/chizidotdev/copia/internal/app/dashboard"
	"github.com/chizidotdev/copia/internal/app/item"
	"github.com/chizidotdev/copia/internal/app/sale"
	"github.com/chizidotdev/copia/internal/app/user"
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

	// User routes
	userService := user.NewUserService(server.Store)
	server.router.POST("/signup", userService.CreateUser)
	server.router.POST("/login", userService.Login)
	server.router.GET("/logout", userService.Logout)
	server.router.GET("/users", userService.IsAuthenticated, userService.ListUsers)
	server.router.GET("/validateToken", userService.IsAuthenticated, userService.ValidateToken)

	// Item routes
	itemService := item.NewItemService(server.Store)
	itemRoutes := server.router.Group("/items")
	itemRoutes.Use(userService.IsAuthenticated)
	{
		itemRoutes.POST("", itemService.CreateItem)
		itemRoutes.GET("", itemService.ListItems)
		itemRoutes.GET("/:id", itemService.GetItemByID)
		itemRoutes.PUT("/:id", itemService.UpdateItem)
		itemRoutes.DELETE("/:id", itemService.DeleteItem)
	}

	// Sale routes
	saleService := sale.NewSaleService(server.Store)
	saleRoutes := server.router.Group("/sales")
	saleRoutes.Use(userService.IsAuthenticated)
	{
		saleRoutes.POST("", saleService.CreateSale)
		saleRoutes.GET("", saleService.ListSales)
		saleRoutes.GET("/:saleID", saleService.GetSaleByID)
		saleRoutes.PUT("/:saleID", saleService.UpdateSale)
		saleRoutes.DELETE("/:saleID", saleService.DeleteSale)
	}

	// Dashboard routes
	dashboardService := dashboard.NewDashboardService(server.Store)
	server.router.GET("/inventory", userService.IsAuthenticated, dashboardService.GetDashboard)
}
