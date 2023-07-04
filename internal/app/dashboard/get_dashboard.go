package dashboard

import (
	"fmt"
	"github.com/chizidotdev/copia/internal/dto"
	"github.com/chizidotdev/copia/pkg/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (d *dashboardService) GetDashboard(ctx *gin.Context) {
	user := ctx.MustGet("user").(dto.UserJWT)
	inventory, err := d.Store.GetInventoryStats(ctx, user.ID)
	if err != nil {
		errMsg := fmt.Errorf("failed to get inventory stats: %w", err)
		ctx.JSON(http.StatusNotFound, utils.ErrorResponse(errMsg.Error()))
		return
	}

	salesPerformance, err := d.getSalesPerformance(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, utils.ErrorResponse(err.Error()))
		return
	}

	priceSoldByDate, err := d.getPriceSoldByDate(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, utils.ErrorResponse(err.Error()))
		return
	}

	priceSoldByWeek, err := d.getPriceSoldByWeek(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, utils.ErrorResponse(err.Error()))
		return
	}

	dashboard := dto.DashboardResponse{
		TotalItems:       inventory.TotalItems,
		LowStockItems:    inventory.LowStockItems,
		RecentSales:      inventory.RecentSales,
		PendingOrders:    inventory.PendingOrders,
		SalesPerformance: salesPerformance,
		PriceSoldByDate:  priceSoldByDate,
		PriceSoldByWeek:  priceSoldByWeek,
	}

	ctx.JSON(http.StatusOK, dashboard)
}
