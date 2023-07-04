package sale

import (
	"github.com/chizidotdev/copia/internal/repository"
	"github.com/gin-gonic/gin"
)

type Sale interface {
	CreateSale(ctx *gin.Context)
	GetSaleByID(ctx *gin.Context)
	ListSalesByUserID(ctx *gin.Context)
	ListSales(ctx *gin.Context)
	UpdateSale(ctx *gin.Context)
	DeleteSale(ctx *gin.Context)
}

type saleService struct {
	Store *repository.Store
}

func NewSaleService(store *repository.Store) Sale {
	return &saleService{
		Store: store,
	}
}
