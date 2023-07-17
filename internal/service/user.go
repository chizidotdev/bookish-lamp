package service

import (
	"context"
	"errors"
	"fmt"
	"github.com/chizidotdev/copia/internal/repository"
	"github.com/chizidotdev/copia/pkg/utils"
)

type UserService interface {
	GetUserByEmail(ctx context.Context, email string) (repository.User, error)
	ListUsers(ctx context.Context) ([]repository.User, error)
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

func (u *userService) GetUserByEmail(ctx context.Context, email string) (repository.User, error) {
	user, err := u.Store.GetUser(ctx, email)
	if err != nil {
		return repository.User{}, errors.New(utils.ErrorMessages.UserNotFound)
	}

	return user, nil
}

func (u *userService) ListUsers(ctx context.Context) ([]repository.User, error) {
	users, err := u.Store.ListUsers(ctx)
	if err != nil {
		return nil, fmt.Errorf("error listing users: %w", err)
	}

	return users, nil
}
