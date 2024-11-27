package models

import (
	"gorm.io/gorm"
)

type Station struct {
	gorm.Model
	Name   string  `json:"name"`
	CityID uint    `json:"city_id"`
	City   City    `gorm:"foreignkey:CityID" json:"city"`
	Routes []Route `gorm:"foreignKey:StartStationID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"routes"`
}

func (Station) TableName() string {
	return "stations"
}
