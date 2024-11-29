package models

import (
	"gorm.io/gorm"
)

type Booking struct {
	gorm.Model
	ScheduleID uint `json:"schedule_id"`
	SeatID     uint `json:"seat_id"`
	RouteID    uint `json:"route_id"`
	UserID     uint `json:"user_id"`
}

func (Booking) TableName() string {
	return "bookings"
}
