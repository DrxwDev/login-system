package user

type RegisterParams struct {
	Name     string
	Email    string
	Password string
}

type LoginParams struct {
	Email    string
	Password string
}
