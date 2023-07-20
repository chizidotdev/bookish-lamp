package service

import (
	"context"
	"errors"
	"fmt"
	"github.com/chizidotdev/copia/internal/repository/sqlx"
	"github.com/chizidotdev/copia/pkg/utils"
)

type UserService interface {
	GetUserByEmail(ctx context.Context, email string) (sqlx.User, error)
	ListUsers(ctx context.Context) ([]sqlx.User, error)
	//UpdateUser(ctx *gin.Context)
	//DeleteUser(ctx *gin.Context)
}

type userService struct {
	Store        *sqlx.Store
	tokenManager TokenManager
}

func NewUserService(store *sqlx.Store, tokenManger TokenManager) UserService {
	return &userService{
		Store:        store,
		tokenManager: tokenManger,
	}
}

func (u *userService) GetUserByEmail(ctx context.Context, email string) (sqlx.User, error) {
	user, err := u.Store.GetUser(ctx, email)
	if err != nil {
		return sqlx.User{}, errors.New(utils.ErrorMessages.UserNotFound)
	}

	return user, nil
}

func (u *userService) ListUsers(ctx context.Context) ([]sqlx.User, error) {
	users, err := u.Store.ListUsers(ctx)
	if err != nil {
		return nil, fmt.Errorf("error listing users: %w", err)
	}

	return users, nil
}
