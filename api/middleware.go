package api

import (
	"fmt"
	"net/http"
	"time"

	"github.com/chizidotdev/copia/utils"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func (server *Server) isAuthenticated(ctx *gin.Context) {
	cookie, err := ctx.Cookie("Authorization")
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, utils.ErrorResponse(err))
		return
	}

	token, err := jwt.Parse(cookie, func(token *jwt.Token) (interface{}, error) {
		return []byte(utils.EnvVars.AuthSecret), nil
	})
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, utils.ErrorResponse(err))
		return
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		if float64(time.Now().Unix()) > claims["exp"].(float64) {
			ctx.JSON(http.StatusUnauthorized, utils.ErrorResponse(err))
			return
		}
		fmt.Println("user", claims["sub"].(string))

		user, err := server.store.GetUser(ctx, claims["sub"].(string))
		if err != nil {
			ctx.JSON(http.StatusUnauthorized, utils.ErrorResponse(err))
			return
		}

		ctx.Set("user", gin.H{
			"Email": user.Email,
			"ID":    user.ID,
		})
		ctx.Next()
	} else {
		ctx.JSON(http.StatusUnauthorized, utils.ErrorResponse(err))
	}
}
