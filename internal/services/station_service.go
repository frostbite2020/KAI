package services

import (
	"MsKAI/internal/models"

	"gorm.io/gorm"
)

func CreateStation(db *gorm.DB, station *models.Station) error {
	return db.Create(station).Error
}

func GetStations(db *gorm.DB) ([]models.Station, error) {
	var stations []models.Station
	err := db.Preload("City").Find(&stations).Error
	return stations, err
}

func UpdateStation(db *gorm.DB, station *models.Station) error {
	return db.Save(station).Error
}

func DeleteStation(db *gorm.DB, stationID uint) error {
	return db.Delete(&models.Station{}, stationID).Error
}
