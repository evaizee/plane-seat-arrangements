package models

import (
	"time"
)

// Aircraft represents an aircraft in the system
type Aircraft struct {
	ID        string    `json:"id"`
	Code      string    `json:"code"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// Cabin represents a cabin in an aircraft
type Cabin struct {
	ID          string    `json:"id"`
	AircraftID  string    `json:"aircraft_id"`
	SegmentID   string    `json:"segment_id"`
	Deck        string    `json:"deck"`
	FirstRow    int       `json:"first_row"`
	LastRow     int       `json:"last_row"`
	SeatColumns string    `json:"seat_columns"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

// CabinWithSeats represents a cabin with its seats
type CabinWithSeats struct {
	Cabin *Cabin  `json:"cabin"`
	Seats []*Seat `json:"seats"`
}