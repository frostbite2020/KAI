package models

import (
	"time"

	"gorm.io/gorm"
)

type Schedule struct {
	gorm.Model
	TrainID   uint       `json:"train_id"` // Foreign key for Train
	Train     Train      `json:"train"`    // Association with Train
	RouteID   uint       `json:"route_id"` // Foreign key for Route
	Route     Route      `json:"route"`    // Association with Route
	Departure time.Time  `json:"departure"`
	Arrival   time.Time  `json:"arrival"`
	Carriages []Carriage `json:"carriages"` // One-to-many relation with Carriage
}
