package app

import (
	"github.com/chizidotdev/copia/internal/datastruct"
	"github.com/chizidotdev/copia/pkg/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (server *Server) listItems(ctx *gin.Context) {
	user := ctx.MustGet("user").(datastruct.UserJWT)

	items, err := server.ItemService.ListItems(ctx, user)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, utils.ErrorResponse(err.Error()))
		return
	}

	ctx.JSON(http.StatusOK, items)
}
