package app

import (
	"github.com/chizidotdev/copia/internal/datastruct"
	"github.com/chizidotdev/copia/pkg/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (server *Server) isAuth(ctx *gin.Context) {
	cookie, err := ctx.Cookie("Authorization")

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, utils.ErrorResponse(err.Error()))
		return
	}

	subject, err := server.TokenManager.Parse(cookie)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, utils.ErrorResponse(err.Error()))
		return
	}

	user, err := server.AuthService.GetUser(ctx, subject)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, utils.ErrorResponse(err.Error()))
		return
	}

	ctx.Set("user", datastruct.UserJWT{
		Email: user.Email,
		ID:    user.ID,
	})
	ctx.Next()
}
