package postgres

import (
	"database/sql"
	"fmt"
	"strconv"

	"github.com/evaizee/seat-arrangements/backend/models"
	"github.com/evaizee/seat-arrangements/backend/repositories"
)

// PassengerRepository is a PostgreSQL implementation of the PassengerRepository interface
type PassengerRepository struct {
	db *sql.DB
}

// NewPassengerRepository creates a new PostgreSQL passenger repository
func NewPassengerRepository(db *sql.DB) repositories.PassengerRepository {
	return &PassengerRepository{db: db}
}

// Create inserts a new passenger into the database
func (r *PassengerRepository) Create(passenger *models.Passenger) error {
	query := `
		INSERT INTO passengers (
			segment_id, passenger_index, passenger_name_number, first_name, last_name, 
			date_of_birth, gender, type, email, phone, street, city, country, postcode, 
			address_type, document_type, issuing_country, country_of_birth, nationality
		) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17, $18, $19)
		RETURNING id, created_at, updated_at
	`

	var dateOfBirth *string
	if passenger.DateOfBirth != nil {
		dateStr := passenger.DateOfBirth.Format("2006-01-02")
		dateOfBirth = &dateStr
	}

	return r.db.QueryRow(
		query,
		passenger.SegmentID,
		passenger.PassengerIndex,
		passenger.PassengerNameNumber,
		passenger.FirstName,
		passenger.LastName,
		dateOfBirth,
		passenger.Gender,
		passenger.Type,
		passenger.Email,
		passenger.Phone,
		passenger.Street,
		passenger.City,
		passenger.Country,
		passenger.Postcode,
		passenger.AddressType,
		passenger.DocumentType,
		passenger.IssuingCountry,
		passenger.CountryOfBirth,
		passenger.Nationality,
	).Scan(&passenger.ID, &passenger.CreatedAt, &passenger.UpdatedAt)
}

// GetByID retrieves a passenger by ID
func (r *PassengerRepository) GetByID(id string) (*models.Passenger, error) {
	passengerID, err := strconv.Atoi(id)
	if err != nil {
		return nil, fmt.Errorf("invalid passenger ID: %w", err)
	}

	query := `
		SELECT id, segment_id, passenger_index, passenger_name_number, first_name, last_name, 
			date_of_birth, gender, type, email, phone, street, city, country, postcode, 
			address_type, document_type, issuing_country, country_of_birth, nationality, 
			created_at, updated_at
		FROM passengers
		WHERE id = $1
	`

	passenger := &models.Passenger{}
	var dateOfBirth sql.NullTime

	err = r.db.QueryRow(query, passengerID).Scan(
		&passenger.ID,
		&passenger.SegmentID,
		&passenger.PassengerIndex,
		&passenger.PassengerNameNumber,
		&passenger.FirstName,
		&passenger.LastName,
		&dateOfBirth,
		&passenger.Gender,
		&passenger.Type,
		&passenger.Email,
		&passenger.Phone,
		&passenger.Street,
		&passenger.City,
		&passenger.Country,
		&passenger.Postcode,
		&passenger.AddressType,
		&passenger.DocumentType,
		&passenger.IssuingCountry,
		&passenger.CountryOfBirth,
		&passenger.Nationality,
		&passenger.CreatedAt,
		&passenger.UpdatedAt,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("passenger not found: %w", err)
		}
		return nil, fmt.Errorf("error getting passenger: %w", err)
	}

	if dateOfBirth.Valid {
		passenger.DateOfBirth = &dateOfBirth.Time
	}

	return passenger, nil
}

// GetBySegmentID retrieves passengers by segment ID
func (r *PassengerRepository) GetBySegmentID(segmentID string) ([]*models.Passenger, error) {
	segID, err := strconv.Atoi(segmentID)
	if err != nil {
		return nil, fmt.Errorf("invalid segment ID: %w", err)
	}

	query := `
		SELECT id, segment_id, passenger_index, passenger_name_number, first_name, last_name, 
			date_of_birth, gender, type, email, phone, street, city, country, postcode, 
			address_type, document_type, issuing_country, country_of_birth, nationality, 
			created_at, updated_at
		FROM passengers
		WHERE segment_id = $1
		ORDER BY passenger_index
	`

	rows, err := r.db.Query(query, segID)
	if err != nil {
		return nil, fmt.Errorf("error querying passengers: %w", err)
	}
	defer rows.Close()

	passengers := []*models.Passenger{}

	for rows.Next() {
		passenger := &models.Passenger{}
		var dateOfBirth sql.NullTime

		err := rows.Scan(
			&passenger.ID,
			&passenger.SegmentID,
			&passenger.PassengerIndex,
			&passenger.PassengerNameNumber,
			&passenger.FirstName,
			&passenger.LastName,
			&dateOfBirth,
			&passenger.Gender,
			&passenger.Type,
			&passenger.Email,
			&passenger.Phone,
			&passenger.Street,
			&passenger.City,
			&passenger.Country,
			&passenger.Postcode,
			&passenger.AddressType,
			&passenger.DocumentType,
			&passenger.IssuingCountry,
			&passenger.CountryOfBirth,
			&passenger.Nationality,
			&passenger.CreatedAt,
			&passenger.UpdatedAt,
		)

		if err != nil {
			return nil, fmt.Errorf("error scanning passenger: %w", err)
		}

		if dateOfBirth.Valid {
			passenger.DateOfBirth = &dateOfBirth.Time
		}

		passengers = append(passengers, passenger)
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterating passengers: %w", err)
	}

	return passengers, nil
}

