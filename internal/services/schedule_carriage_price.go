package services

import (
	"MsKAI/internal/models"
	"encoding/json"
	"net/http"

	"gorm.io/gorm"
)

func CreateScheduleCarriagePrice(w http.ResponseWriter, r *http.Request, db *gorm.DB) {
	var price models.ScheduleCarriagePrice

	// Decode JSON payload
	if err := json.NewDecoder(r.Body).Decode(&price); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Save the schedule-carriage price to the database
	if err := db.Create(&price).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(price)
}

func GetScheduleCarriagePrices(w http.ResponseWriter, r *http.Request, db *gorm.DB) {
	var prices []models.ScheduleCarriagePrice

	if err := db.Preload("Schedule.Train").
		Preload("Schedule.Route.StartStation").
		Preload("Schedule.Route.EndStation").
		Preload("Carriage").
		Find(&prices).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(prices)
}
