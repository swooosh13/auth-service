package main

import (
	"context"
	"fmt"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/swooosh13/auth-service/pkg/logger"
	"go.uber.org/zap"

	"github.com/swooosh13/auth-service/internal/composites"
	"github.com/swooosh13/auth-service/internal/config"
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

	logger.Info("mongodb composite has been created")
	userComposite, err := composites.NewUserComposite(mongodbComposite)
	if err != nil {
		logger.Fatal("failed to create user composite", zap.String("error", err.Error()))
		return
	}
	userComposite.Handler.Register(r)

	logger.Info("server has been started", zap.String("host", cfg.Listen.Host), zap.String("port", cfg.Listen.Port))
	http.ListenAndServe(fmt.Sprintf("%s:%s", cfg.Listen.Host, cfg.Listen.Port), r)
}
