package app

import (
	"github.com/chizidotdev/copia/internal/dto"
	"github.com/chizidotdev/copia/pkg/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (server *Server) login(ctx *gin.Context) {
	var req dto.LoginRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, utils.ErrorResponse(err.Error()))
		return
	}

	tokenString, err := server.AuthService.Login(ctx, req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, utils.ErrorResponse(err.Error()))
		return
	}

	ctx.SetSameSite(http.SameSiteNoneMode)
	ctx.SetCookie("Authorization", tokenString, 3600*24, "/", utils.EnvVars.ClientDomain, true, true)

	ctx.JSON(http.StatusOK, gin.H{"token": tokenString})
}
