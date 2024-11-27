package services

import (
	"MsKAI/internal/models"
	"MsKAI/internal/utils"
	"context"
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/go-chi/chi"
	"github.com/gorilla/sessions"
	"github.com/markbates/goth/gothic"
	"gorm.io/gorm"
)

const (
	key    = "randomString"
	MaxAge = 86400 * 30
	IsProd = false
)

var store = sessions.NewCookieStore([]byte(key))

func HealthCheckHandler(w http.ResponseWriter, r *http.Request) {
	response := map[string]string{
		"status":  "OK",
		"message": "Service is up and running!",
	}

	jsonResp, err := json.Marshal(response)
	if err != nil {
		log.Printf("Error marshaling response: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonResp)
}

func HelloWorldHandler(w http.ResponseWriter, r *http.Request) {
	response := map[string]string{
		"message": "Hello World",
	}

	jsonResp, err := json.Marshal(response)
	if err != nil {
		log.Printf("Error marshaling response: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonResp)
}

func GetAuthCallbackFunction(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Tambahkan provider ke context
		r = r.WithContext(context.WithValue(r.Context(), "provider", "google"))

		// Selesaikan proses autentikasi dengan Goth
		user, err := gothic.CompleteUserAuth(w, r)
		if err != nil {
			log.Printf("Authentication failed: %v", err)
			http.Redirect(w, r, "http://localhost:5173/", http.StatusFound)
			return
		}

		// Cek atau buat user baru di database
		var dbUser models.User
		if err := db.Where("email = ?", user.Email).First(&dbUser).Error; err != nil {
			if err == gorm.ErrRecordNotFound {
				// Jika user belum ada, buat user baru
				dbUser = models.User{
					Email: user.Email,
					Name:  user.Name,
				}
				if err := db.Create(&dbUser).Error; err != nil {
					log.Printf("Failed to create user: %v", err)
					http.Error(w, "Failed to create user", http.StatusInternalServerError)
					return
				}
			} else {
				// Jika error lain, log dan hentikan
				log.Printf("Failed to query user: %v", err)
				http.Error(w, "Failed to query user", http.StatusInternalServerError)
				return
			}
		}

		// Buat token sesi baru
		token := utils.GenerateToken()
		session := models.Session{
			UserID:    dbUser.ID,
			Token:     token,
			ExpiresAt: time.Now().Add(24 * time.Hour), // Sesi berlaku 24 jam
		}

		// Simpan sesi ke database (update jika sudah ada)
		if err := db.Where("user_id = ?", dbUser.ID).Save(&session).Error; err != nil {
			log.Printf("Failed to create session: %v", err)
			http.Error(w, "Failed to create session", http.StatusInternalServerError)
			return
		}

		// Redirect ke frontend dengan token
		redirectURL := r.URL.Query().Get("redirect_to")
		if redirectURL == "" {
			redirectURL = "http://localhost:5173/"
		}
		redirectURL += "?token=" + token
		http.Redirect(w, r, redirectURL, http.StatusFound)
	}
}

func LogoutHandler(w http.ResponseWriter, r *http.Request) {
	gothic.Logout(w, r)

	http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
}

func Login(res http.ResponseWriter, req *http.Request) {
	provider := chi.URLParam(req, "provider")

	req = req.WithContext(context.WithValue(req.Context(), "provider", "google"))

	if user, err := gothic.CompleteUserAuth(res, req); err == nil {
		log.Printf("User already authenticated: %+v", user)

		http.Redirect(res, req, "http://localhost:5173/", http.StatusFound)
		return
	}

	log.Printf("Starting authentication with provider: %s", provider)
	gothic.BeginAuthHandler(res, req)
}
