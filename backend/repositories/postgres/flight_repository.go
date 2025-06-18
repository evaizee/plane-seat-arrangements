package postgres

import (
	"database/sql"
	"errors"

	"github.com/evaizee/seat-arrangements/backend/models"
	"github.com/evaizee/seat-arrangements/backend/repositories"
)

// FlightRepository is a PostgreSQL implementation of the FlightRepository interface
type FlightRepository struct {
	db *sql.DB
}

// NewFlightRepository creates a new FlightRepository
func NewFlightRepository(db *sql.DB) repositories.FlightRepository {
	return &FlightRepository{db: db}
}

// Create creates a new flight in the database
func (r *FlightRepository) Create(flight *models.Flight) error {
	query := `
		INSERT INTO flights (
			id, itinerary_id, segment_id, flight_number, airline_code,
			operating_airline, origin, destination, departure, arrival,
			departure_terminal, arrival_terminal, cabin_class, equipment,
			duration, booking_class, created_at, updated_at
		)
		VALUES (
			$1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17, $18
		)
	`

	_, err := r.db.Exec(
		query,
		flight.ID,
		flight.ItineraryID,
		flight.SegmentID,
		flight.FlightNumber,
		flight.AirlineCode,
		flight.OperatingAirline,
		flight.Origin,
		flight.Destination,
		flight.Departure,
		flight.Arrival,
		flight.DepartureTerminal,
		flight.ArrivalTerminal,
		flight.CabinClass,
		flight.Equipment,
		flight.Duration,
		flight.BookingClass,
		flight.CreatedAt,
		flight.UpdatedAt,
	)

	return err
}

// GetByID gets a flight by its ID
func (r *FlightRepository) GetByID(id string) (*models.Flight, error) {
	query := `
		SELECT
			id, itinerary_id, flight_number, operating_airline_code,
			operating_flight_number, origin, destination, departure, arrival,
			departure_terminal, arrival_terminal, cabin_class, equipment,
			duration, booking_class, created_at, updated_at
		FROM segments
		WHERE id = $1
	`

	var flight models.Flight
	err := r.db.QueryRow(query, id).Scan(
		&flight.ID,
		&flight.ItineraryID,
		&flight.FlightNumber,
		&flight.AirlineCode,
		&flight.OperatingAirline,
		&flight.Origin,
		&flight.Destination,
		&flight.Departure,
		&flight.Arrival,
		&flight.DepartureTerminal,
		&flight.ArrivalTerminal,
		&flight.CabinClass,
		&flight.Equipment,
		&flight.Duration,
		&flight.BookingClass,
		&flight.CreatedAt,
		&flight.UpdatedAt,
	)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}

	return &flight, nil
}

// GetAll gets all flights
func (r *FlightRepository) GetAll() ([]*models.Flight, error) {
	query := `
		SELECT
			id, itinerary_id, segment_id, flight_number, airline_code,
			operating_airline, origin, destination, departure, arrival,
			departure_terminal, arrival_terminal, cabin_class, equipment,
			duration, booking_class, created_at, updated_at
		FROM flights
		ORDER BY departure
	`

	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var flights []*models.Flight
	for rows.Next() {
		var flight models.Flight
		err := rows.Scan(
			&flight.ID,
			&flight.ItineraryID,
			&flight.SegmentID,
			&flight.FlightNumber,
			&flight.AirlineCode,
			&flight.OperatingAirline,
			&flight.Origin,
			&flight.Destination,
			&flight.Departure,
			&flight.Arrival,
			&flight.DepartureTerminal,
			&flight.ArrivalTerminal,
			&flight.CabinClass,
			&flight.Equipment,
			&flight.Duration,
			&flight.BookingClass,
			&flight.CreatedAt,
			&flight.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		flights = append(flights, &flight)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return flights, nil
}

// Update updates a flight in the database
func (r *FlightRepository) Update(flight *models.Flight) error {
	query := `
		UPDATE flights
		SET
			itinerary_id = $2,
			segment_id = $3,
			flight_number = $4,
			airline_code = $5,
			operating_airline = $6,
			origin = $7,
			destination = $8,
			departure = $9,
			arrival = $10,
			departure_terminal = $11,
			arrival_terminal = $12,
			cabin_class = $13,
			equipment = $14,
			duration = $15,
			booking_class = $16,
			updated_at = $17
		WHERE id = $1
	`

	_, err := r.db.Exec(
		query,
		flight.ID,
		flight.ItineraryID,
		flight.SegmentID,
		flight.FlightNumber,
		flight.AirlineCode,
		flight.OperatingAirline,
		flight.Origin,
		flight.Destination,
		flight.Departure,
		flight.Arrival,
		flight.DepartureTerminal,
		flight.ArrivalTerminal,
		flight.CabinClass,
		flight.Equipment,
		flight.Duration,
		flight.BookingClass,
		flight.UpdatedAt,
	)

	return err
}

// Delete deletes a flight by its ID
func (r *FlightRepository) Delete(id string) error {
	query := `DELETE FROM flights WHERE id = $1`
	_, err := r.db.Exec(query, id)
	return err
}
