package models

import (
	"time"
)

// Passenger represents a passenger in the system
type Passenger struct {
	ID                int       `json:"id" db:"id"`
	SegmentID         int       `json:"segment_id" db:"segment_id"`
	PassengerIndex    int       `json:"passenger_index" db:"passenger_index"`
	PassengerNameNumber string   `json:"passenger_name_number" db:"passenger_name_number"`
	FirstName         string    `json:"first_name" db:"first_name"`
	LastName          string    `json:"last_name" db:"last_name"`
	DateOfBirth       *time.Time `json:"date_of_birth,omitempty" db:"date_of_birth"`
	Gender            *string   `json:"gender,omitempty" db:"gender"`
	Type              *string   `json:"type,omitempty" db:"type"`
	Email             *string   `json:"email,omitempty" db:"email"`
	Phone             *string   `json:"phone,omitempty" db:"phone"`
	Street            *string   `json:"street,omitempty" db:"street"`
	City              *string   `json:"city,omitempty" db:"city"`
	Country           *string   `json:"country,omitempty" db:"country"`
	Postcode          *string   `json:"postcode,omitempty" db:"postcode"`
	AddressType       *string   `json:"address_type,omitempty" db:"address_type"`
	DocumentType      *string   `json:"document_type,omitempty" db:"document_type"`
	IssuingCountry    *string   `json:"issuing_country,omitempty" db:"issuing_country"`
	CountryOfBirth    *string   `json:"country_of_birth,omitempty" db:"country_of_birth"`
	Nationality       *string   `json:"nationality,omitempty" db:"nationality"`
	CreatedAt         time.Time `json:"created_at" db:"created_at"`
	UpdatedAt         time.Time `json:"updated_at" db:"updated_at"`
}

// FrequentFlyer represents a frequent flyer record for a passenger
type FrequentFlyer struct {
	ID          int       `json:"id" db:"id"`
	PassengerID int       `json:"passenger_id" db:"passenger_id"`
	Airline     string    `json:"airline" db:"airline"`
	Number      string    `json:"number" db:"number"`
	TierLevel   *string   `json:"tier_level,omitempty" db:"tier_level"`
	CreatedAt   time.Time `json:"created_at" db:"created_at"`
	UpdatedAt   time.Time `json:"updated_at" db:"updated_at"`
}

// PassengerWithDetails represents a passenger with their frequent flyer details
type PassengerWithDetails struct {
	Passenger     Passenger      `json:"passenger"`
	FrequentFlyers []*FrequentFlyer `json:"frequent_flyers,omitempty"`
}

// NewPassenger creates a new passenger
func NewPassenger(segmentID, passengerIndex int, passengerNameNumber, firstName, lastName string) *Passenger {
	return &Passenger{
		SegmentID:         segmentID,
		PassengerIndex:    passengerIndex,
		PassengerNameNumber: passengerNameNumber,
		FirstName:         firstName,
		LastName:          lastName,
		CreatedAt:         time.Now(),
		UpdatedAt:         time.Now(),
	}
}

// NewFrequentFlyer creates a new frequent flyer record
func NewFrequentFlyer(passengerID int, airline, number string) *FrequentFlyer {
	return &FrequentFlyer{
		PassengerID: passengerID,
		Airline:     airline,
		Number:      number,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}
}