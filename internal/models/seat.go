package models

import (
	"gorm.io/gorm"
)

type Seat struct {
	gorm.Model
	CarriageID uint     `json:"carriage_id"`
	Carriage   Carriage `gorm:"foreignkey:CarriageID" json:"carriage"`
	SeatNumber string   `json:"seat_number"`
	Booked     bool     `json:"booked"`
}

func (Seat) TableName() string {
	return "seats"
}
