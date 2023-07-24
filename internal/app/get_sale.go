package app

import (
	"github.com/chizidotdev/copia/internal/datastruct"
	"github.com/chizidotdev/copia/pkg/utils"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"net/http"
)

func (server *Server) getSaleByID(ctx *gin.Context) {
	user := ctx.MustGet("user").(*datastruct.UserInfo)
	saleID := uuid.MustParse(ctx.Param("saleID"))

	sale, err := server.SaleService.GetSaleByID(ctx, datastruct.GetSaleParams{
		ID:        saleID,
		UserEmail: user.Email,
	})
	if err != nil {
		ctx.JSON(http.StatusNotFound, utils.ErrorResponse(err.Error()))
		return
	}

	ctx.JSON(http.StatusOK, sale)
}
