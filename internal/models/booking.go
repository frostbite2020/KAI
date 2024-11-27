package models

import (
	"time"

	"gorm.io/gorm"
)

type Booking struct {
	gorm.Model
	UserID      uint      `json:"user_id"`      // Foreign key to User
	User        User      `json:"user"`         // Association to User model (not shown here)
	ScheduleID  uint      `json:"schedule_id"`  // Foreign key to Schedule
	Schedule    Schedule  `json:"schedule"`     // Association to Schedule
	BookedSeats []Seat    `json:"booked_seats"` // Many-to-many relationship with Seat
	BookingDate time.Time `json:"booking_date"` // Date of booking
	TotalAmount float64   `json:"total_amount"` // Total booking amount
}

func (Booking) TableName() string {
	return "bookings"
}
