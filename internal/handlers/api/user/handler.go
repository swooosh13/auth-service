package user

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/go-chi/chi"
	"github.com/go-playground/validator/v10"
	"github.com/swooosh13/quest-auth/internal/domain/user"
	"github.com/swooosh13/quest-auth/internal/handlers/api"
	"github.com/swooosh13/quest-auth/pkg/logger"
	"go.uber.org/zap"
)

var validate = validator.New()

type handler struct {
	service user.Service
}

func NewHandler(service user.Service) api.Handler {
	return &handler{service: service}
}

func (h *handler) Register(r *chi.Mux) {
	r.Route("/user", func(r chi.Router) {
		r.Use(api.LogRequest)
		r.Post("/sign-in", h.SignInHandler)
		r.Post("/sign-up", h.SignUpHandler)
		r.With(api.Authentication).Get("/private", h.PrivateExample)
	})
}

func (h *handler) SignInHandler(rw http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()

	var reqUser user.SignInUserDTO

	if err := json.NewDecoder(r.Body).Decode(&reqUser); err != nil {
		http.Error(rw, fmt.Sprintf("%s:%s", "invalid parse body", err), http.StatusBadRequest)
		return
	}

	validationErr := validate.Struct(reqUser)
	if validationErr != nil {
		http.Error(rw, validationErr.Error(), http.StatusBadRequest)
		return
	}

	foundUser, err := h.service.SignIn(ctx, reqUser)
	if err != nil {
		http.Error(rw, fmt.Sprintf("%s:%s", "invalid sign in error", err), http.StatusInternalServerError)
		return
	}

	jsonData, err := json.Marshal(foundUser)
	if err != nil {
		http.Error(rw, fmt.Sprintf("%s:%s", "error parse user body", err), http.StatusInternalServerError)
		return
	}

	rw.Header().Set("Content-Type", "application/json")
	rw.Write(jsonData)
	return
}

func (h *handler) SignUpHandler(rw http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()

	var user user.CreateUserDTO
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(rw, fmt.Sprintf("%s:%s", "invalid parse body", err), http.StatusBadRequest)
		return
	}

	validationErr := validate.Struct(user)
	if validationErr != nil {
		http.Error(rw, validationErr.Error(), http.StatusBadRequest)
		return
	}

	insertedId, err := h.service.SignUp(ctx, user)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}
	defer cancel()

	rw.Header().Set("Content-Type", "application/json")

	jsonData, err := json.Marshal(insertedId)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}

	rw.Write(jsonData)
	return
}

func (h *handler) PrivateExample(rw http.ResponseWriter, r *http.Request) {
	requestID := r.Context().Value("login").(string)
	logger.Info("private-info", zap.String("login", requestID))
	rw.Write([]byte("hello"))
}
