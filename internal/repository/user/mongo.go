package user

import (
	"context"
	"fmt"
	"time"

	"github.com/swooosh13/quest-auth/internal/domain/user"
	"github.com/swooosh13/quest-auth/pkg/client/mongodb"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type userStorage struct {
	db *mongo.Database
}

func NewStorage(db *mongo.Database) user.Storage {
	return &userStorage{db: db}
}

func (u *userStorage) Count(ctx context.Context, field string, value string) (int, error) {
	userCollection := mongodb.OpenCollection(u.db, "user")
	count, err := userCollection.CountDocuments(ctx, bson.M{field: value})
	if err != nil {
		return int(count), err
	}

	return int(count), err
}

func (u *userStorage) SignUp(ctx context.Context, userEntity user.User) (*user.InsertedId, error) {
	userCollection := mongodb.OpenCollection(u.db, "user")

	resultInsertionNumber, err := userCollection.InsertOne(ctx, userEntity)
	if err != nil {
		return nil, fmt.Errorf("user was not created")
	}

	return &user.InsertedId{InsertedId: resultInsertionNumber.InsertedID}, nil
}

func (u *userStorage) FindOne(ctx context.Context, field, value string) (*user.User, error) {
	userCollection := mongodb.OpenCollection(u.db, "user")

	var userEntity user.User

	err := userCollection.FindOne(ctx, bson.M{field: value}).Decode(&userEntity)
	if err != nil {
		return nil, err
	}

	return &userEntity, nil
}

func (u *userStorage) UpdateAllTokens(ctx context.Context, signedToken string, signedRefreshToken string, userId string) error {
	userCollection := mongodb.OpenCollection(u.db, "user")

	ctx, cancel := context.WithTimeout(ctx, 100*time.Second)

	var updateObj primitive.D

	updateObj = append(updateObj, bson.E{"token", signedToken})
	updateObj = append(updateObj, bson.E{"refresh_token", signedRefreshToken})

	Updated_at, _ := time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
	updateObj = append(updateObj, bson.E{"updated_at", Updated_at})

	upsert := true
	filter := bson.M{"user_id": userId}
	opt := options.UpdateOptions{
		Upsert: &upsert,
	}

	_, err := userCollection.UpdateOne(
		ctx,
		filter,
		bson.D{
			{"$set", updateObj},
		},
		&opt,
	)
	defer cancel()

	if err != nil {
		return err
	}

	return nil
}
