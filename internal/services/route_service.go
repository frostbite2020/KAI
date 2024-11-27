package services

import (
	"MsKAI/internal/models"
	"encoding/json"
	"net/http"

	"gorm.io/gorm"
)

func CreateRoute(w http.ResponseWriter, r *http.Request, db *gorm.DB) {
	var route models.Route

	if err := json.NewDecoder(r.Body).Decode(&route); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Save the route to the database
	if err := db.Create(&route).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(route)
}

func GetRoutes(w http.ResponseWriter, r *http.Request, db *gorm.DB) {
	var routes []models.Route

	if err := db.Preload("StartStation.City").
		Preload("EndStation.City").
		Find(&routes).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(routes)
}
