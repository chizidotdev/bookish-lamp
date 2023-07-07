package item

import (
	"github.com/chizidotdev/copia/internal/datastruct"
	"github.com/chizidotdev/copia/internal/dto"
	"github.com/chizidotdev/copia/internal/repository"
	"github.com/chizidotdev/copia/pkg/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (i *itemService) CreateItem(ctx *gin.Context) {
	var req dto.CreateItemRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, utils.ErrorResponse(err.Error()))
		return
	}

	user := ctx.MustGet("user").(datastruct.UserJWT)

	arg := repository.CreateItemParams{
		Title:        req.Title,
		BuyingPrice:  req.BuyingPrice,
		SellingPrice: req.SellingPrice,
		Quantity:     req.Quantity,
		UserID:       user.ID,
	}

	newItem, err := i.Store.CreateItem(ctx, arg)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, utils.ErrorResponse(err.Error()))
		return
	}

	ctx.JSON(http.StatusOK, newItem)
}
