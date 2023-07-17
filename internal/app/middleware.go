package app

import (
	"encoding/json"
	"github.com/chizidotdev/copia/internal/datastruct"
	"github.com/chizidotdev/copia/pkg/utils"
	"github.com/gin-gonic/gin"
	"io"
	"net/http"
	"strings"
)

func (server *Server) isAuth(ctx *gin.Context) {
	reqToken := ctx.Request.Header.Get("Authorization")
	if reqToken == "" {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, utils.ErrorResponse("Invalid Auth token"))
		return
	}

	_, err := server.TokenManager.ValidateToken(ctx, reqToken)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, utils.ErrorResponse(err.Error()))
		return
	}

	url := "https://" + utils.EnvVars.Auth0Domain + "/userinfo"
	request, _ := http.NewRequest("GET", url, nil)

	request.Header.Add("authorization", reqToken)
	res, _ := http.DefaultClient.Do(request)
	buf := new(strings.Builder)
	_, err = io.Copy(buf, res.Body)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, utils.ErrorResponse("Unable to get user info"))
		return
	}

	user := &datastruct.UserInfo{}
	data := []byte(buf.String())
	_ = json.Unmarshal(data, user)

	ctx.Set("user", user)
	ctx.Next()
}
