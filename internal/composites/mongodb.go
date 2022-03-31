package composites

import (
	"context"

	"github.com/swooosh13/auth-service/pkg/client/mongodb"
	"go.mongodb.org/mongo-driver/mongo"
)

type MongoDBComposite struct {
	db *mongo.Database
}

func NewMongoDBComposite(ctx context.Context, Host, Port, Username, Password, Database, AuthSource string) (*MongoDBComposite, error) {
	mongoclient, err := mongodb.NewClient(ctx, Host, Port, Username, Password, Database, AuthSource)
	if err != nil {
		return nil, err
	}

	return &MongoDBComposite{db: mongoclient}, err
}

func (c *MongoDBComposite) GetDb() *mongo.Database {
	return c.db
}
