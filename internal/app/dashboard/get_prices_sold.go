package dashboard

import (
	"fmt"
	"github.com/chizidotdev/copia/internal/datastruct"
	"github.com/chizidotdev/copia/internal/repository"
	"github.com/gin-gonic/gin"
)

func (d *dashboardService) getPriceSoldByDate(ctx *gin.Context) ([]repository.PriceSoldByDateRow, error) {
	user := ctx.MustGet("user").(datastruct.UserJWT)
	priceSoldByDate, err := d.Store.PriceSoldByDate(ctx, user.ID)
	if err != nil {
		errMsg := fmt.Errorf("failed to get price sold by date: %w", err)
		return nil, errMsg
	}

	return priceSoldByDate, nil
}

func (d *dashboardService) getPriceSoldByWeek(ctx *gin.Context) ([]repository.PriceSoldByWeekRow, error) {
	user := ctx.MustGet("user").(datastruct.UserJWT)
	priceSoldByWeek, err := d.Store.PriceSoldByWeek(ctx, user.ID)
	if err != nil {
		errMsg := fmt.Errorf("failed to get price sold by week: %w", err)
		return nil, errMsg
	}

	return priceSoldByWeek, nil
}
