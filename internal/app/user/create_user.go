package user

import (
	"fmt"
	"github.com/chizidotdev/copia/internal/dto"
	"github.com/chizidotdev/copia/internal/repository"
	"github.com/chizidotdev/copia/pkg/utils"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"net/http"
)

func (u *userService) CreateUser(ctx *gin.Context) {
	var req dto.CreateUserRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, utils.ErrorResponse(err.Error()))
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), 10)
	if err != nil {
		errMsg := fmt.Errorf("error hashing password: %w", err)
		ctx.JSON(http.StatusInternalServerError, utils.ErrorResponse(errMsg.Error()))
		return
	}
	arg := repository.CreateUserParams{
		Email:    req.Email,
		Password: string(hashedPassword),
	}

	_, err = u.Store.CreateUser(ctx, arg)
	if err != nil {
		errMsg := fmt.Errorf("error creating user: %w", err)
		ctx.JSON(http.StatusInternalServerError, utils.ErrorResponse(errMsg.Error()))
		return
	}

	ctx.JSON(http.StatusOK, "User created successfully")
}
