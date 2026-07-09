package user

import (
	"context"

	"github.com/DrxwDev/login-system/internal/core/domain/user"
	"github.com/DrxwDev/login-system/internal/core/ports"
)

type LoginUseCase struct {
	repo   ports.UserRepository
	hasher ports.PasswordHasher
	auth   ports.TokenManager
}

func NewLoginUseCase(repo ports.UserRepository, hasher ports.PasswordHasher, auth ports.TokenManager) *LoginUseCase {
	return &LoginUseCase{
		repo:   repo,
		hasher: hasher,
		auth:   auth,
	}
}

func (uc *LoginUseCase) Login(ctx context.Context, params LoginParams) (user.User, string, error) {
	u, err := uc.repo.FindByEmail(ctx, user.Email(params.Email))
	if err != nil {
		return user.User{}, "", err
	}

	if err := uc.hasher.Compare(params.Password, string(u.Password)); err != nil {
		return user.User{}, "", err
	}

	token, err := uc.auth.Generate(u.ID.String())
	if err != nil {
		return user.User{}, "", err
	}

	return u, token, nil
}
