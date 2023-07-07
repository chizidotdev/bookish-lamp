package user

import (
	"github.com/chizidotdev/copia/internal/datastruct"
	"github.com/chizidotdev/copia/pkg/utils"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"net/http"
	"time"
)

func (u *userService) IsAuthenticated(ctx *gin.Context) {
	cookie, err := ctx.Cookie("Authorization")

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, utils.ErrorResponse(err.Error()))
		return
	}

	token, err := jwt.Parse(cookie, func(token *jwt.Token) (interface{}, error) {
		return []byte(utils.EnvVars.AuthSecret), nil
	})
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, utils.ErrorResponse(err.Error()))
		return
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		if float64(time.Now().Unix()) > claims["exp"].(float64) {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, utils.ErrorResponse(err.Error()))
			return
		}

		user, err := u.Store.GetUser(ctx, claims["sub"].(string))
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, utils.ErrorResponse(err.Error()))
			return
		}

		ctx.Set("user", datastruct.UserJWT{
			Email: user.Email,
			ID:    user.ID,
		})
		ctx.Next()
	} else {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, utils.ErrorResponse(err.Error()))
	}
}
