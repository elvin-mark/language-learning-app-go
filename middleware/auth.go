package middleware

import (
	"context"
	"language-learning-app/auth"
	"net/http"
	"strings"
)

type contextKey string

const userContextKey contextKey = "user"

func JWTAuth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			http.Error(w, "Missing Authorization Header", http.StatusUnauthorized)
			return
		}

		// Header is usually "Bearer <token>"
		tokenString := strings.TrimPrefix(authHeader, "Bearer ")

		claims, err := auth.VerifyWithClaims(tokenString)
		if err != nil {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		// Store claims in context to use in subsequent handlers (To be used later)
		ctx := context.WithValue(r.Context(), userContextKey, claims)

		// Support current implementation
		r.Header.Set("User-Id", claims.UserID)

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
