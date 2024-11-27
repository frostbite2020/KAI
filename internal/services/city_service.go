package services

import (
	"MsKAI/internal/models"
	"encoding/json"
	"fmt"
	"net/http"

	"gorm.io/gorm"
)

func CreateCity(w http.ResponseWriter, r *http.Request, db *gorm.DB) {
	var city models.City
	err := json.NewDecoder(r.Body).Decode(&city)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	if err := db.Create(&city).Error; err != nil {
		http.Error(w, fmt.Sprintf("Error creating city: %v", err), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(city)
}

func GetCities(w http.ResponseWriter, r *http.Request, db *gorm.DB) {
	var cities []models.City
	if err := db.Find(&cities).Error; err != nil {
		http.Error(w, fmt.Sprintf("Error fetching cities: %v", err), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(cities)
}
