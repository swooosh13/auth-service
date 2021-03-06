package api

import (
	"context"
	"fmt"
	"net/http"

	"github.com/swooosh13/auth-service/pkg/logger"
	"github.com/swooosh13/auth-service/pkg/token"
	"go.uber.org/zap"
)

func LogRequest(next http.Handler) http.Handler {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			logger.Info("request_info", zap.String("method", r.Method), zap.String("request", r.URL.Path))
			next.ServeHTTP(w, r)
		},
	)
}

func Authentication(next http.Handler) http.Handler {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			clientToken := r.Header.Get("token")
			if clientToken == "" {
				http.Error(w, fmt.Sprintf("Error token"), http.StatusUnauthorized)
				return
			}

			claims, err := token.ValidateToken(clientToken)
			if err != "" {
				http.Error(w, fmt.Sprintf("Invalid token: %s", err), http.StatusInternalServerError)
				return
			}

			ctx := context.WithValue(r.Context(), "login", claims.Login)
			next.ServeHTTP(w, r.WithContext(ctx))
		},
	)
}
