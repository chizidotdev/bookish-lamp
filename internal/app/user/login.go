package user

import (
	"github.com/chizidotdev/copia/internal/dto"
	"github.com/chizidotdev/copia/pkg/utils"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"time"
)

func (u *userService) Login(ctx *gin.Context) {
	var req dto.LoginRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, utils.ErrorResponse(err.Error()))
		return
	}

	user, err := u.Store.GetUser(ctx, req.Email)
	if err != nil {
		ctx.JSON(http.StatusNotFound, utils.ErrorResponse(err.Error()))
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password))
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, utils.ErrorResponse(err.Error()))
		return
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": user.Email,
		"exp": time.Now().Add(time.Hour * 24).Unix(), // 1 day
	})

	tokenString, err := token.SignedString([]byte(utils.EnvVars.AuthSecret))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, utils.ErrorResponse(err.Error()))
		return
	}

	ctx.Header("Access-Control-Allow-Credentials", "true")
	ctx.SetSameSite(http.SameSiteNoneMode)
	ctx.SetCookie("Authorization", tokenString, 3600*24, "/", "localhost", false, true)

	ctx.JSON(http.StatusOK, gin.H{"token": tokenString})
}
