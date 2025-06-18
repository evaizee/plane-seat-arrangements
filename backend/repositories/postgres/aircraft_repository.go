package postgres

import (
	"database/sql"
	"errors"

	"github.com/evaizee/seat-arrangements/backend/models"
	"github.com/evaizee/seat-arrangements/backend/repositories"
)

// AircraftRepository is a PostgreSQL implementation of the AircraftRepository interface
type AircraftRepository struct {
	db *sql.DB
}

// NewAircraftRepository creates a new AircraftRepository
func NewAircraftRepository(db *sql.DB) repositories.AircraftRepository {
	return &AircraftRepository{db: db}
}

// Create creates a new aircraft in the database
func (r *AircraftRepository) Create(aircraft *models.Aircraft) error {
	query := `
		INSERT INTO aircraft (id, code, name, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5)
	`

	_, err := r.db.Exec(
		query,
		aircraft.ID,
		aircraft.Code,
		aircraft.Name,
		aircraft.CreatedAt,
		aircraft.UpdatedAt,
	)

	return err
}

// GetByID gets an aircraft by its ID
func (r *AircraftRepository) GetByID(id string) (*models.Aircraft, error) {
	query := `
		SELECT id, code, name, created_at, updated_at
		FROM aircraft
		WHERE id = $1
	`

	var aircraft models.Aircraft
	err := r.db.QueryRow(query, id).Scan(
		&aircraft.ID,
		&aircraft.Code,
		&aircraft.Name,
		&aircraft.CreatedAt,
		&aircraft.UpdatedAt,
	)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}

	return &aircraft, nil
}

// GetByCode gets an aircraft by its code
func (r *AircraftRepository) GetByCode(code string) (*models.Aircraft, error) {
	query := `
		SELECT id, code, name, created_at, updated_at
		FROM aircraft
		WHERE code = $1
	`

	var aircraft models.Aircraft
	err := r.db.QueryRow(query, code).Scan(
		&aircraft.ID,
		&aircraft.Code,
		&aircraft.Name,
		&aircraft.CreatedAt,
		&aircraft.UpdatedAt,
	)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}

	return &aircraft, nil
}

// GetAll gets all aircraft
func (r *AircraftRepository) GetAll() ([]*models.Aircraft, error) {
	query := `
		SELECT id, code, name, created_at, updated_at
		FROM aircraft
		ORDER BY code
	`

	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var aircraft []*models.Aircraft
	for rows.Next() {
		var a models.Aircraft
		err := rows.Scan(
			&a.ID,
			&a.Code,
			&a.Name,
			&a.CreatedAt,
			&a.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		aircraft = append(aircraft, &a)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return aircraft, nil
}

// Update updates an aircraft in the database
func (r *AircraftRepository) Update(aircraft *models.Aircraft) error {
	query := `
		UPDATE aircraft
		SET code = $2, name = $3, updated_at = $4
		WHERE id = $1
	`

	_, err := r.db.Exec(
		query,
		aircraft.ID,
		aircraft.Code,
		aircraft.Name,
		aircraft.UpdatedAt,
	)

	return err
}

// Delete deletes an aircraft by its ID
func (r *AircraftRepository) Delete(id string) error {
	query := `DELETE FROM aircraft WHERE id = $1`
	_, err := r.db.Exec(query, id)
	return err
}
