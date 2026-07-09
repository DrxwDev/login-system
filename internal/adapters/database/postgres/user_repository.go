// Package postgres
package postgres

import (
	"context"

	"github.com/DrxwDev/login-system/internal/adapters/database/sqlc"
	"github.com/DrxwDev/login-system/internal/core/domain/user"
	"github.com/DrxwDev/login-system/internal/core/ports"
)

type UserRepository struct {
	queries *sqlc.Queries
}

var _ ports.UserRepository = (*UserRepository)(nil)

func NewUserRepository(q *sqlc.Queries) *UserRepository {
	return &UserRepository{
		queries: q,
	}
}

func (r UserRepository) Create(ctx context.Context, u user.User) error {
	params, err := userDomainToSaveParams(u)
	if err != nil {
		return err
	}

	err = r.queries.Save(ctx, params)
	if err != nil {
		return err
	}

	return nil
}

func (r UserRepository) FindByEmail(ctx context.Context, email user.Email) (user.User, error) {
	panic("not implemented") // TODO: Implement
}

func (r UserRepository) FindByID(ctx context.Context, id user.UserID) (user.User, error) {
	panic("not implemented") // TODO: Implement
}
