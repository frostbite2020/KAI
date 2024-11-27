package services

import (
	"MsKAI/internal/models"

	"gorm.io/gorm"
)

func CreateCarriage(db *gorm.DB, carriage *models.Carriage) error {
	return db.Create(carriage).Error
}

func GetCarriages(db *gorm.DB) ([]models.Carriage, error) {
	var carriages []models.Carriage
	err := db.Preload("Train").Find(&carriages).Error
	return carriages, err
}

func UpdateCarriage(db *gorm.DB, carriage *models.Carriage) error {
	return db.Save(carriage).Error
}

func DeleteCarriage(db *gorm.DB, carriageID uint) error {
	return db.Delete(&models.Carriage{}, carriageID).Error
}
