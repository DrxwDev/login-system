package http

import (
	"github.com/gin-gonic/gin"

	"github.com/DrxwDev/login-system/internal/platform/config"
)

func NewGinRouter(cfg config.AppConfig) *gin.Engine {
	gin.SetMode(cfg.GinMode)
	router := gin.New()
	router.HandleMethodNotAllowed = true
	router.Use(gin.Recovery(), gin.Logger())

	return router
}
