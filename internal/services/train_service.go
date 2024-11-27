package services

import (
	"MsKAI/internal/models"
	"encoding/json"
	"log"
	"net/http"

	"gorm.io/gorm"
)

func CreateTrain(w http.ResponseWriter, r *http.Request, db *gorm.DB) error {
	var train models.Train

	json.NewDecoder(r.Body).Decode(&train)

	db.Create(&train)

	w.WriteHeader(http.StatusCreated)
	return json.NewEncoder(w).Encode(train)
}

func GetTrains(w http.ResponseWriter, r *http.Request, db *gorm.DB) {
	var trains []models.Train

	if err := db.Find(&trains).Error; err != nil {
		http.Error(w, "Failed to fetch trains", http.StatusInternalServerError)
		log.Println("Error fetching trains:", err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(trains)
}

func UpdateTrain(db *gorm.DB, train *models.Train) error {
	return db.Save(train).Error
}

func DeleteTrain(db *gorm.DB, trainID uint) error {
	return db.Delete(&models.Train{}, trainID).Error
}
