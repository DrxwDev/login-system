package postgres

import (
	"github.com/google/uuid"

	"github.com/DrxwDev/login-system/internal/adapters/database/sqlc"
	"github.com/DrxwDev/login-system/internal/core/domain/user"
)

func userFromDBToDomain(params sqlc.User) user.User {
	return user.User{
		ID:        user.UserID(params.ID.String()),
		Name:      params.Name,
		Email:     user.Email(params.Email),
		Password:  user.Password(params.Password),
		CreatedAt: params.CreatedAt,
	}
}

func userDomainToSaveParams(u user.User) (sqlc.SaveParams, error) {
	id, err := uuid.Parse(u.ID.String())
	if err != nil {
		return sqlc.SaveParams{}, err
	}

	return sqlc.SaveParams{
		ID:       id,
		Name:     u.Name,
		Email:    u.Email.String(),
		Password: u.Password.String(),
	}, nil
}
