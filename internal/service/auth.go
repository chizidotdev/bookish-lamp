package service

import (
	"context"
	"errors"
	"fmt"
	"github.com/chizidotdev/copia/internal/dto"
	"github.com/chizidotdev/copia/internal/repository/sqlx"
	"github.com/chizidotdev/copia/pkg/utils"
	"golang.org/x/crypto/bcrypt"
)

type AuthService interface {
	CreateUser(ctx context.Context, req sqlx.CreateUserParams) error
	Login(ctx context.Context, request dto.LoginRequest) (string, error)
}

type authService struct {
	Store        *sqlx.Store
	tokenManager TokenManager
}

func NewAuthService(store *sqlx.Store, tokenManger TokenManager) AuthService {
	return &authService{
		Store:        store,
		tokenManager: tokenManger,
	}
}

func (a *authService) CreateUser(ctx context.Context, req sqlx.CreateUserParams) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), 10)
	if err != nil {
		return fmt.Errorf("error hashing password: %w", err)
	}

	req.Password = string(hashedPassword)
	_, err = a.Store.CreateUser(ctx, req)
	if err != nil {
		return errors.New(utils.ErrorMessages.SignUpError)
	}

	return nil
}

func (a *authService) Login(ctx context.Context, req dto.LoginRequest) (string, error) {
	user, err := a.Store.GetUser(ctx, req.Email)
	if err != nil {
		return "", errors.New(utils.ErrorMessages.LoginError)
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password))
	if err != nil {
		return "", errors.New(utils.ErrorMessages.LoginError)
	}

	token, err := a.tokenManager.NewJWT(user.Email)
	if err != nil {
		return "", errors.New(utils.ErrorMessages.LoginError)
	}

	return token, nil
}
