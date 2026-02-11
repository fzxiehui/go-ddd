package middleware

import (
	"context"
	"net/http"
	"strings"

	"ddd/internal/application/service/auth"
)

type JWTMiddleware struct {
	tokenService *auth.TokenService
}

func NewJWTMiddleware(ts *auth.TokenService) *JWTMiddleware {
	return &JWTMiddleware{tokenService: ts}
}

func (m *JWTMiddleware) Handler(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			http.Error(w, "missing token", http.StatusUnauthorized)
			return
		}

		tokenStr := strings.TrimPrefix(authHeader, "Bearer ")

		claims, err := m.tokenService.Parse(tokenStr)
		if err != nil {
			http.Error(w, "invalid token", http.StatusUnauthorized)
			return
		}

		ctx := context.WithValue(r.Context(), "user_id", claims.UserID)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
