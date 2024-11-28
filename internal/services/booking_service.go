package services

import (
	"MsKAI/internal/models"
	"encoding/json"
	"net/http"

	"gorm.io/gorm"
)

type CreateBookingRequest struct {
	ScheduleID   uint   `json:"schedule_id"`
	UserID       uint   `json:"user_id"`
	CarriageType string `json:"carriage_type"` // Carriage type like "economy", "business"
}

func CreateBooking(w http.ResponseWriter, r *http.Request, db *gorm.DB) {
	var bookingReq CreateBookingRequest
	if err := json.NewDecoder(r.Body).Decode(&bookingReq); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	if bookingReq.ScheduleID == 0 || bookingReq.UserID == 0 || bookingReq.CarriageType == "" {
		http.Error(w, "schedule_id, user_id, and carriage_type are required", http.StatusBadRequest)
		return
	}

	var carriage models.Carriage
	if err := db.Preload("Seats").Where("train_id = ? AND type = ?", bookingReq.ScheduleID, bookingReq.CarriageType).First(&carriage).Error; err != nil {
		http.Error(w, "No carriages available for this schedule with type "+bookingReq.CarriageType, http.StatusNotFound)
		return
	}

	var seat models.Seat
	if err := db.Where("carriage_id = ? AND booked = ?", carriage.ID, false).First(&seat).Error; err != nil {
		http.Error(w, "No available seats in this carriage", http.StatusConflict)
		return
	}

	// Mark the seat as booked
	seat.Booked = true
	if err := db.Save(&seat).Error; err != nil {
		http.Error(w, "Error updating seat status", http.StatusInternalServerError)
		return
	}

	booking := models.Booking{
		ScheduleID: bookingReq.ScheduleID,
		UserID:     bookingReq.UserID,
		SeatID:     seat.ID,
	}

	if err := db.Create(&booking).Error; err != nil {
		http.Error(w, "Error creating booking", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(booking)
}
