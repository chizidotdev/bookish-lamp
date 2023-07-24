package service

import (
	"context"
	"github.com/chizidotdev/copia/internal/datastruct"
	"github.com/chizidotdev/copia/internal/repository"
)

type UserService interface {
	GetUserByEmail(ctx context.Context, email string) (datastruct.UserJWT, error)
	ListUsers(ctx context.Context) ([]datastruct.UserJWT, error)
	//UpdateUser(ctx *gin.Context)
	//DeleteUser(ctx *gin.Context)
}

type userService struct {
	Store        *repository.Store
	tokenManager TokenManager
}

func NewUserService(store *repository.Store, tokenManger TokenManager) UserService {
	return &userService{
		Store:        store,
		tokenManager: tokenManger,
	}
}

// GetUserByEmail returns a user by email
// TODO: Implement.
func (u *userService) GetUserByEmail(ctx context.Context, email string) (datastruct.UserJWT, error) {
	return datastruct.UserJWT{}, nil
}

// ListUsers returns a list of users
// TODO: Implement.
func (u *userService) ListUsers(ctx context.Context) ([]datastruct.UserJWT, error) {
	return nil, nil
}
