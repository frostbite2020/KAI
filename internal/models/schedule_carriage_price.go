package models

import "gorm.io/gorm"

type ScheduleCarriagePrice struct {
	gorm.Model
	ScheduleID   uint     `json:"schedule_id"`
	Schedule     Schedule `gorm:"foreignKey:ScheduleID" json:"schedule"`
	CarriageID   uint     `json:"carriage_id"`
	Carriage     Carriage `gorm:"foreignKey:CarriageID" json:"carriage"`
	TrainType    string   `json:"train_type"`
	CarriageType string   `json:"carriage_type"`
	Price        float64  `json:"price"`
}

func (ScheduleCarriagePrice) TableName() string {
	return "schedule_carriage_prices"
}
