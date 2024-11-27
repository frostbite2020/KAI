package services

import (
	"MsKAI/internal/models"

	"gorm.io/gorm"
)

func CreateSchedule(db *gorm.DB, schedule *models.Schedule) error {
	return db.Create(schedule).Error
}

func GetSchedules(db *gorm.DB) ([]models.Schedule, error) {
	var schedules []models.Schedule
	err := db.Preload("Train").Preload("Route").Preload("Carriages").Find(&schedules).Error
	return schedules, err
}

func UpdateSchedule(db *gorm.DB, schedule *models.Schedule) error {
	return db.Save(schedule).Error
}

func DeleteSchedule(db *gorm.DB, scheduleID uint) error {
	return db.Delete(&models.Schedule{}, scheduleID).Error
}
