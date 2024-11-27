package models

import (
	"gorm.io/gorm"
)

type Carriage struct {
	gorm.Model
	TrainID  uint   `json:"train_id"` // Foreign key to Train
	Train    Train  `json:"train"`    // Association with Train
	Seats    []Seat `json:"seats"`    // One-to-many relation with Seat
	Number   string `json:"number"`   // Carriage number (e.g., "1", "2", "3", etc.)
	Type     string `json:"type"`     // e.g., "Sleeper", "Seat", etc.
	Capacity int    `json:"capacity"` // Total number of seats in this carriage
}

func (Carriage) TableName() string {
	return "carriages"
}
