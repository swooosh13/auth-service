package mongodb

import "go.mongodb.org/mongo-driver/mongo"

func OpenCollection(db *mongo.Database, c string) *mongo.Collection {
	collection := db.Collection(c)
	return collection
}
