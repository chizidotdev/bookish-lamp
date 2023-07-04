package item

import (
	"github.com/chizidotdev/copia/internal/dto"
	"github.com/chizidotdev/copia/pkg/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (i *itemService) ListItems(ctx *gin.Context) {
	user := ctx.MustGet("user").(dto.UserJWT)

	items, err := i.Store.ListItems(ctx, user.ID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, utils.ErrorResponse(err.Error()))
		return
	}

	ctx.JSON(http.StatusOK, items)
}
