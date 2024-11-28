package routes

import (
	"MsKAI/internal/database"
	"MsKAI/internal/services"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func RegisterRoutes() http.Handler {
	db := database.GetDB()
	r := chi.NewRouter()

	// Middleware for logging and request handling
	r.Use(middleware.Logger)

	r.Get("/auth/{provider}/callback", services.GetAuthCallbackFunction(db))
	r.Get("/auth/{provider}", services.Login)

	// Register specific protected routes
	RegisterProtectedRoutes(r, db)

	return r
}
