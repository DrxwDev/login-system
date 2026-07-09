package token

import (
	"go.uber.org/fx"

	"github.com/DrxwDev/login-system/internal/core/ports"
)

var Module = fx.Module(
	"token",
	fx.Provide(
		fx.Annotate(
			NewTokenService,
			fx.As(new(ports.TokenManager)),
		),
	),
)
