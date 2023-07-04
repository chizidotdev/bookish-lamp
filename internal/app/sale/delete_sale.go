package sale

import (
	"fmt"
	"github.com/chizidotdev/copia/internal/dto"
	"github.com/chizidotdev/copia/internal/repository"
	"github.com/chizidotdev/copia/pkg/utils"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"net/http"
)

func (s *saleService) DeleteSale(ctx *gin.Context) {
	user := ctx.MustGet("user").(dto.UserJWT)
	saleID := uuid.MustParse(ctx.Param("saleID"))

	sale, err := s.Store.GetSale(ctx, repository.GetSaleParams{
		ID:     saleID,
		UserID: user.ID,
	})
	if err != nil {
		errMessage := fmt.Errorf("sale not found: %v", err)
		ctx.JSON(http.StatusNotFound, utils.ErrorResponse(errMessage.Error()))
		return
	}

	args := repository.DeleteSaleParams{
		ID:     saleID,
		UserID: user.ID,
	}

	itemArgs := repository.UpdateItemQuantityParams{
		ID:       sale.ItemID,
		Quantity: sale.QuantitySold,
		UserID:   user.ID,
	}

	err = s.Store.ExecTx(ctx, func(query *repository.Queries) error {
		var err error
		err = query.DeleteSale(ctx, args)
		if err != nil {
			return err
		}
		_, err = query.UpdateItemQuantity(ctx, itemArgs)
		return err
	})

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, utils.ErrorResponse(err.Error()))
		return
	}

	ctx.JSON(http.StatusOK, "Sale deleted successfully")
}
