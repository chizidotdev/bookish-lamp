package api

import (
	"fmt"
	"net/http"

	"github.com/chizidotdev/copia/utils"
	"github.com/gin-gonic/gin"
)

func (server *Server) getInventoryStats(ctx *gin.Context) {
	user := ctx.MustGet("user").(userJWT)
	dashboard, err := server.store.GetInventoryStats(ctx, user.ID)
	if err != nil {
		errMessage := fmt.Errorf("failed to get inventory stats: %w", err)
		ctx.JSON(http.StatusNotFound, utils.ErrorResponse(errMessage.Error()))
		return
	}

	ctx.JSON(http.StatusOK, dashboard)
}
