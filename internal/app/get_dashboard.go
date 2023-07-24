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

func (server *Server) getReport(ctx *gin.Context) {
	user := ctx.MustGet("user").(*datastruct.UserInfo)
	dashboard, err := server.DashboardService.GetDashboard(ctx, user.Email)
	if err != nil {
		ctx.JSON(http.StatusNotFound, utils.ErrorResponse(DashboardError))
	}

	ctx.JSON(http.StatusOK, dashboard)
}
