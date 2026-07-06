// Package argon
package argon

import "go.uber.org/fx"

var Module = fx.Module(
	"argon",
	fx.Provide(
		NewHasher,
	),
)
