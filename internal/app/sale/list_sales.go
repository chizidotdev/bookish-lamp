package sale

import (
	"github.com/chizidotdev/copia/internal/dto"
	"github.com/chizidotdev/copia/internal/repository"
	"github.com/chizidotdev/copia/pkg/utils"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"net/http"
)

func (s *saleService) ListSalesByUserID(ctx *gin.Context) {
	user := ctx.MustGet("user").(dto.UserJWT)

	sales, err := s.Store.ListSalesByUserId(ctx, user.ID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, utils.ErrorResponse(err.Error()))
		return
	}

	ctx.JSON(http.StatusOK, sales)
}

func (s *saleService) ListSales(ctx *gin.Context) {
	user := ctx.MustGet("user").(dto.UserJWT)

	itemID, err := uuid.Parse(ctx.Query("itemID"))
	if err != nil {
		var allSales []repository.ListSalesByUserIdRow

		allSales, err = s.Store.ListSalesByUserId(ctx, user.ID)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, utils.ErrorResponse(err.Error()))
			return
		}

		ctx.JSON(http.StatusOK, allSales)
	} else {
		var itemSales []repository.ListSalesRow

		_, err = s.Store.GetItem(ctx, itemID)
		if err != nil {
			ctx.JSON(http.StatusNotFound, utils.ErrorResponse("item not found"))
			return
		}
		itemSales, err = s.Store.ListSales(ctx, repository.ListSalesParams{
			ItemID: itemID,
			UserID: user.ID,
		})
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, utils.ErrorResponse(err.Error()))
			return
		}

		ctx.JSON(http.StatusOK, itemSales)
	}
}
