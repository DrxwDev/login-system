// Package http
package http

import (
	"net/http"

	"go.uber.org/fx"
)

var Module = fx.Module(
	"http",
	fx.Provide(
		NewGinRouter,
		NewServer,
	),
	fx.Invoke(
		RegisterRoutes,
		func(_ *http.Server) {
		},
	),
)
