package models

import (
	"gorm.io/gorm"
)

type Train struct {
	gorm.Model
	Name string `json:"name"`
	Type string `json:"type"`
	// Schedules []Schedule `json:"schedules"`
	// Carriages []Carriage `json:"carriages"`
}

func (Train) TableName() string {
	return "trains"
}
