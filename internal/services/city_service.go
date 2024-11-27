package services

import (
	"MsKAI/internal/models"

	"gorm.io/gorm"
)

func CreateCity(db *gorm.DB, city *models.City) error {
	return db.Create(city).Error
}

func GetCities(db *gorm.DB) ([]models.City, error) {
	var cities []models.City
	err := db.Preload("Stations").Find(&cities).Error
	return cities, err
}

func UpdateCity(db *gorm.DB, city *models.City) error {
	return db.Save(city).Error
}

func DeleteCity(db *gorm.DB, cityID uint) error {
	return db.Delete(&models.City{}, cityID).Error
}
