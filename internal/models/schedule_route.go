package models

import (
	"gorm.io/gorm"
)

type ScheduleRoute struct {
	gorm.Model
	ScheduleID uint     `json:"schedule_id"`
	Schedule   Schedule `gorm:"foreignKey:ScheduleID" json:"schedule"`
	RouteID    uint     `json:"route_id"`
	Route      Route    `gorm:"foreignKey:RouteID" json:"route"`
	Order      uint     `json:"order"`
}

func (ScheduleRoute) TableName() string {
	return "schedule_routes"
}
