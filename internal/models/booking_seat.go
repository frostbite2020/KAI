package models

import "gorm.io/gorm"

type BookingSeat struct {
	gorm.Model
	SeatID    uint    `json:"seat_id"`
	Seat      Seat    `gorm:"foreignkey:SeatID" json:"seat"`
	BookingID uint    `json:"booking_id"`
	Booking   Booking `gorm:"foreignkey:BookingID" json:"booking"`
}

func (BookingSeat) TableName() string {
	return "booking_seats"
}
