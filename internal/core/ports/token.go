package ports

import "time"

type TokenClaims struct {
	Subject   string
	Issuer    string
	ExpiresAt time.Time
	IssuedAt  time.Time
	NotBefore time.Time
}

type TokenManager interface {
	Generate(userID string) (string, error)
	Validate(token string) (*TokenClaims, error)
}
