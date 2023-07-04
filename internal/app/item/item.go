package item

import (
	"github.com/chizidotdev/copia/internal/repository"
	"github.com/gin-gonic/gin"
)

type Item interface {
	CreateItem(ctx *gin.Context)
	GetItemByID(ctx *gin.Context)
	ListItems(ctx *gin.Context)
	UpdateItem(ctx *gin.Context)
	DeleteItem(ctx *gin.Context)
}

type itemService struct {
	Store *repository.Store
}

func NewItemService(store *repository.Store) Item {
	return &itemService{
		Store: store,
	}
}
