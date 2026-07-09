package sqlc

import (
	"github.com/jackc/pgx/v5/pgxpool"
	"go.uber.org/fx"
)

func NewQueries(pool *pgxpool.Pool) *Queries {
	return New(pool)
}

var Module = fx.Module(
	"sqlc",
	fx.Provide(
		NewQueries,
	),
)
