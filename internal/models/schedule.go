package models

import (
	"time"

	"gorm.io/gorm"
)

type Schedule struct {
	gorm.Model
	TrainID        uint            `json:"train_id"`
	Train          Train           `gorm:"foreignKey:TrainID" json:"train"`
	Departure      time.Time       `json:"departure"`
	ScheduleRoutes []ScheduleRoute `gorm:"foreignKey:ScheduleID" json:"schedule_routes"`
}

func (Schedule) TableName() string {
	return "schedules"
}
