// Package token
package token

import (
	"time"

	"github.com/golang-jwt/jwt/v5"

	"github.com/DrxwDev/login-system/internal/platform/config"
)

type TokenService struct {
	cfg config.AppConfig
}

func NewTokenService(cfg config.AppConfig) *TokenService {
	return &TokenService{
		cfg: cfg,
	}
}

func (s TokenService) Generate(userID string) (string, error) {
	// validate user id
	if userID == "" {
		return "", ErrUserIDNotProvided
	}

	// register token claims
	claims := jwt.RegisteredClaims{
		Issuer:    "access-token",
		Subject:   userID,
		ExpiresAt: jwt.NewNumericDate(time.Now().Add((24 * time.Hour) * 30)),
		NotBefore: jwt.NewNumericDate(time.Now()),
		IssuedAt:  jwt.NewNumericDate(time.Now()),
	}

	// create token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// signed to string
	tokenString, err := token.SignedString([]byte(s.cfg.JwtSecret))
	if err != nil {
		return "", ErrUnableToCreateToken
	}

	return tokenString, nil
}

func (s TokenService) Validate(token string) (string, error) {
	claims := &jwt.RegisteredClaims{}

	// parse token
	t, err := jwt.ParseWithClaims(token, claims, func(t *jwt.Token) (any, error) {
		if t.Method != jwt.SigningMethodHS256 {
			return nil, ErrInvalidTokenMethod
		}
		return []byte(s.cfg.JwtSecret), nil
	})
	if err != nil {
		return "", ErrUnableToParseToken
	}

	// validate token
	if !t.Valid {
		return "", ErrInvalidAccessToken
	}

	// return user id as Subject
	return claims.Subject, nil
}
