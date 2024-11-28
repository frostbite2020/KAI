package middleware

import (
	"MsKAI/internal/models"
	"context"
	"net/http"
	"strings"
	"time"

	"gorm.io/gorm"
)

func AuthMiddleware(db *gorm.DB) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			// Extract Authorization header
			authHeader := r.Header.Get("Authorization")
			if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
				http.Error(w, "Unauthorized", http.StatusUnauthorized)
				return
			}

			// Extract token from the header
			token := strings.TrimPrefix(authHeader, "Bearer ")

			// Check the session in the database
			var session models.Session
			if err := db.Where("token = ? AND expires_at > ?", token, time.Now()).
				Preload("User").First(&session).Error; err != nil {
				http.Error(w, "Unauthorized", http.StatusUnauthorized)
				return
			}

			// Add user to the context
			ctx := context.WithValue(r.Context(), "user", session.User)
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}
