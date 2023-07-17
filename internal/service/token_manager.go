package service

import (
	"context"
	"errors"
	"github.com/auth0/go-jwt-middleware/v2/jwks"
	"github.com/auth0/go-jwt-middleware/v2/validator"
	"github.com/chizidotdev/copia/pkg/utils"
	"github.com/golang-jwt/jwt/v5"
	"log"
	"net/url"
	"strings"
	"time"
)

type TokenManager interface {
	NewJWT(subject string) (string, error)
	Parse(accessToken string) (string, error)
	ValidateToken(ctx context.Context, accessToken string) (interface{}, error)
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

// ValidateToken is a middleware that will check the validity of our JWT.
func (t *tokenMangerService) ValidateToken(ctx context.Context, accessToken string) (interface{}, error) {
	issuerURL, err := url.Parse("https://" + utils.EnvVars.Auth0Domain + "/")
	if err != nil {
		log.Fatalf("Failed to parse the issuer url: %v", err)
	}

	provider := jwks.NewCachingProvider(issuerURL, 5*time.Minute)

	jwtValidator, err := validator.New(
		provider.KeyFunc,
		validator.RS256,
		issuerURL.String(),
		[]string{utils.EnvVars.Auth0Audience},
		validator.WithAllowedClockSkew(time.Minute),
	)
	if err != nil {
		log.Fatal("failed to set up the jwt validator")
	}

	splitToken := strings.Split(accessToken, "Bearer ")
	accessToken = splitToken[1]

	tokenInfo, err := jwtValidator.ValidateToken(ctx, accessToken)
	if err != nil {
		return nil, errors.New("invalid token")
	}

	return tokenInfo, nil
}
