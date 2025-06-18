package postgres

import (
	"database/sql"
	"fmt"
	"strconv"

	"github.com/evaizee/seat-arrangements/backend/models"
	"github.com/evaizee/seat-arrangements/backend/repositories"
)

// FrequentFlyerRepository is a PostgreSQL implementation of the FrequentFlyerRepository interface
type FrequentFlyerRepository struct {
	db *sql.DB
}

// NewFrequentFlyerRepository creates a new PostgreSQL frequent flyer repository
func NewFrequentFlyerRepository(db *sql.DB) repositories.FrequentFlyerRepository {
	return &FrequentFlyerRepository{db: db}
}

// Create inserts a new frequent flyer into the database
func (r *FrequentFlyerRepository) Create(frequentFlyer *models.FrequentFlyer) error {
	query := `
		INSERT INTO frequent_flyers (passenger_id, airline, number, tier_level)
		VALUES ($1, $2, $3, $4)
		RETURNING id, created_at, updated_at
	`

	return r.db.QueryRow(
		query,
		frequentFlyer.PassengerID,
		frequentFlyer.Airline,
		frequentFlyer.Number,
		frequentFlyer.TierLevel,
	).Scan(&frequentFlyer.ID, &frequentFlyer.CreatedAt, &frequentFlyer.UpdatedAt)
}

// GetByID retrieves a frequent flyer by ID
func (r *FrequentFlyerRepository) GetByID(id string) (*models.FrequentFlyer, error) {
	frequentFlyerID, err := strconv.Atoi(id)
	if err != nil {
		return nil, fmt.Errorf("invalid frequent flyer ID: %w", err)
	}

	query := `
		SELECT id, passenger_id, airline, number, tier_level, created_at, updated_at
		FROM frequent_flyers
		WHERE id = $1
	`

	frequentFlyer := &models.FrequentFlyer{}

	err = r.db.QueryRow(query, frequentFlyerID).Scan(
		&frequentFlyer.ID,
		&frequentFlyer.PassengerID,
		&frequentFlyer.Airline,
		&frequentFlyer.Number,
		&frequentFlyer.TierLevel,
		&frequentFlyer.CreatedAt,
		&frequentFlyer.UpdatedAt,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("frequent flyer not found: %w", err)
		}
		return nil, fmt.Errorf("error getting frequent flyer: %w", err)
	}

	return frequentFlyer, nil
}

// GetByPassengerID retrieves frequent flyers by passenger ID
func (r *FrequentFlyerRepository) GetByPassengerID(passengerID string) ([]*models.FrequentFlyer, error) {
	pID, err := strconv.Atoi(passengerID)
	if err != nil {
		return nil, fmt.Errorf("invalid passenger ID: %w", err)
	}

	query := `
		SELECT id, passenger_id, airline, number, tier_level, created_at, updated_at
		FROM frequent_flyers
		WHERE passenger_id = $1
		ORDER BY airline
	`

	rows, err := r.db.Query(query, pID)
	if err != nil {
		return nil, fmt.Errorf("error querying frequent flyers: %w", err)
	}
	defer rows.Close()

	frequentFlyers := []*models.FrequentFlyer{}

	for rows.Next() {
		frequentFlyer := &models.FrequentFlyer{}

		err := rows.Scan(
			&frequentFlyer.ID,
			&frequentFlyer.PassengerID,
			&frequentFlyer.Airline,
			&frequentFlyer.Number,
			&frequentFlyer.TierLevel,
			&frequentFlyer.CreatedAt,
			&frequentFlyer.UpdatedAt,
		)

		if err != nil {
			return nil, fmt.Errorf("error scanning frequent flyer: %w", err)
		}

		frequentFlyers = append(frequentFlyers, frequentFlyer)
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterating frequent flyers: %w", err)
	}

	return frequentFlyers, nil
}

// GetByAirlineAndNumber retrieves a frequent flyer by airline and number
func (r *FrequentFlyerRepository) GetByAirlineAndNumber(airline, number string) (*models.FrequentFlyer, error) {
	query := `
		SELECT id, passenger_id, airline, number, tier_level, created_at, updated_at
		FROM frequent_flyers
		WHERE airline = $1 AND number = $2
	`

	frequentFlyer := &models.FrequentFlyer{}

	err := r.db.QueryRow(query, airline, number).Scan(
		&frequentFlyer.ID,
		&frequentFlyer.PassengerID,
		&frequentFlyer.Airline,
		&frequentFlyer.Number,
		&frequentFlyer.TierLevel,
		&frequentFlyer.CreatedAt,
		&frequentFlyer.UpdatedAt,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("frequent flyer not found: %w", err)
		}
		return nil, fmt.Errorf("error getting frequent flyer: %w", err)
	}

	return frequentFlyer, nil
}

// Update updates a frequent flyer in the database
func (r *FrequentFlyerRepository) Update(frequentFlyer *models.FrequentFlyer) error {
	query := `
		UPDATE frequent_flyers
		SET passenger_id = $1, airline = $2, number = $3, tier_level = $4, updated_at = CURRENT_TIMESTAMP
		WHERE id = $5
		RETURNING updated_at
	`

	return r.db.QueryRow(
		query,
		frequentFlyer.PassengerID,
		frequentFlyer.Airline,
		frequentFlyer.Number,
		frequentFlyer.TierLevel,
		frequentFlyer.ID,
	).Scan(&frequentFlyer.UpdatedAt)
}

// Delete deletes a frequent flyer from the database
func (r *FrequentFlyerRepository) Delete(id string) error {
	frequentFlyerID, err := strconv.Atoi(id)
	if err != nil {
		return fmt.Errorf("invalid frequent flyer ID: %w", err)
	}

	query := "DELETE FROM frequent_flyers WHERE id = $1"
	_, err = r.db.Exec(query, frequentFlyerID)
	return err
}