package models

import "gorm.io/gorm"

// Schedule model
type Schedule struct {
	gorm.Model
	TrainID   uint   `json:"train_id"`
	Train     Train  `gorm:"foreignKey:TrainID" json:"train"`
	RouteID   uint   `json:"route_id"`
	Route     Route  `gorm:"foreignKey:RouteID" json:"route"`
	Departure string `json:"departure"`
	Arrival   string `json:"arrival"`
}

func (Schedule) TableName() string {
	return "schedules"
}
