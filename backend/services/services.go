package services

import (
	"github.com/evaizee/seat-arrangements/backend/models"
)

// UserService defines the interface for user business logic
type UserService interface {
	Register(email, password, firstName, lastName string) (*models.User, error)
	GetByID(id string) (*models.User, error)
	GetByEmail(email string) (*models.User, error)
	UpdateUser(user *models.User) error
	ChangePassword(userID, currentPassword, newPassword string) error
}

// FlightService defines the interface for flight business logic
type FlightService interface {
	GetByID(id string) (*models.Flight, error)
	GetAll() ([]*models.Flight, error)
	GetWithDetails(id string) (*models.FlightWithDetails, error)
}

// AircraftService defines the interface for aircraft business logic
type AircraftService interface {
	GetByID(id string) (*models.Aircraft, error)
	GetAll() ([]*models.Aircraft, error)
}

// CabinService defines the interface for cabin business logic
type CabinService interface {
	GetByID(id string) (*models.Cabin, error)
	GetByAircraftID(aircraftID string) ([]*models.Cabin, error)
	GetWithSeats(id string) (*models.CabinWithSeats, error)
}

// SeatService defines the interface for seat business logic
type SeatService interface {
	GetByID(id string) (*models.SeatWithPrice, error)
	GetByRowID(rowID string) ([]*models.SeatWithPrice, error)
	GetByFlightID(flightID string) ([]*models.SeatWithPrice, error)
	GetSeatMap(flightID string, passengerID string) (*models.SeatMapResponse, error)
	UpdateAvailability(seatID string, available bool) error
}

// BookingService defines the interface for booking business logic
type BookingService interface {
	CreateBooking(userID, flightID string, seatIDs []string) (*models.BookingWithDetails, error)
	GetByID(id string) (*models.BookingWithDetails, error)
	GetByUserID(userID string) ([]*models.BookingWithDetails, error)
	CancelBooking(id string) error
}

// AuthService defines the interface for authentication business logic
type AuthService interface {
	Register(email, password, firstName, lastName string) (*models.User, error)
	Login(email, password string) (string, error) // Returns JWT token
	ValidateToken(token string) (string, error)   // Returns user ID if valid
}

// PassengerService defines the interface for passenger business logic
type PassengerService interface {
	CreatePassenger(segmentID int, passengerIndex int, passengerNameNumber, firstName, lastName string) (*models.Passenger, error)
	GetByID(id string) (*models.Passenger, error)
	GetBySegmentID(segmentID string) ([]*models.Passenger, error)
	GetWithFrequentFlyers(id string) (*models.PassengerWithDetails, error)
	UpdatePassenger(passenger *models.Passenger) error
}

// FrequentFlyerService defines the interface for frequent flyer business logic
type FrequentFlyerService interface {
	CreateFrequentFlyer(passengerID int, airline, number string) (*models.FrequentFlyer, error)
	GetByID(id string) (*models.FrequentFlyer, error)
	GetByPassengerID(passengerID string) ([]*models.FrequentFlyer, error)
	UpdateFrequentFlyer(frequentFlyer *models.FrequentFlyer) error
}
