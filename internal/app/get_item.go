package app

import (
	"github.com/chizidotdev/copia/pkg/utils"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"net/http"
)

func (server *Server) getItemByID(ctx *gin.Context) {
	idParam := ctx.Params.ByName("id")
	itemID, err := uuid.Parse(idParam)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, utils.ErrorResponse(err.Error()))
		return
	}

	item, err := server.ItemService.GetItemByID(ctx, itemID)
	if err != nil {
		ctx.JSON(http.StatusNotFound, utils.ErrorResponse(err.Error()))
		return
	}

	ctx.JSON(http.StatusOK, item)
}
