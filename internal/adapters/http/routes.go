package http

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/DrxwDev/login-system/internal/adapters/http/handlers"
)

func RegisterRoutes(router *gin.Engine, userH *handlers.UserHandler) {
	router.GET("/health", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"status":  "OK",
			"message": "Server running and healthy",
		})
	})

	apiV1 := router.Group("/api/v1")

	apiV1.POST("/register", userH.Register)
	apiV1.POST("/login", userH.Login)
}
