package app

import (
	"github.com/chizidotdev/copia/pkg/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (server *Server) logout(ctx *gin.Context) {
	ctx.SetSameSite(http.SameSiteNoneMode)
	ctx.SetCookie("Authorization", "", -1, "", utils.EnvVars.ClientDomain, true, true)
	ctx.JSON(http.StatusOK, "Logged out successfully")
}
