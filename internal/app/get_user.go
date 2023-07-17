package app

import (
	"github.com/chizidotdev/copia/pkg/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (server *Server) getUser(ctx *gin.Context) {
	email, exists := ctx.GetQuery("email")
	if !exists {
		ctx.JSON(http.StatusBadRequest, utils.ErrorResponse("Invalid email provided"))
		return
	}

	userInfo, err := server.UserService.GetUserByEmail(ctx, email)
	if err != nil {
		ctx.JSON(http.StatusNotFound, utils.ErrorResponse(err.Error()))
		return
	}

	ctx.JSON(http.StatusOK, userInfo)
}
