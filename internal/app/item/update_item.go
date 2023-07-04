package item

import (
	"github.com/chizidotdev/copia/internal/dto"
	"github.com/chizidotdev/copia/internal/repository"
	"github.com/chizidotdev/copia/pkg/utils"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"net/http"
)

func (i *itemService) UpdateItem(ctx *gin.Context) {
	var req dto.UpdateItemRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, utils.ErrorResponse(err.Error()))
		return
	}

	idParam := ctx.Params.ByName("id")
	itemID, err := uuid.Parse(idParam)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, utils.ErrorResponse(err.Error()))
		return
	}

	user := ctx.MustGet("user").(dto.UserJWT)

	arg := repository.UpdateItemParams{
		ID:           itemID,
		Title:        req.Title,
		BuyingPrice:  req.BuyingPrice,
		SellingPrice: req.SellingPrice,
		Quantity:     req.Quantity,
		UserID:       user.ID,
	}

	item, err := i.Store.UpdateItem(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, utils.ErrorResponse(err.Error()))
	}

	ctx.JSON(http.StatusOK, item)
}