// GetByPassengerNameNumber retrieves a passenger by name number
func (r *PassengerRepository) GetByPassengerNameNumber(nameNumber string) (*models.Passenger, error) {
	query := `
		SELECT id, segment_id, passenger_index, passenger_name_number, first_name, last_name, 
			date_of_birth, gender, type, email, phone, street, city, country, postcode, 
			address_type, document_type, issuing_country, country_of_birth, nationality, 
			created_at, updated_at
		FROM passengers
		WHERE passenger_name_number = $1
	`

	passenger := &models.Passenger{}
	var dateOfBirth sql.NullTime

	err := r.db.QueryRow(query, nameNumber).Scan(
		&passenger.ID,
		&passenger.SegmentID,
		&passenger.PassengerIndex,
		&passenger.PassengerNameNumber,
		&passenger.FirstName,
		&passenger.LastName,
		&dateOfBirth,
		&passenger.Gender,
		&passenger.Type,
		&passenger.Email,
		&passenger.Phone,
		&passenger.Street,
		&passenger.City,
		&passenger.Country,
		&passenger.Postcode,
		&passenger.AddressType,
		&passenger.DocumentType,
		&passenger.IssuingCountry,
		&passenger.CountryOfBirth,
		&passenger.Nationality,
		&passenger.CreatedAt,
		&passenger.UpdatedAt,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("passenger not found: %w", err)
		}
		return nil, fmt.Errorf("error getting passenger: %w", err)
	}

	if dateOfBirth.Valid {
		passenger.DateOfBirth = &dateOfBirth.Time
	}

	return passenger, nil
}

// GetAll retrieves all passengers
func (r *PassengerRepository) GetAll() ([]*models.Passenger, error) {
	query := `
		SELECT id, segment_id, passenger_index, passenger_name_number, first_name, last_name, 
			date_of_birth, gender, type, email, phone, street, city, country, postcode, 
			address_type, document_type, issuing_country, country_of_birth, nationality, 
			created_at, updated_at
		FROM passengers
		ORDER BY segment_id, passenger_index
	`

	rows, err := r.db.Query(query)
	if err != nil {
		return nil, fmt.Errorf("error querying passengers: %w", err)
	}
	defer rows.Close()

	passengers := []*models.Passenger{}

	for rows.Next() {
		passenger := &models.Passenger{}
		var dateOfBirth sql.NullTime

		err := rows.Scan(
			&passenger.ID,
			&passenger.SegmentID,
			&passenger.PassengerIndex,
			&passenger.PassengerNameNumber,
			&passenger.FirstName,
			&passenger.LastName,
			&dateOfBirth,
			&passenger.Gender,
			&passenger.Type,
			&passenger.Email,
			&passenger.Phone,
			&passenger.Street,
			&passenger.City,
			&passenger.Country,
			&passenger.Postcode,
			&passenger.AddressType,
			&passenger.DocumentType,
			&passenger.IssuingCountry,
			&passenger.CountryOfBirth,
			&passenger.Nationality,
			&passenger.CreatedAt,
			&passenger.UpdatedAt,
		)

		if err != nil {
			return nil, fmt.Errorf("error scanning passenger: %w", err)
		}

		if dateOfBirth.Valid {
			passenger.DateOfBirth = &dateOfBirth.Time
		}

		passengers = append(passengers, passenger)
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterating passengers: %w", err)
	}

	return passengers, nil
}

// Update updates a passenger in the database
func (r *PassengerRepository) Update(passenger *models.Passenger) error {
	query := `
		UPDATE passengers
		SET segment_id = $1, passenger_index = $2, passenger_name_number = $3, 
			first_name = $4, last_name = $5, date_of_birth = $6, gender = $7, 
			type = $8, email = $9, phone = $10, street = $11, city = $12, 
			country = $13, postcode = $14, address_type = $15, document_type = $16, 
			issuing_country = $17, country_of_birth = $18, nationality = $19, 
			updated_at = CURRENT_TIMESTAMP
		WHERE id = $20
		RETURNING updated_at
	`

	var dateOfBirth *string
	if passenger.DateOfBirth != nil {
		dateStr := passenger.DateOfBirth.Format("2006-01-02")
		dateOfBirth = &dateStr
	}

	return r.db.QueryRow(
		query,
		passenger.SegmentID,
		passenger.PassengerIndex,
		passenger.PassengerNameNumber,
		passenger.FirstName,
		passenger.LastName,
		dateOfBirth,
		passenger.Gender,
		passenger.Type,
		passenger.Email,
		passenger.Phone,
		passenger.Street,
		passenger.City,
		passenger.Country,
		passenger.Postcode,
		passenger.AddressType,
		passenger.DocumentType,
		passenger.IssuingCountry,
		passenger.CountryOfBirth,
		passenger.Nationality,
		passenger.ID,
	).Scan(&passenger.UpdatedAt)
}

// Delete deletes a passenger from the database
func (r *PassengerRepository) Delete(id string) error {
	passengerID, err := strconv.Atoi(id)
	if err != nil {
		return fmt.Errorf("invalid passenger ID: %w", err)
	}

	query := "DELETE FROM passengers WHERE id = $1"
	_, err = r.db.Exec(query, passengerID)
	return err
}