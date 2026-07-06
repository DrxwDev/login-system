// Package database
package database

import (
	"fmt"

	"go.uber.org/fx"
)

var Module = fx.Module(
	"database",
	fx.Provide(
		NewPool,
	),
	fx.Invoke(
		func() {
			fmt.Println("Connected to Postgres")
		},
	),
)
