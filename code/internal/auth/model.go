package auth

type LoginForm struct {
	Email    string
	Password string
}

type RegisterForm struct {
	Username string
	Email    string
	Password string
}
