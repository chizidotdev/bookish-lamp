package app

import (
	"github.com/chizidotdev/copia/internal/datastruct"
	"github.com/chizidotdev/copia/internal/dto"
	"github.com/chizidotdev/copia/pkg/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (server *Server) createItem(ctx *gin.Context) {
	var req dto.CreateItemRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, utils.ErrorResponse(err.Error()))
		return
	}

	user := ctx.MustGet("user").(*datastruct.UserInfo)

	item, err := server.ItemService.CreateItem(ctx, datastruct.CreateItemParams{
		Title:        req.Title,
		BuyingPrice:  req.BuyingPrice,
		SellingPrice: req.SellingPrice,
		Quantity:     req.Quantity,
		UserEmail:    user.Email,
	})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, utils.ErrorResponse(err.Error()))
		return
	}

	ctx.JSON(http.StatusOK, item)
}
