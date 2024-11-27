package models

import (
	"gorm.io/gorm"
)

type City struct {
	gorm.Model
	Name     string    `json:"name"`
	Stations []Station `json:"stations"`
}

func (City) TableName() string {
	return "cities"
}
