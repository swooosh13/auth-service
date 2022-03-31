package user

import (
	"context"
	"fmt"
	"time"

	"github.com/swooosh13/auth-service/pkg/crypto"
	"github.com/swooosh13/auth-service/pkg/token"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type userService struct {
	storage Storage
}

func NewService(storage Storage) Service {
	return &userService{storage: storage}
}

func (u *userService) SignUp(ctx context.Context, createUserDTO CreateUserDTO) (*InsertedId, error) {
	count, err := u.Count(ctx, "login", createUserDTO.Login)

	if err != nil {
		return nil, fmt.Errorf("error occured while checking for the email: %w", err)
	}

	if count > 0 {
		return nil, fmt.Errorf("error: this login already exists: %w", err)
	}

	var user *User = &User{}

	user.Login = createUserDTO.Login

	password := crypto.HashPassword(createUserDTO.Password)
	user.Password = password

	user.CreatedAt, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
	user.UpdatedAt, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
	user.ID = primitive.NewObjectID()
	user.UserId = user.ID.Hex()

	token, refreshToken, _ := token.GenerateAllTokens(createUserDTO.Login, user.UserId)
	user.Token = token
	user.RefreshToken = refreshToken

	resultInsertionNumber, insertErr := u.storage.SignUp(ctx, *user)
	if insertErr != nil {
		return nil, fmt.Errorf("error: %w", err)
	}
	return resultInsertionNumber, nil
}

func (u *userService) SignIn(ctx context.Context, user SignInUserDTO) (*User, error) {
	foundUser, err := u.storage.FindOne(ctx, "login", user.Login)
	if err != nil {
		return nil, err
	}

	ok, msg := crypto.VerifyPassword(user.Password, foundUser.Password)
	if !ok {
		return nil, fmt.Errorf("invalid password error: %s", msg)
	}

	token, refreshToken, _ := token.GenerateAllTokens(foundUser.Login, foundUser.UserId)
	err = u.storage.UpdateAllTokens(ctx, token, refreshToken, foundUser.UserId)
	if err != nil {
		return nil, err
	}

	return foundUser, nil
}

func (u *userService) Count(ctx context.Context, field string, value string) (int, error) {
	return u.storage.Count(ctx, field, value)
}
