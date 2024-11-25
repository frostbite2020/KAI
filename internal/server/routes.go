package server

import (
	"context"
	"encoding/json"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"github.com/markbates/goth/gothic"
)

func (s *Server) RegisterRoutes() http.Handler {
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	r.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS", "PATCH"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type"},
		AllowCredentials: true,
		MaxAge:           300,
	}))

	r.Get("/", s.HelloWorldHandler)

	r.Get("/health", s.healthHandler)

	r.Get("/auth/{provider}/callback", s.getAuthCallbackFunction)

	r.Get("/logout/{provider}", func(res http.ResponseWriter, req *http.Request) {
		gothic.Logout(res, req)
		res.Header().Set("Location", "/")
		res.WriteHeader(http.StatusTemporaryRedirect)
	})

	r.Get("/auth/{provider}", func(res http.ResponseWriter, req *http.Request) {
		provider := chi.URLParam(req, "provider")

		// Add the provider to the request context
		req = req.WithContext(context.WithValue(req.Context(), "provider", provider))

		// Try to complete authentication (if user is already authenticated)
		if user, err := gothic.CompleteUserAuth(res, req); err == nil {
			log.Printf("User already authenticated: %+v", user)

			// Redirect to the frontend (e.g., dashboard or another page)
			http.Redirect(res, req, "http://localhost:5173/", http.StatusFound)
			return
		}

		// Start authentication flow
		log.Printf("Starting authentication with provider: %s", provider)
		gothic.BeginAuthHandler(res, req)
	})
	return r
}

func (s *Server) HelloWorldHandler(w http.ResponseWriter, r *http.Request) {
	resp := make(map[string]string)
	resp["message"] = "Hello World"

	jsonResp, err := json.Marshal(resp)
	if err != nil {
		log.Fatalf("error handling JSON marshal. Err: %v", err)
	}

	_, _ = w.Write(jsonResp)
}

func (s *Server) healthHandler(w http.ResponseWriter, r *http.Request) {
	jsonResp, _ := json.Marshal(s.db.Health())
	_, _ = w.Write(jsonResp)
}

func (s *Server) getAuthCallbackFunction(w http.ResponseWriter, r *http.Request) {
	provider := chi.URLParam(r, "provider")
	r = r.WithContext(context.WithValue(context.Background(), "provider", provider))

	// Complete the authentication process
	user, err := gothic.CompleteUserAuth(w, r)
	if err != nil {
		log.Printf("Authentication failed: %v", err)
		http.Redirect(w, r, "http://localhost:5173/", http.StatusFound)
		return
	}

	log.Printf("User authenticated: %+v", user)

	// Redirect to frontend (dashboard or dynamic URL)
	redirectURL := r.URL.Query().Get("redirect_to")
	if redirectURL == "" {
		redirectURL = "http://localhost:5173/"
	}

	http.Redirect(w, r, redirectURL, http.StatusFound)
}
