package user

import (
	"context"

	"github.com/google/uuid"

	"github.com/DrxwDev/login-system/internal/core/domain/user"
	"github.com/DrxwDev/login-system/internal/core/ports"
)

type RegisterUseCase struct {
	repo   ports.UserRepository
	hasher ports.PasswordHasher
}

func NewRegisterUseCase(repo ports.UserRepository, hasher ports.PasswordHasher) *RegisterUseCase {
	return &RegisterUseCase{
		repo:   repo,
		hasher: hasher,
	}
}

func (uc *RegisterUseCase) Register(ctx context.Context, params RegisterParams) (user.User, error) {
	id, err := user.NewUserID(uuid.NewString())
	if err != nil {
		return user.User{}, err
	}

	email, err := user.NewEmail(params.Email)
	if err != nil {
		return user.User{}, err
	}

	hash, err := uc.hasher.Hash(params.Password)
	if err != nil {
		return user.User{}, err
	}

	password, err := user.NewHashedPassword(hash)
	if err != nil {
		return user.User{}, err
	}

	newUser, err := user.NewUser(id, params.Name, email, password)
	if err != nil {
		return user.User{}, err
	}

	err = uc.repo.Create(ctx, newUser)
	if err != nil {
		return user.User{}, err
	}

	return newUser, nil
}
