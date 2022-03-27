package user

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/swooosh13/quest-auth/internal/domain/user"
	"github.com/swooosh13/quest-auth/internal/handlers/api"
)

type handler struct {
	service user.Service
}

func NewHandler(service user.Service) api.Handler {
	return &handler{service: service}
}

func (h *handler) Register(r *chi.Mux) {
	r.Route("/user", func(r chi.Router) {
		r.Get("/sign-in", h.SignInHandler)
		r.Get("/sign-up", h.SignUpHandler)
	})
}

func (h *handler) SignInHandler(rw http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(rw, "SignIn handker")
}

func (h *handler) SignUpHandler(rw http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(rw, "SIgnUp handler")
}
