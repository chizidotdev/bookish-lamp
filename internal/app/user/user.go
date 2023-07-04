package user

import (
	"github.com/chizidotdev/copia/internal/repository"
	"github.com/gin-gonic/gin"
)

type User interface {
	CreateUser(ctx *gin.Context)
	Login(ctx *gin.Context)
	Logout(ctx *gin.Context)
	ValidateToken(ctx *gin.Context)
	ListUsers(ctx *gin.Context)
	IsAuthenticated(ctx *gin.Context)
	//UpdateUser(ctx *gin.Context)
	//DeleteUser(ctx *gin.Context)
}

type userService struct {
	Store *repository.Store
}

func NewUserService(store *repository.Store) User {
	return &userService{
		Store: store,
	}
}
