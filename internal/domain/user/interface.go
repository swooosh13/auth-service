package user

type Storage interface {
	SignIn() (string, error)
	SignUp() error
}

type Service interface {
	SignIn() (string, error)
	SignUp() error
}
