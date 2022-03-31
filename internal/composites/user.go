package composites

import (
	user2 "github.com/swooosh13/auth-service/internal/domain/user"
	"github.com/swooosh13/auth-service/internal/handlers/api"
	"github.com/swooosh13/auth-service/internal/handlers/api/user"
	user1 "github.com/swooosh13/auth-service/internal/repository/user"
)

type UserComposite struct {
	Storage user2.Storage
	Service user2.Service
	Handler api.Handler
}

func NewUserComposite(mongodbComposite *MongoDBComposite) (*UserComposite, error) {
	storage := user1.NewStorage(mongodbComposite.db)
	service := user2.NewService(storage)
	handler := user.NewHandler(service)
	return &UserComposite{Storage: storage, Service: service, Handler: handler}, nil
}
