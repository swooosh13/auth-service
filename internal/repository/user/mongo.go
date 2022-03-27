package user

import (
	"github.com/swooosh13/quest-auth/internal/domain/user"
	"go.mongodb.org/mongo-driver/mongo"
)

// mongo code
type userStorage struct {
	db *mongo.Database
}

func NewStorage(db *mongo.Database) user.Storage {
	return &userStorage{db: db}
}

func (u *userStorage) SignIn() (string, error) {
	return "", nil
}

func (u *userStorage) SignUp() error {
	return nil
}
