package sale

import (
	"github.com/chizidotdev/copia/internal/datastruct"
	"github.com/chizidotdev/copia/internal/dto"
	"github.com/chizidotdev/copia/internal/repository"
	"github.com/chizidotdev/copia/pkg/utils"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"net/http"
)

func (s *saleService) CreateSale(ctx *gin.Context) {
	var req dto.CreateSaleRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, utils.ErrorResponse(err.Error()))
		return
	}

	user := ctx.MustGet("user").(datastruct.UserJWT)
	args := repository.CreateSaleParams{
		UserID:       user.ID,
		ItemID:       uuid.MustParse(ctx.Query("itemID")),
		QuantitySold: req.QuantitySold,
		SalePrice:    req.SalePrice,
		CustomerName: req.CustomerName,
		SaleDate:     req.SaleDate,
	}

	itemArg := repository.UpdateItemQuantityParams{
		ID:       args.ItemID,
		Quantity: -args.QuantitySold,
		UserID:   user.ID,
	}

	var sale repository.Sale
	err := s.Store.ExecTx(ctx, func(query *repository.Queries) error {
		var err error
		sale, err = query.CreateSale(ctx, args)
		if err != nil {
			return err
		}
		_, err = query.UpdateItemQuantity(ctx, itemArg)
		return err
	})

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, utils.ErrorResponse(err.Error()))
		return
	}

	ctx.JSON(http.StatusOK, sale)
}
