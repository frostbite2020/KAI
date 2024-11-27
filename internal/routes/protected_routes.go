package routes

import (
	"MsKAI/internal/middleware"
	"MsKAI/internal/services"
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

		// City Endpoints
		api.Get("/cities", func(w http.ResponseWriter, r *http.Request) {
			services.GetCities(w, r, db)
		})

		api.Post("/city", func(w http.ResponseWriter, r *http.Request) {
			services.CreateCity(w, r, db)
		})

		// Train Endpoints
		api.Get("/trains", func(w http.ResponseWriter, r *http.Request) {
			services.GetTrains(w, r, db)
		})

		api.Post("/train", func(w http.ResponseWriter, r *http.Request) {
			services.CreateTrain(w, r, db)
		})

		// Station Endpoints
		api.Post("/station", func(w http.ResponseWriter, r *http.Request) {
			services.CreateStation(w, r, db)
		})
		api.Get("/stations", func(w http.ResponseWriter, r *http.Request) {
			services.GetStations(w, r, db)
		})

		// Route Endpoints
		api.Post("/route", func(w http.ResponseWriter, r *http.Request) {
			services.CreateRoute(w, r, db)
		})
		api.Get("/routes", func(w http.ResponseWriter, r *http.Request) {
			services.GetRoutes(w, r, db)
		})

		// Carriage Endpoints
		api.Post("/carriage", func(w http.ResponseWriter, r *http.Request) {
			services.CreateCarriage(w, r, db)
		})
		api.Get("/carriages", func(w http.ResponseWriter, r *http.Request) {
			services.GetCarriages(w, r, db)
		})

		// Schedules Endpoints
		api.Post("/schedule", func(w http.ResponseWriter, r *http.Request) {
			services.CreateSchedule(w, r, db)
		})
		api.Get("/schedules", func(w http.ResponseWriter, r *http.Request) {
			services.GetSchedules(w, r, db)
		})

		api.Post("/schedule-carriage-price", func(w http.ResponseWriter, r *http.Request) {
			services.CreateScheduleCarriagePrice(w, r, db)
		})
		api.Get("/schedule-carriage-prices", func(w http.ResponseWriter, r *http.Request) {
			services.GetScheduleCarriagePrices(w, r, db)
		})

		api.Get("/schedules", func(w http.ResponseWriter, r *http.Request) {
			services.GetSchedules(w, r, db)
		})
	})
}
