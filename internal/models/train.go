package models

import (
	"gorm.io/gorm"
)

type Train struct {
	gorm.Model
	Name      string     `json:"name"`
	Type      string     `json:"type"`      // e.g., Economy, Business, etc.
	Schedules []Schedule `json:"schedules"` // One-to-many relation with Schedule
}
