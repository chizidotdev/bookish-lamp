package user

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func (u *userService) Logout(ctx *gin.Context) {
	ctx.SetSameSite(http.SameSiteLaxMode)
	ctx.SetCookie("Authorization", "", -1, "", "", false, true)
	ctx.JSON(http.StatusOK, "Logged out successfully")
}
