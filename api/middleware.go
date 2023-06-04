package api

import (
	"net/http"
	"time"

	"github.com/chizidotdev/copia/utils"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

type userJWT struct {
	Email string    `json:"email"`
	ID    uuid.UUID `json:"id"`
}

func (server *Server) isAuthenticated(ctx *gin.Context) {
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

		user, err := server.store.GetUser(ctx, claims["sub"].(string))
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, utils.ErrorResponse(err.Error()))
			return
		}

		ctx.Set("user", userJWT{
			Email: user.Email,
			ID:    user.ID,
		})
		ctx.Next()
	} else {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, utils.ErrorResponse(err.Error()))
	}
}
