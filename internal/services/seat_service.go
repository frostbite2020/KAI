package services

import (
	"MsKAI/internal/models"

	"gorm.io/gorm"
)

func GetAvailableSeats(db *gorm.DB, carriageID uint) ([]models.Seat, error) {
	var seats []models.Seat
	err := db.Where("carriage_id = ? AND status = ?", carriageID, "Available").Find(&seats).Error
	return seats, err
}

func GetBookedSeats(db *gorm.DB, scheduleID uint) ([]models.Seat, error) {
	var bookedSeats []models.Seat
	err := db.Joins("JOIN booking_seats ON seats.id = booking_seats.seat_id").
		Joins("JOIN bookings ON bookings.id = booking_seats.booking_id").
		Where("bookings.schedule_id = ?", scheduleID).
		Find(&bookedSeats).Error
	return bookedSeats, err
}

func CancelBooking(db *gorm.DB, bookingID uint) error {
	var booking models.Booking
	err := db.Preload("BookedSeats").First(&booking, bookingID).Error
	if err != nil {
		return err
	}

	for _, seat := range booking.BookedSeats {
		err := db.Model(&models.Seat{}).Where("id = ?", seat.ID).Update("status", "Available").Error
		if err != nil {
			return err
		}
	}

	err = db.Delete(&booking).Error
	if err != nil {
		return err
	}

	return nil
}