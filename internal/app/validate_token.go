package app

import (
	"github.com/chizidotdev/copia/pkg/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (server *Server) validateToken(ctx *gin.Context) {
	user, exists := ctx.Get("user")
	if !exists {
		ctx.JSON(http.StatusUnauthorized, utils.ErrorResponse("Unauthorized"))
		return
	}

	ctx.JSON(http.StatusOK, user)
}
