package models

import (
	"time"

	"gorm.io/gorm"
)

type Booking struct {
	gorm.Model
	UserID      uint      `json:"user_id"`
	User        User      `gorm:"foreignkey:UserID" json:"user"`
	ScheduleID  uint      `json:"schedule_id"`
	Schedule    Schedule  `gorm:"foreignkey:ScheduleID" json:"schedule"`
	BookedSeats []Seat    `gorm:"many2many:booking_seats;" json:"booked_seats"`
	BookingDate time.Time `json:"booking_date"`
	TotalAmount float64   `json:"total_amount"`
}

func (Booking) TableName() string {
	return "bookings"
}
