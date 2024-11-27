package services

import (
	"MsKAI/internal/models"

	"gorm.io/gorm"
)

func CreateTrain(db *gorm.DB, train *models.Train) error {
	return db.Create(train).Error
}

func GetTrains(db *gorm.DB) ([]models.Train, error) {
	var trains []models.Train
	err := db.Preload("Schedules").Find(&trains).Error
	return trains, err
}

func UpdateTrain(db *gorm.DB, train *models.Train) error {
	return db.Save(train).Error
}

func DeleteTrain(db *gorm.DB, trainID uint) error {
	return db.Delete(&models.Train{}, trainID).Error
}
