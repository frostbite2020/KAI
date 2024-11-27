package services

import (
	"MsKAI/internal/models"
	"encoding/json"
	"net/http"

	"gorm.io/gorm"
)

func CreateCarriage(w http.ResponseWriter, r *http.Request, db *gorm.DB) {
	var carriage models.Carriage

	// Decode JSON payload
	if err := json.NewDecoder(r.Body).Decode(&carriage); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Save the carriage to the database
	if err := db.Create(&carriage).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(carriage)
}

func GetCarriages(w http.ResponseWriter, r *http.Request, db *gorm.DB) {
	var carriages []models.Carriage

	// Query the database for all carriages and preload Train
	if err := db.Preload("Train").Find(&carriages).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(carriages)
}
