package models

import "gorm.io/gorm"

// ScheduleCarriagePrice model
type ScheduleCarriagePrice struct {
	gorm.Model
	ScheduleID uint     `json:"schedule_id"`
	Schedule   Schedule `gorm:"foreignKey:ScheduleID" json:"schedule"`
	CarriageID uint     `json:"carriage_id"`
	Carriage   Carriage `gorm:"foreignKey:CarriageID" json:"carriage"`
	Price      float64  `json:"price"`
}

func (ScheduleCarriagePrice) TableName() string {
	return "schedule_carriage_prices"
}
