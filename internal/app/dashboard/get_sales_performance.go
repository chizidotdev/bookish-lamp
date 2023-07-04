package dashboard

import (
	"fmt"
	"github.com/chizidotdev/copia/internal/dto"
	"github.com/chizidotdev/copia/pkg/utils"
	"github.com/gin-gonic/gin"
	"math"
)

func (d *dashboardService) getSalesPerformance(ctx *gin.Context) (float64, error) {
	user := ctx.MustGet("user").(dto.UserJWT)
	currWeekSale, err := d.Store.CurrentWeekSales(ctx, user.ID)
	if err != nil {
		errMsg := fmt.Errorf("failed to get current week sales: %w", err)
		return 0, errMsg
	}

	lastWeekSales, err := d.Store.LastWeekSales(ctx, user.ID)
	if err != nil {
		errMsg := fmt.Errorf("failed to get last week sales: %w", err)
		return 0, errMsg
	}

	var salesPerformance float64
	if lastWeekSales == 0 {
		if currWeekSale == 0 {
			salesPerformance = 0
		} else {
			salesPerformance = 100
		}
	} else {
		diff := utils.CalcPercentageDiff(lastWeekSales, currWeekSale)
		salesPerformance = math.Floor(diff)
	}

	return salesPerformance, nil
}
