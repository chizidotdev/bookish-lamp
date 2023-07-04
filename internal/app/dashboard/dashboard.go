package dashboard

import (
	"github.com/chizidotdev/copia/internal/repository"
	"github.com/gin-gonic/gin"
)

type Dashboard interface {
	GetDashboard(ctx *gin.Context)
}

type dashboardService struct {
	Store *repository.Store
}

func NewDashboardService(store *repository.Store) Dashboard {
	return &dashboardService{
		Store: store,
	}
}
