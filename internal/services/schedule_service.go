package services

import (
	"MsKAI/internal/models"

	"gorm.io/gorm"
)

func UpdateSchedule(db *gorm.DB, schedule *models.Schedule) error {
	return db.Save(schedule).Error
}

func DeleteSchedule(db *gorm.DB, scheduleID uint) error {
	return db.Delete(&models.Schedule{}, scheduleID).Error
}
