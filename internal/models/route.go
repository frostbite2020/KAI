package models

import (
	"gorm.io/gorm"
)

type Route struct {
	gorm.Model
	StartStationID uint    `json:"start_station_id"`
	EndStationID   uint    `json:"end_station_id"`
	Distance       int     `json:"distance"`
	TravelTime     int     `json:"travel_time"`
	StartStation   Station `gorm:"foreignkey:StartStationID" json:"start_station"`
	EndStation     Station `gorm:"foreignkey:EndStationID" json:"end_station"`
}

func (Route) TableName() string {
	return "routes"
}
