package app

import (
	"github.com/chizidotdev/copia/internal/datastruct"
	"github.com/chizidotdev/copia/pkg/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

const (
	DashboardError = "An error occurred while getting dashboard"
)

func (server *Server) getDashboard(ctx *gin.Context) {
	user := ctx.MustGet("user").(datastruct.UserJWT)
	dashboard, err := server.DashboardService.GetDashboard(ctx, user.ID)
	if err != nil {
		ctx.JSON(http.StatusNotFound, utils.ErrorResponse(DashboardError))
	}

	ctx.JSON(http.StatusOK, dashboard)
}
