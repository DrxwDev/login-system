package postgres

import (
	"github.com/google/uuid"

	"github.com/DrxwDev/login-system/internal/adapters/database/sqlc"
	"github.com/DrxwDev/login-system/internal/core/domain/user"
)

func userFromDBToDomain(row sqlc.User) (user.User, error) {
	id, err := user.NewUserID(row.ID.String())
	if err != nil {
		return user.User{}, err
	}

	email, err := user.NewEmail(row.Email)
	if err != nil {
		return user.User{}, err
	}

	password, err := user.NewHashedPassword(row.Password)
	if err != nil {
		return user.User{}, err
	}

	u, err := user.NewUser(
		id,
		row.Name,
		email,
		password,
	)
	if err != nil {
		return user.User{}, err
	}

	u.CreatedAt = row.CreatedAt

	return u, nil
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
