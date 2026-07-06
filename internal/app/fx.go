// Package app
package app

import (
	"go.uber.org/fx"

	"github.com/DrxwDev/login-system/internal/adapters/crypto/argon"
	"github.com/DrxwDev/login-system/internal/adapters/http"
	"github.com/DrxwDev/login-system/internal/platform/config"
	"github.com/DrxwDev/login-system/internal/platform/logger"
)

var Module = fx.Options(
	config.Module,
	logger.Module,
	http.Module,
	argon.Module,
)
