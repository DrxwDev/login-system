package http

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/DrxwDev/login-system/internal/adapters/http/handlers"
	"github.com/DrxwDev/login-system/internal/adapters/http/middlewares"
	"github.com/DrxwDev/login-system/internal/core/ports"
)

func RegisterRoutes(router *gin.Engine, userH *handlers.UserHandler, auth ports.TokenManager) {
	router.GET("/health", middlewares.AuthMiddleware(auth), func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"status":  "OK",
			"message": "Server running and healthy",
		})
	})

	apiV1 := router.Group("/api/v1")

	apiV1.POST("/register", userH.Register)
	apiV1.POST("/login", userH.Login)
	apiV1.POST("/logout", middlewares.AuthMiddleware(auth), userH.Logout)
	apiV1.GET("/me", middlewares.AuthMiddleware(auth), userH.User)
}
