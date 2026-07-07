package ports

type TokenManager interface {
	Generate(userID string) (string, error)
	Validate(token string) (string, error)
}
