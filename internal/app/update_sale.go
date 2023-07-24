package app

import (
	"github.com/chizidotdev/copia/internal/datastruct"
	"github.com/chizidotdev/copia/internal/dto"
	"github.com/chizidotdev/copia/pkg/utils"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"net/http"
)

func (server *Server) updateSale(ctx *gin.Context) {
	var req dto.UpdateSaleRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, utils.ErrorResponse(err.Error()))
		return
	}

	user := ctx.MustGet("user").(*datastruct.UserInfo)
	saleID := uuid.MustParse(ctx.Param("saleID"))

	sale, err := server.SaleService.UpdateSale(ctx, datastruct.UpdateSaleParams{
		ID:           saleID,
		UserEmail:    user.Email,
		QuantitySold: req.QuantitySold,
		SaleDate:     req.SaleDate,
		SalePrice:    req.SalePrice,
		CustomerName: req.CustomerName,
	})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, utils.ErrorResponse(err.Error()))
		return
	}

	ctx.JSON(http.StatusOK, sale)
}
