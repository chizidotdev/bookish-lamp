package service

import (
	"context"
	"errors"
	"fmt"
	"github.com/chizidotdev/copia/internal/dto"
	"github.com/chizidotdev/copia/internal/repository"
	"github.com/chizidotdev/copia/pkg/utils"
	"golang.org/x/crypto/bcrypt"
)

type AuthService interface {
	CreateUser(ctx context.Context, req repository.CreateUserParams) error
	Login(ctx context.Context, request dto.LoginRequest) (string, error)
	GetUser(ctx context.Context, email string) (repository.User, error)
	ListUsers(ctx context.Context) ([]repository.User, error)
	//ValidateToken(ctx *gin.Context)
	//IsAuthenticated(ctx *gin.Context)
	//Logout(ctx context.Context) error
	//UpdateUser(ctx *gin.Context)
	//DeleteUser(ctx *gin.Context)
}

type authService struct {
	Store        *repository.Store
	tokenManager TokenManager
}

func NewAuthService(store *repository.Store, tokenManger TokenManager) AuthService {
	return &authService{
		Store:        store,
		tokenManager: tokenManger,
	}
}

func (a *authService) CreateUser(ctx context.Context, req repository.CreateUserParams) error {
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

func (a *authService) GetUser(ctx context.Context, email string) (repository.User, error) {
	user, err := a.Store.GetUser(ctx, email)
	if err != nil {
		return repository.User{}, fmt.Errorf("error getting user: %w", err)
	}

	return user, nil
}

func (a *authService) ListUsers(ctx context.Context) ([]repository.User, error) {
	users, err := a.Store.ListUsers(ctx)
	if err != nil {
		return nil, fmt.Errorf("error listing users: %w", err)
	}

	return users, nil
}
