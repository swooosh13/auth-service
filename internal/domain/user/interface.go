package user

import (
	"context"
)

type Storage interface {
	SignUp(ctx context.Context, user User) (*InsertedId, error)
	Count(ctx context.Context, field string, value string) (int, error)
	FindOne(ctx context.Context, field string, value string) (*User, error)
	UpdateAllTokens(ctx context.Context, signedToken string, signedRefreshToken string, userId string) error
}

type Service interface {
	SignIn(ctx context.Context, user SignInUserDTO) (*User, error)
	SignUp(ctx context.Context, user CreateUserDTO) (*InsertedId, error)
	Count(ctx context.Context, field string, value string) (int, error)
}
