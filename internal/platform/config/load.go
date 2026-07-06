package config

import (
	"os"
	"time"
)

func LoadApp() AppConfig {
	return AppConfig{
		HOST:    os.Getenv("HOST"),
		PORT:    os.Getenv("PORT"),
		GinMode: os.Getenv("GIN_MODE"),
	}
}

func LoadServer() ServerConfig {
	return ServerConfig{
		ReadHeaderTimeout: 2 * time.Second,
		WriteTimeout:      5 * time.Second,
		ReadTimeout:       5 * time.Second,
		IdleTimeout:       30 * time.Second,
		MaxHeaderBytes:    1 << 20,
	}
}

func LoadDatabase() DatabaseConfig {
	return DatabaseConfig{
		DSN: os.Getenv("DSN"),
	}
}
