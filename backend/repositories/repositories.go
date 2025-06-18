package repositories

import (
	"github.com/evaizee/seat-arrangements/backend/models"
)

// UserRepository defines the interface for user data access
type UserRepository interface {
	Create(user *models.User) error
	GetByID(id string) (*models.User, error)
	GetByEmail(email string) (*models.User, error)
	Update(user *models.User) error
	Delete(id string) error
}

// FlightRepository defines the interface for flight data access
type FlightRepository interface {
	Create(flight *models.Flight) error
	GetByID(id string) (*models.Flight, error)
	GetAll() ([]*models.Flight, error)
	Update(flight *models.Flight) error
	Delete(id string) error
}

// AircraftRepository defines the interface for aircraft data access
type AircraftRepository interface {
	Create(aircraft *models.Aircraft) error
	GetByID(id string) (*models.Aircraft, error)
	GetByCode(code string) (*models.Aircraft, error)
	GetAll() ([]*models.Aircraft, error)
	Update(aircraft *models.Aircraft) error
	Delete(id string) error
}

// CabinRepository defines the interface for cabin data access
type CabinRepository interface {
	Create(cabin *models.Cabin) error
	GetByID(id string) (*models.Cabin, error)
	GetByAircraftID(aircraftID string) ([]*models.Cabin, error)
	GetBySegmentID(segmentID string) ([]*models.Cabin, error)
	GetAll() ([]*models.Cabin, error)
	Update(cabin *models.Cabin) error
	Delete(id string) error
}

// SeatRepository defines the interface for seat repository operations
type SeatRepository interface {
	Create(seat *models.Seat) error
	CreatePrice(price *models.SeatPrice) error
	GetByID(id string) (*models.Seat, error)
	GetByRowID(rowID string) ([]*models.Seat, error)
	GetBySegmentID(segmentID string) ([]*models.Seat, error)
	GetByFlightID(flightID string) ([]*models.Seat, error)
	GetAll() ([]*models.Seat, error)
	Update(seat *models.Seat) error
	UpdatePrice(price *models.SeatPrice) error
	Delete(id string) error
	DeletePrice(id string) error
	GetPriceBySeatID(seatID string) (*models.SeatPrice, error)
}

// BookingRepository defines the interface for booking data access
type BookingRepository interface {
	Create(booking *models.Booking) error
	CreateSeat(bookingSeat *models.BookingSeat) error
	GetByID(id string) (*models.Booking, error)
	GetByUserID(userID string) ([]*models.Booking, error)
	GetSeatsByBookingID(bookingID string) ([]*models.BookingSeat, error)
	Update(booking *models.Booking) error
	Delete(id string) error
	DeleteSeats(bookingID string) error
}

// PassengerRepository defines the interface for passenger data access
type PassengerRepository interface {
	Create(passenger *models.Passenger) error
	GetByID(id string) (*models.Passenger, error)
	GetBySegmentID(segmentID string) ([]*models.Passenger, error)
	GetByPassengerNameNumber(nameNumber string) (*models.Passenger, error)
	GetAll() ([]*models.Passenger, error)
	Update(passenger *models.Passenger) error
	Delete(id string) error
}

// FrequentFlyerRepository defines the interface for frequent flyer data access
type FrequentFlyerRepository interface {
	Create(frequentFlyer *models.FrequentFlyer) error
	GetByID(id string) (*models.FrequentFlyer, error)
	GetByPassengerID(passengerID string) ([]*models.FrequentFlyer, error)
	GetByAirlineAndNumber(airline, number string) (*models.FrequentFlyer, error)
	Update(frequentFlyer *models.FrequentFlyer) error
	Delete(id string) error
}

// RowRepository defines the interface for seat row data access
type RowRepository interface {
	GetByCabinID(cabinID string) ([]*models.SeatRow, error)
}