package models

import (
	"gorm.io/gorm"
)

type Station struct {
	gorm.Model
	Name   string `json:"name"`
	CityID uint   `json:"city_id"`
	City   City   `gorm:"foreignkey:CityID" json:"cities"`
}

func (Station) TableName() string {
	return "stations"
}
