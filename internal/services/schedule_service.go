package services

import (
	"MsKAI/internal/models"
	"encoding/json"
	"net/http"
	"time"

	"gorm.io/gorm"
)

func CreateSchedule(w http.ResponseWriter, r *http.Request, db *gorm.DB) {
	var schedule models.Schedule

	if err := json.NewDecoder(r.Body).Decode(&schedule); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	var route models.Route
	if err := db.First(&route, schedule.RouteID).Error; err != nil {
		http.Error(w, "Route not found", http.StatusNotFound)
		return
	}

	// Parse departure time from the request
	departureTime, err := time.Parse(time.RFC3339, schedule.Departure)
	if err != nil {
		http.Error(w, "Invalid departure time format", http.StatusBadRequest)
		return
	}

	arrivalTime := departureTime.Add(time.Duration(route.TravelTime) * time.Second)

	schedule.Arrival = arrivalTime.Format(time.RFC3339)

	if err := db.Create(&schedule).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Return the created schedule
	json.NewEncoder(w).Encode(schedule)
}

func GetSchedules(w http.ResponseWriter, r *http.Request, db *gorm.DB) {
	var schedules []models.Schedule

	// Get query parameters
	routeID := r.URL.Query().Get("route_id")
	departure := r.URL.Query().Get("departure")

	// Validate required parameters
	if routeID == "" || departure == "" {
		http.Error(w, "route_id and departure query parameters are required", http.StatusBadRequest)
		return
	}

	// Query schedules with filters
	err := db.Preload("Train").
		Preload("Route").
		Preload("Route.StartStation").
		Preload("Route.StartStation.City").
		Preload("Route.EndStation").
		Preload("Route.EndStation.City").
		Where("route_id = ? AND departure > ?", routeID, departure).
		Find(&schedules).Error

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Return the schedules as JSON
	json.NewEncoder(w).Encode(schedules)
}
