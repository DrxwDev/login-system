package middlewares

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/DrxwDev/login-system/internal/core/ports"
)

func AuthMiddleware(auth ports.TokenManager) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token, err := ctx.Cookie(AccessTokenCookie)
		if err != nil {
			ctx.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		claims, err := auth.Validate(token)
		if err != nil {
			ctx.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		ctx.Set(ClaimsContextKey, claims)
		ctx.Next()
	}
}
