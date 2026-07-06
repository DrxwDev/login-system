package config

import "time"

type ServerConfig struct {
	ReadHeaderTimeout time.Duration
	WriteTimeout      time.Duration
	ReadTimeout       time.Duration
	IdleTimeout       time.Duration
	MaxHeaderBytes    int
}
