package services

import (
	"MsKAI/internal/models"
	"encoding/json"
	"net/http"

	"gorm.io/gorm"
)

func CreateStation(w http.ResponseWriter, r *http.Request, db *gorm.DB) {
	var station models.Station

	// Decode JSON payload
	if err := json.NewDecoder(r.Body).Decode(&station); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Save the station to the database
	if err := db.Create(&station).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(station)
}

func GetStations(w http.ResponseWriter, r *http.Request, db *gorm.DB) {
	var stations []models.Station

	// Query the database for all stations and preload related City
	if err := db.Preload("City").Find(&stations).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(stations)
}
