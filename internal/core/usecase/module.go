// Package usecase
package usecase

import (
	"go.uber.org/fx"

	"github.com/DrxwDev/login-system/internal/core/usecase/user"
)

var Module = fx.Options(
	fx.Provide(
		user.NewRegisterUseCase,
		user.NewLoginUseCase,
		user.NewGetUserUseCase,
	),
)
