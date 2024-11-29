package services

import (
	"MsKAI/internal/models"
	"encoding/json"
	"net/http"
	"time"

	"gorm.io/gorm"
)

// func CreateSchedule(w http.ResponseWriter, r *http.Request, db *gorm.DB) {
// 	type CreateScheduleRequest struct {
// 		TrainID   uint   `json:"train_id"`
// 		Departure string `json:"departure"`
// 		Routes    []uint `json:"routes"` // ID rute yang dilewati
// 	}

// 	var req CreateScheduleRequest
// 	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
// 		http.Error(w, err.Error(), http.StatusBadRequest)
// 		return
// 	}

// 	departureTime, err := time.Parse(time.RFC3339, req.Departure)
// 	if err != nil {
// 		http.Error(w, "Invalid departure time", http.StatusBadRequest)
// 		return
// 	}

// 	schedule := models.Schedule{
// 		TrainID:   req.TrainID,
// 		Departure: departureTime,
// 	}

// 	if err := db.Create(&schedule).Error; err != nil {
// 		http.Error(w, err.Error(), http.StatusInternalServerError)
// 		return
// 	}

// 	for i, routeID := range req.Routes {
// 		scheduleRoute := models.ScheduleRoute{
// 			ScheduleID: schedule.ID,
// 			RouteID:    routeID,
// 			Order:      i + 1,
// 		}
// 		if err := db.Create(&scheduleRoute).Error; err != nil {
// 			http.Error(w, err.Error(), http.StatusInternalServerError)
// 			return
// 		}
// 	}

// 	json.NewEncoder(w).Encode(schedule)
// }

func GetSchedules(w http.ResponseWriter, r *http.Request, db *gorm.DB) {
	startStationName := r.URL.Query().Get("startstation")
	endStationName := r.URL.Query().Get("endstation")
	departureDate := r.URL.Query().Get("departuredate")

	if startStationName == "" || endStationName == "" || departureDate == "" {
		http.Error(w, "Missing required parameters", http.StatusBadRequest)
		return
	}

	var startStation, endStation models.Station
	if err := db.Where("name = ?", startStationName).First(&startStation).Error; err != nil {
		http.Error(w, "Start station not found", http.StatusNotFound)
		return
	}
	if err := db.Where("name = ?", endStationName).First(&endStation).Error; err != nil {
		http.Error(w, "End station not found", http.StatusNotFound)
		return
	}

	// Fetch train schedules
	var schedule models.Schedule
	err := db.Preload("Train").
		Preload("ScheduleRoutes.Route.StartStation.City").
		Preload("ScheduleRoutes.Route.EndStation.City").
		Joins("JOIN schedule_routes sr_start ON sr_start.schedule_id = schedules.id").
		Joins("JOIN routes r_start ON r_start.id = sr_start.route_id").
		Joins("JOIN stations start_station ON start_station.id = r_start.start_station_id").
		Where("start_station.id = ?", startStation.ID).
		First(&schedule).Error

	if err != nil {
		http.Error(w, "No schedule available for the requested route", http.StatusNotFound)
		return
	}

	// variabel untuk menghubungkan rute
	var currentStation = startStation
	var totalDistance, totalTravelTime int
	var connectedRoutes []models.Route
	var firstRoute models.Route

	for {
		var nextRoute models.Route
		err := db.Preload("StartStation.City").
			Preload("EndStation.City").
			Where("start_station_id = ?", currentStation.ID).
			First(&nextRoute).Error
		if err != nil {
			http.Error(w, "Failed to connect all routes", http.StatusNotFound)
			return
		}

		// Simpan rute pertama
		if len(connectedRoutes) == 0 {
			firstRoute = nextRoute
		}

		connectedRoutes = append(connectedRoutes, nextRoute)
		totalDistance += nextRoute.Distance
		totalTravelTime += nextRoute.TravelTime
		currentStation = nextRoute.EndStation

		if currentStation.ID == endStation.ID {
			break
		}
	}

	// Hitung waktu tiba
	departureTime, _ := time.Parse(time.RFC3339, departureDate)
	arrivalTime := departureTime.Add(time.Duration(totalTravelTime) * time.Minute)

	// respons
	type RouteResponse struct {
		Distance     int            `json:"distance"`
		TravelTime   int            `json:"travel_time"`
		StartStation models.Station `json:"start_station"`
		EndStation   models.Station `json:"end_station"`
	}
	type ScheduleRouteResponse struct {
		ScheduleID     uint          `json:"schedule_id"`
		StartStationID uint          `json:"start_station_id"`
		EndStationID   uint          `json:"end_station_id"`
		Route          RouteResponse `json:"route"`
	}
	type TrainResponse struct {
		ID   uint   `json:"ID"`
		Name string `json:"name"`
		Type string `json:"type"`
	}
	type ScheduleResponse struct {
		ID             uint                  `json:"ID"`
		CreatedAt      time.Time             `json:"CreatedAt"`
		TrainID        uint                  `json:"train_id"`
		Train          TrainResponse         `json:"train"`
		Departure      time.Time             `json:"departure"`
		ArrivalTime    time.Time             `json:"arrival_time"`
		ScheduleRoutes ScheduleRouteResponse `json:"schedule_routes"`
	}

	response := ScheduleResponse{
		ID:        schedule.ID,
		CreatedAt: schedule.CreatedAt,
		TrainID:   schedule.TrainID,
		Train: TrainResponse{
			ID:   schedule.Train.ID,
			Name: schedule.Train.Name,
			Type: schedule.Train.Type,
		},
		Departure:   departureTime,
		ArrivalTime: arrivalTime,
		ScheduleRoutes: ScheduleRouteResponse{
			ScheduleID:     schedule.ID,
			StartStationID: startStation.ID,
			EndStationID:   endStation.ID,
			Route: RouteResponse{
				Distance:     totalDistance,
				TravelTime:   totalTravelTime,
				StartStation: firstRoute.StartStation,
				EndStation:   currentStation,
			},
		},
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
