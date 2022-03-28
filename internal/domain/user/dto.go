package user

type CreateUserDTO struct {
	Login    string `json:"login" validate:"required,min=2,max=100"`
	Password string `json:"password" validate:"required,min=2,max=100"`
}

type SignInUserDTO struct {
	Login    string `json:"login" validate:"required,min=2,max=100"`
	Password string `json:"password" validate:"required,min=2,max=100"`
}

type UpdateUserDTO struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}

type InsertedId struct {
	InsertedId interface{} `json:"InsertedID"`
}
