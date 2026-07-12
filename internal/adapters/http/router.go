package http

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	"github.com/DrxwDev/login-system/internal/platform/config"
)

func NewGinRouter(cfg config.AppConfig) *gin.Engine {
	gin.SetMode(cfg.GinMode)

	router := gin.New()
	router.HandleMethodNotAllowed = true
	router.Use(gin.Recovery(), gin.Logger())

	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000"},
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))
	return router
}
