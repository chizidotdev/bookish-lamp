package app

import (
	"github.com/chizidotdev/copia/internal/datastruct"
	"github.com/chizidotdev/copia/pkg/utils"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"net/http"
)

func (server *Server) listSales(ctx *gin.Context) {
	user := ctx.MustGet("user").(*datastruct.UserInfo)

	itemID, err := uuid.Parse(ctx.Query("itemID"))
	if err != nil {
		sales, err := server.SaleService.ListSalesByUser(ctx, user.Email)
		if err != nil {
			ctx.JSON(http.StatusNotFound, utils.ErrorResponse(err.Error()))
			return
		}

		ctx.JSON(http.StatusOK, sales)
	} else {
		sales, err := server.SaleService.ListSalesByItem(ctx, datastruct.ListSalesParams{
			ItemID:    itemID,
			UserEmail: user.Email,
		})
		if err != nil {
			ctx.JSON(http.StatusNotFound, utils.ErrorResponse(err.Error()))
			return
		}

		ctx.JSON(http.StatusOK, sales)
	}
}
