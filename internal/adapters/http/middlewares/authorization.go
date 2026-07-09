// Package middlewares
package middlewares

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"

	"github.com/DrxwDev/login-system/internal/core/ports"
)

const (
	bearerPrefix      = "Bearer "
	ClaimsContextKey  = "claims"
	AccessTokenCookie = "access_token"
)

func GetBearerToken(auth ports.TokenManager) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		header := strings.TrimSpace(ctx.GetHeader("Authorization"))
		if header == "" {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"status": "failed",
				"error":  ErrMissingAuthHeader,
			})
			return
		}

		if !strings.HasPrefix(header, bearerPrefix) {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"status": "failed",
				"error":  ErrInvalidAuthFormat,
			})
			return
		}

		token := strings.TrimSpace(strings.TrimPrefix(header, bearerPrefix))
		claims, err := auth.Validate(token)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"status": "failed",
				"error":  ErrInvalidToken,
			})
			return
		}

		ctx.Set(ClaimsContextKey, claims)
		ctx.Next()
	}
}
