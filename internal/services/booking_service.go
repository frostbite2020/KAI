package services

import (
	"MsKAI/internal/models"

	"gorm.io/gorm"
)

func CreateBooking(db *gorm.DB, booking *models.Booking, seatIDs []uint) error {
	err := db.Create(booking).Error
	if err != nil {
		return err
	}

	for _, seatID := range seatIDs {
		err = db.Model(&models.Seat{}).Where("id = ?", seatID).Update("status", "Booked").Error
		if err != nil {
			return err
		}

		bookingSeat := models.BookingSeat{
			BookingID: booking.ID,
			SeatID:    seatID,
		}
		err = db.Create(&bookingSeat).Error
		if err != nil {
			return err
		}
	}

	return nil
}
