package models

import (
	"time"
)

// SeatRow represents a row of seats in an aircraft cabin
type SeatRow struct {
	ID        string    `json:"id" db:"id"`
	CabinID   string    `json:"cabin_id" db:"cabin_id"`
	RowNumber int       `json:"row_number" db:"row_number"`
	SeatCodes string    `json:"seat_codes" db:"seat_codes"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
}

// SeatRowWithSeats represents a seat row with its associated seats
type SeatRowWithSeats struct {
	*SeatRow
	Seats []*SeatWithPrice `json:"seats"`
}
