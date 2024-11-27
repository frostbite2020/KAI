package services

import (
	"MsKAI/internal/models"

	"gorm.io/gorm"
)

func CreateRoute(db *gorm.DB, route *models.Route) error {
	return db.Create(route).Error
}

func GetRoutes(db *gorm.DB) ([]models.Route, error) {
	var routes []models.Route
	err := db.Preload("StartStation").Preload("EndStation").Find(&routes).Error
	return routes, err
}

func UpdateRoute(db *gorm.DB, route *models.Route) error {
	return db.Save(route).Error
}

func DeleteRoute(db *gorm.DB, routeID uint) error {
	return db.Delete(&models.Route{}, routeID).Error
}
