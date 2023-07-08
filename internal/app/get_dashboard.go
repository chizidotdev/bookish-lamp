package app

import (
	"github.com/chizidotdev/copia/internal/datastruct"
	"github.com/chizidotdev/copia/pkg/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

const (
	DASHBOARD_ERROR = "An error occurred while getting dashboard"
)

func (server *Server) getDashboard(ctx *gin.Context) {
	user := ctx.MustGet("user").(datastruct.UserJWT)
	dashboard, err := server.DashboardService.GetDashboard(ctx, user)
	if err != nil {
		ctx.JSON(http.StatusNotFound, utils.ErrorResponse(DASHBOARD_ERROR))
	}

	ctx.JSON(http.StatusOK, dashboard)
}
