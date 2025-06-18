package models

import (
	"time"
)

// Flight represents a flight in the system
type Flight struct {
	ID                string    `json:"id"`
	ItineraryID       string    `json:"itinerary_id"`
	SegmentID         string    `json:"segment_id"`
	FlightNumber      string    `json:"flight_number"`
	AirlineCode       string    `json:"airline_code"`
	OperatingAirline  string    `json:"operating_airline"`
	Origin            string    `json:"origin"`
	Destination       string    `json:"destination"`
	Departure         time.Time `json:"departure"`
	Arrival           time.Time `json:"arrival"`
	DepartureTerminal string    `json:"departure_terminal"`
	ArrivalTerminal   string    `json:"arrival_terminal"`
	CabinClass        string    `json:"cabin_class"`
	Equipment         string    `json:"equipment"`
	Duration          int       `json:"duration"`
	BookingClass      string    `json:"booking_class"`
	CreatedAt         time.Time `json:"created_at"`
	UpdatedAt         time.Time `json:"updated_at"`
}

// FlightWithDetails represents a flight with its aircraft and cabin details
type FlightWithDetails struct {
	Flight   *Flight   `json:"flight"`
	Aircraft *Aircraft `json:"aircraft"`
	Cabins   []*Cabin  `json:"cabins"`
}