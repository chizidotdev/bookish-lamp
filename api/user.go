package api

import (
	"net/http"
	"time"

	db "github.com/chizidotdev/copia/db/sqlc"
	"github.com/chizidotdev/copia/utils"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

type userRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (server *Server) listUsers(ctx *gin.Context) {
	users, err := server.store.ListUsers(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, utils.ErrorResponse(err.Error()))
		return
	}

	ctx.JSON(http.StatusOK, users)
}

func (server *Server) signup(ctx *gin.Context) {
	var req userRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, utils.ErrorResponse(err.Error()))
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), 10)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, utils.ErrorResponse(err.Error()))
		return
	}
	arg := db.CreateUserParams{
		Email:    req.Email,
		Password: string(hashedPassword),
	}

	_, err = server.store.CreateUser(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, utils.ErrorResponse(err.Error()))
		return
	}

	ctx.JSON(http.StatusOK, "User created successfully")
}

func (server *Server) login(ctx *gin.Context) {
	var req userRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, utils.ErrorResponse(err.Error()))
		return
	}

	user, err := server.store.GetUser(ctx, req.Email)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, utils.ErrorResponse(err.Error()))
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
	ctx.SetSameSite(http.SameSiteLaxMode)
	ctx.SetCookie("Authorization", tokenString, 3600*24, "/", "localhost", false, true)

	ctx.JSON(http.StatusOK, gin.H{"token": tokenString})
}

func (server *Server) logout(ctx *gin.Context) {
	ctx.SetSameSite(http.SameSiteLaxMode)
	ctx.SetCookie("Authorization", "", -1, "", "", false, true)
	ctx.JSON(http.StatusOK, "Logged out successfully")
}

func (server *Server) validateToken(ctx *gin.Context) {
	user, exists := ctx.Get("user")
	if !exists {
		ctx.JSON(http.StatusUnauthorized, utils.ErrorResponse("Unauthorized"))
		return
	}

	ctx.JSON(http.StatusOK, user)
}
