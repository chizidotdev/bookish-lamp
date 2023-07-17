package service

import (
	"github.com/chizidotdev/copia/internal/repository"
	"github.com/chizidotdev/copia/pkg/utils"
)

type Service struct {
	DashboardService
	ItemService
	SaleService
	AuthService
	TokenManager
	UserService
}

func NewService(store *repository.Store) *Service {
	dashboard := NewDashboardService(store)
	item := NewItemService(store)
	sale := NewSaleService(store)
	tokenManger := NewTokenManager(utils.EnvVars.AuthSecret)
	auth := NewAuthService(store, tokenManger)
	user := NewUserService(store, tokenManger)

	return &Service{
		DashboardService: dashboard,
		ItemService:      item,
		SaleService:      sale,
		AuthService:      auth,
		TokenManager:     tokenManger,
		UserService:      user,
	}
}
