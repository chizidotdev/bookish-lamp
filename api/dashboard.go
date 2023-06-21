package api

import (
	"fmt"
	"math"
	"net/http"

	"github.com/chizidotdev/copia/utils"
	"github.com/gin-gonic/gin"
)

type DashboardInfo struct {
	TotalItems       int64   `json:"total_items"`
	LowStockItems    int64   `json:"low_stock_items"`
	RecentSales      int64   `json:"recent_sales"`
	PendingOrders    int64   `json:"pending_orders"`
	SalesPerformance float64 `json:"sales_performance"`
}

func (server *Server) getDashboard(ctx *gin.Context) {
	var dashboard DashboardInfo

	user := ctx.MustGet("user").(userJWT)
	inventory, err := server.store.GetInventoryStats(ctx, user.ID)
	if err != nil {
		errMessage := fmt.Errorf("failed to get inventory stats: %w", err)
		ctx.JSON(http.StatusNotFound, utils.ErrorResponse(errMessage.Error()))
		return
	}

	dashboard.TotalItems = inventory.TotalItems
	dashboard.LowStockItems = inventory.LowStockItems
	dashboard.RecentSales = inventory.RecentSales
	dashboard.PendingOrders = inventory.PendingOrders

	currWeekSale, err := server.store.CurrentWeekSales(ctx, user.ID)
	if err != nil {
		errMessage := fmt.Errorf("failed to get current week sales: %w", err)
		ctx.JSON(http.StatusInternalServerError, utils.ErrorResponse(errMessage.Error()))
		return
	}

	lastWeekSales, err := server.store.LastWeekSales(ctx, user.ID)
	if err != nil {
		errMessage := fmt.Errorf("failed to get last week sales: %w", err)
		ctx.JSON(http.StatusInternalServerError, utils.ErrorResponse(errMessage.Error()))
		return
	}

	if lastWeekSales == 0 {
		if currWeekSale == 0 {
			dashboard.SalesPerformance = 0
		} else {
			dashboard.SalesPerformance = 100
		}
	} else {
		salesPerformance := utils.CalcPercentageDiff(lastWeekSales, currWeekSale)
		dashboard.SalesPerformance = math.Floor(salesPerformance)
	}

	ctx.JSON(http.StatusOK, dashboard)
}
