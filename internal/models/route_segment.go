package models

import "gorm.io/gorm"

type RouteSegment struct {
	gorm.Model
	RouteID        uint   `json:"route_id"`
	StartStationID uint   `json:"start_station_id"`
	EndStationID   uint   `json:"end_station_id"`
	Order          int    `json:"order"`
	Distance       int    `json:"distance"`
	TravelTime     int    `json:"travel_time"`
	Seats          []Seat `gorm:"foreignKey:RouteSegmentID" json:"seats"`
}

func (RouteSegment) TableName() string {
	return "route_segments"
}
