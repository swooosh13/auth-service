package user

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID             primitive.ObjectID `bson:"_id"`
	FirstName      string `json:"first_name"`
	LastName       string
	Password       string
	Email          string
	Phone          string
	Token          string
	RefreshedToken string
	CreatedAt      time.Time
	UpdatedAt      time.Time
}
