package models

import (
	"gorm.io/gorm"
)

type Seat struct {
	gorm.Model
	CarriageID uint     `json:"carriage_id"` // Foreign key to Carriage
	Carriage   Carriage `json:"carriage"`    // Association with Carriage
	SeatNumber string   `json:"seat_number"` // Seat identifier (e.g., A1, B2, etc.)
	SeatType   string   `json:"seat_type"`   // e.g., "Economy", "Business", etc.
	Status     string   `json:"status"`      // Status of the seat: "Available", "Booked"
}

func (Seat) TableName() string {
	return "seats"
}
