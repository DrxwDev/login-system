package middlewares

import (
	"github.com/gin-gonic/gin"

	"github.com/DrxwDev/login-system/internal/core/ports"
)

func GetClaims(ctx *gin.Context) (*ports.TokenClaims, bool) {
	claims, ok := ctx.Get(ClaimsContextKey)
	if !ok {
		return nil, false
	}

	tokenClaims, ok := claims.(*ports.TokenClaims)
	if !ok {
		return nil, false
	}

	return tokenClaims, true
}
