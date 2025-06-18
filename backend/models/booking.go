package models

import (
	"time"

	"github.com/google/uuid"
)

// Booking represents a booking in the system
type Booking struct {
	ID        string    `json:"id"`
	UserID    string    `json:"user_id"`
	FlightID  string    `json:"flight_id"`
	Status    string    `json:"status"` // confirmed, cancelled, pending
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// BookingSeat represents a seat in a booking
type BookingSeat struct {
	ID        string    `json:"id"`
	BookingID string    `json:"booking_id"`
	SeatID    string    `json:"seat_id"`
	Price     float64   `json:"price"`
	Currency  string    `json:"currency"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// BookingWithDetails represents a booking with its flight and seats
type BookingWithDetails struct {
	Booking *Booking      `json:"booking"`
	Flight  *Flight       `json:"flight"`
	Seats   []*BookingSeat `json:"seats"`
	Total   float64       `json:"total"`
}

// NewBooking creates a new booking
func NewBooking(userID, flightID string) *Booking {
	return &Booking{
		ID:        uuid.New().String(),
		UserID:    userID,
		FlightID:  flightID,
		Status:    "confirmed",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
}

// NewBookingSeat creates a new booking seat
func NewBookingSeat(bookingID, seatID string, price float64, currency string) *BookingSeat {
	return &BookingSeat{
		ID:        uuid.New().String(),
		BookingID: bookingID,
		SeatID:    seatID,
		Price:     price,
		Currency:  currency,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
}