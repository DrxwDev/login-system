package postgres

import (
	"go.uber.org/fx"

	"github.com/DrxwDev/login-system/internal/core/ports"
)

var Module = fx.Module(
	"postgres",
	fx.Provide(
		fx.Annotate(
			NewUserRepository,
			fx.As(new(ports.UserRepository)),
		),
	),
)
