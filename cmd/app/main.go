package main

import (
	"context"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/swooosh13/quest-auth/pkg/logger"
	"go.mongodb.org/mongo-driver/mongo"
	"go.uber.org/zap"

	"github.com/swooosh13/quest-auth/internal/composites"
	"github.com/swooosh13/quest-auth/internal/config"
)

func main() {
	logger.Init()
	cfg := config.GetConfig()

	r := chi.NewRouter()
	mongodbComposite, err := composites.NewMongoDBComposite(context.Background(), cfg.MongoDB.Host, cfg.MongoDB.Port, cfg.MongoDB.Username, cfg.MongoDB.Password, cfg.MongoDB.Database, cfg.MongoDB.AuthDB)
	if err != nil {
		logger.Fatal("failed to create mongodb composite", zap.String("error", err.Error()))
		return
	}

	_ = OpenCollection(mongodbComposite.GetDb(), "users")

	logger.Info("mongodb composite has been created")
	// userComposite, err := composites.NewUserComposite(mongodbComposite)
	// if err != nil {
	// 	logger.Fatal("failed to create user composite", zap.String("error", err.Error()))
	// 	return
	// }
	// userComposite.Handler.Register(r)

	http.ListenAndServe("127.0.0.1:9000", r)
}

func OpenCollection(db *mongo.Database, c string) (collection *mongo.Collection) {
	collection = db.Collection(c)
	return
}
