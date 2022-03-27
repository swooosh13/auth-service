package user

type CreateUserDTO struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}

type UpdateUserDTO struct {
	UUID     string `json:"uuid"`
	Login    string `json:"login"`
	Password string `json:"password"`
}
