package argon

import (
	"runtime"

	"github.com/alexedwards/argon2id"

	"github.com/DrxwDev/login-system/internal/core/ports"
)

type Hasher struct {
	params *argon2id.Params
}

func NewHasher() *Hasher {
	return &Hasher{
		params: &argon2id.Params{
			Memory:      64 * 1024,
			Iterations:  2,
			Parallelism: uint8(runtime.NumCPU()),
			SaltLength:  16,
			KeyLength:   32,
		},
	}
}

var _ ports.PasswordHasher = (*Hasher)(nil)

func (h Hasher) Hash(password string) (string, error) {
	if password == "" {
		return "", ErrPasswordRequired
	}

	return argon2id.CreateHash(password, h.params)
}

func (h Hasher) Compare(password, hash string) error {
	if password == "" {
		return ErrPasswordRequired
	}

	if hash == "" {
		return ErrHashRequired
	}

	match, err := argon2id.ComparePasswordAndHash(password, hash)
	if err != nil {
		return ErrHashMismatch
	}

	if !match {
		return ErrPasswordMismatch
	}

	return nil
}
