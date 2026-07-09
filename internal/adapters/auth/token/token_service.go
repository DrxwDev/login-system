// Package token
package token

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"

	"github.com/DrxwDev/login-system/internal/core/ports"
	"github.com/DrxwDev/login-system/internal/platform/config"
)

const issuer = "access_token"

type TokenService struct {
	cfg config.AppConfig
}

func NewTokenService(cfg config.AppConfig) *TokenService {
	return &TokenService{
		cfg: cfg,
	}
}

func (s *TokenService) Generate(userID string) (string, error) {
	// validate user id
	if userID == "" {
		return "", ErrUserIDNotProvided
	}

	// register token claims
	claims := jwt.RegisteredClaims{
		Issuer:    issuer,
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

func (s *TokenService) Validate(token string) (*ports.TokenClaims, error) {
	claims := &jwt.RegisteredClaims{}

	// parse token
	t, err := jwt.ParseWithClaims(token, claims, func(t *jwt.Token) (any, error) {
		if t.Method.Alg() != jwt.SigningMethodHS256.Alg() {
			return nil, ErrInvalidTokenMethod
		}
		return []byte(s.cfg.JwtSecret), nil
	})
	if err != nil {
		return nil, fmt.Errorf("%w: %v", ErrUnableToParseToken, err)
	}

	// validate token
	if !t.Valid {
		return nil, ErrInvalidAccessToken
	}

	if claims.ExpiresAt == nil ||
		claims.IssuedAt == nil ||
		claims.NotBefore == nil {
		return nil, ErrInvalidAccessToken
	}

	if claims.Subject == "" {
		return nil, ErrInvalidAccessToken
	}

	if claims.Issuer != issuer {
		return nil, ErrInvalidAccessToken
	}

	// return user id as Subject
	return &ports.TokenClaims{
		Subject:   claims.Subject,
		Issuer:    claims.Issuer,
		ExpiresAt: claims.ExpiresAt.Time,
		IssuedAt:  claims.IssuedAt.Time,
		NotBefore: claims.NotBefore.Time,
	}, nil
}
