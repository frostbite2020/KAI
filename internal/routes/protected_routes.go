package routes

import (
	"MsKAI/internal/middleware"
	"MsKAI/internal/services"
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"gorm.io/gorm"
)

// RegisterProtectedRoutes mendaftarkan rute yang membutuhkan otentikasi
func RegisterProtectedRoutes(r chi.Router, db *gorm.DB) {
	r.Route("/api", func(api chi.Router) {
		// Middleware otentikasi untuk semua rute di bawah "/api"
		api.Use(middleware.AuthMiddleware(db))

		// Rute yang dilindungi
		api.Get("/profile", services.GetProfile)

		api.Get("/", services.HelloWorldHandler)

		r.Get("/health", services.HealthCheckHandler)

		api.Post("/logout", services.LogoutHandler)

		api.Get("/trains", func(w http.ResponseWriter, r *http.Request) {
			services.GetTrains(w, r, db)
		})

		api.Post("/train", func(w http.ResponseWriter, r *http.Request) {
			// Call CreateTrain and pass in the necessary arguments
			err := services.CreateTrain(w, r, db)
			if err != nil {
				// Handle error properly
				http.Error(w, fmt.Sprintf("Error creating train: %v", err), http.StatusInternalServerError)
			}
		})
	})
}
