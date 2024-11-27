package services

import (
	"MsKAI/internal/models"
	"encoding/json"
	"net/http"
)

func GetProfile(w http.ResponseWriter, r *http.Request) {
	user, ok := r.Context().Value("user").(models.User)
	if !ok {
		http.Error(w, "User not found", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(user)
}
