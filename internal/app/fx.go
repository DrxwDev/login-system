// Package app
package app

import (
	"go.uber.org/fx"

	"github.com/DrxwDev/login-system/internal/adapters/auth/token"
	"github.com/DrxwDev/login-system/internal/adapters/crypto/argon"
	"github.com/DrxwDev/login-system/internal/adapters/database/postgres"
	"github.com/DrxwDev/login-system/internal/adapters/database/sqlc"
	"github.com/DrxwDev/login-system/internal/adapters/http"
	"github.com/DrxwDev/login-system/internal/adapters/http/handlers"
	"github.com/DrxwDev/login-system/internal/core/usecase"
	"github.com/DrxwDev/login-system/internal/platform/config"
	"github.com/DrxwDev/login-system/internal/platform/database"
	"github.com/DrxwDev/login-system/internal/platform/logger"
	"github.com/DrxwDev/login-system/internal/platform/validation"
)

var Module = fx.Options(
	config.Module,
	logger.Module,
	http.Module,
	argon.Module,
	database.Module,
	postgres.Module,
	sqlc.Module,
	token.Module,
	usecase.Module,
	validation.Module,
	handlers.Module,
)
