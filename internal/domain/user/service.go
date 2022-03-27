package user

type userService struct {
	storage Storage
}

func NewService(storage Storage) Service {
	return &userService{storage: storage}
}

func (u *userService) SignIn() (string, error) {
	return u.storage.SignIn()
}

func (u *userService) SignUp() error {
	return u.storage.SignUp()
}
