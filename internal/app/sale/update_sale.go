package sale

import (
	"fmt"
	"github.com/chizidotdev/copia/internal/datastruct"
	"github.com/chizidotdev/copia/internal/dto"
	"github.com/chizidotdev/copia/internal/repository"
	"github.com/chizidotdev/copia/pkg/utils"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"net/http"
)

func (s *saleService) UpdateSale(ctx *gin.Context) {
	var req dto.UpdateSaleRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, utils.ErrorResponse(err.Error()))
		return
	}

	user := ctx.MustGet("user").(datastruct.UserJWT)
	saleID := uuid.MustParse(ctx.Param("saleID"))

	initialSale, err := s.Store.GetSale(ctx, repository.GetSaleParams{
		ID:     saleID,
		UserID: user.ID,
	})
	if err != nil {
		errMessage := fmt.Errorf("sale not found: %v", err)
		ctx.JSON(http.StatusNotFound, utils.ErrorResponse(errMessage.Error()))
		return
	}

	quantityDiff := req.QuantitySold - initialSale.QuantitySold

	args := repository.UpdateSaleParams{
		ID:           saleID,
		QuantitySold: req.QuantitySold,
		SalePrice:    req.SalePrice,
		CustomerName: req.CustomerName,
		SaleDate:     req.SaleDate,
		UserID:       user.ID,
	}

	itemArgs := repository.UpdateItemQuantityParams{
		ID:       initialSale.ItemID,
		Quantity: -quantityDiff,
		UserID:   user.ID,
	}

	var sale repository.Sale
	err = s.Store.ExecTx(ctx, func(query *repository.Queries) error {
		var err error
		sale, err = query.UpdateSale(ctx, args)
		if err != nil {
			return fmt.Errorf("error updating sale: %w", err)
		}
		_, err = query.UpdateItemQuantity(ctx, itemArgs)
		if err != nil {
			return fmt.Errorf("error updating item quantity: %w", err)
		}

		return err
	})

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, utils.ErrorResponse(err.Error()))
		return
	}

	ctx.JSON(http.StatusOK, sale)
}
