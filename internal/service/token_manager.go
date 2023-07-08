package service

import (
	"github.com/golang-jwt/jwt/v5"
	"time"
)

type TokenManager interface {
	NewJWT(subject string) (string, error)
	Parse(accessToken string) (string, error)
	NewRefreshToken() (string, error)
}

type tokenMangerService struct {
	signingKey string
}

func NewTokenManager(signingKey string) TokenManager {
	return &tokenMangerService{
		signingKey: signingKey,
	}
}

func (t *tokenMangerService) NewJWT(subject string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": subject,
		"exp": time.Now().Add(time.Hour).Unix(), // 1 hour
	})

	return token.SignedString([]byte(t.signingKey))
}

func (t *tokenMangerService) Parse(accessToken string) (string, error) {
	token, err := jwt.Parse(accessToken, func(token *jwt.Token) (interface{}, error) {
		return []byte(t.signingKey), nil
	})
	if err != nil {
		return "", err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return "", err
	}

	if float64(time.Now().Unix()) > claims["exp"].(float64) {
		return "", err
	}

	return claims["sub"].(string), nil
}

func (t *tokenMangerService) NewRefreshToken() (string, error) {
	return "", nil
}
