// Package postgres
package postgres

import (
	"context"

	"github.com/google/uuid"

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
		return ErrUnableToCreateUser
	}

	return nil
}

func (r UserRepository) FindByEmail(ctx context.Context, email user.Email) (user.User, error) {
	row, err := r.queries.GetUserByEmail(ctx, email.String())
	if err != nil {
		return user.User{}, ErrUserNotFound
	}

	return userFromDBToDomain(row)
}

func (r UserRepository) FindByID(ctx context.Context, id user.UserID) (user.User, error) {
	uid, err := uuid.Parse(id.String())
	if err != nil {
		return user.User{}, user.ErrInvalidUserID
	}

	row, err := r.queries.GetUserByID(ctx, uid)
	if err != nil {
		return user.User{}, ErrUserNotFound
	}

	return userFromDBToDomain(row)
}
