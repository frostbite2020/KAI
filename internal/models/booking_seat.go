package models

type BookingSeat struct {
	BookingID uint `gorm:"primaryKey"`
	SeatID    uint `gorm:"primaryKey"`
}

func (BookingSeat) TableName() string {
	return "booking_seats"
}
