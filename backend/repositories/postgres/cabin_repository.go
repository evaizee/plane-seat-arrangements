package postgres

import (
	"database/sql"
	"errors"

	"github.com/evaizee/seat-arrangements/backend/models"
	"github.com/evaizee/seat-arrangements/backend/repositories"
)

// CabinRepository is a PostgreSQL implementation of the CabinRepository interface
type CabinRepository struct {
	db *sql.DB
}

// NewCabinRepository creates a new CabinRepository
func NewCabinRepository(db *sql.DB) repositories.CabinRepository {
	return &CabinRepository{db: db}
}

// Create creates a new cabin in the database
func (r *CabinRepository) Create(cabin *models.Cabin) error {
	query := `
		INSERT INTO cabins (
			id, aircraft_id, segment_id, deck, first_row, last_row, 
			seat_columns, created_at, updated_at
		)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)
	`

	_, err := r.db.Exec(
		query,
		cabin.ID,
		cabin.AircraftID,
		cabin.SegmentID,
		cabin.Deck,
		cabin.FirstRow,
		cabin.LastRow,
		cabin.SeatColumns,
		cabin.CreatedAt,
		cabin.UpdatedAt,
	)

	return err
}

// GetByID gets a cabin by its ID
func (r *CabinRepository) GetByID(id string) (*models.Cabin, error) {
	query := `
		SELECT id, aircraft_id, segment_id, deck, first_row, last_row, 
			seat_columns, created_at, updated_at
		FROM cabins
		WHERE id = $1
	`

	var cabin models.Cabin
	err := r.db.QueryRow(query, id).Scan(
		&cabin.ID,
		&cabin.AircraftID,
		&cabin.SegmentID,
		&cabin.Deck,
		&cabin.FirstRow,
		&cabin.LastRow,
		&cabin.SeatColumns,
		&cabin.CreatedAt,
		&cabin.UpdatedAt,
	)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}

	return &cabin, nil
}

// GetByAircraftID gets cabins by aircraft ID
func (r *CabinRepository) GetByAircraftID(aircraftID string) ([]*models.Cabin, error) {
	query := `
		SELECT id, aircraft_id, segment_id, deck, first_row, last_row, 
			seat_columns, created_at, updated_at
		FROM cabins
		WHERE aircraft_id = $1
		ORDER BY deck, first_row
	`

	rows, err := r.db.Query(query, aircraftID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var cabins []*models.Cabin
	for rows.Next() {
		var cabin models.Cabin
		err := rows.Scan(
			&cabin.ID,
			&cabin.AircraftID,
			&cabin.SegmentID,
			&cabin.Deck,
			&cabin.FirstRow,
			&cabin.LastRow,
			&cabin.SeatColumns,
			&cabin.CreatedAt,
			&cabin.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		cabins = append(cabins, &cabin)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return cabins, nil
}

// GetBySegmentID gets cabins by segment ID
func (r *CabinRepository) GetBySegmentID(segmentID string) ([]*models.Cabin, error) {
	query := `
		SELECT id, aircraft_id, segment_id, deck, first_row, last_row, 
			seat_columns, created_at, updated_at
		FROM cabins
		WHERE segment_id = $1
		ORDER BY deck, first_row
	`

	rows, err := r.db.Query(query, segmentID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var cabins []*models.Cabin
	for rows.Next() {
		var cabin models.Cabin
		err := rows.Scan(
			&cabin.ID,
			&cabin.AircraftID,
			&cabin.SegmentID,
			&cabin.Deck,
			&cabin.FirstRow,
			&cabin.LastRow,
			&cabin.SeatColumns,
			&cabin.CreatedAt,
			&cabin.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		cabins = append(cabins, &cabin)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return cabins, nil
}

// GetAll gets all cabins
func (r *CabinRepository) GetAll() ([]*models.Cabin, error) {
	query := `
		SELECT id, aircraft_id, segment_id, deck, first_row, last_row, 
			seat_columns, created_at, updated_at
		FROM cabins
		ORDER BY aircraft_id, deck, first_row
	`

	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var cabins []*models.Cabin
	for rows.Next() {
		var cabin models.Cabin
		err := rows.Scan(
			&cabin.ID,
			&cabin.AircraftID,
			&cabin.SegmentID,
			&cabin.Deck,
			&cabin.FirstRow,
			&cabin.LastRow,
			&cabin.SeatColumns,
			&cabin.CreatedAt,
			&cabin.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		cabins = append(cabins, &cabin)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return cabins, nil
}

// Update updates a cabin in the database
func (r *CabinRepository) Update(cabin *models.Cabin) error {
	query := `
		UPDATE cabins
		SET aircraft_id = $2, segment_id = $3, deck = $4, first_row = $5, 
			last_row = $6, seat_columns = $7, updated_at = $8
		WHERE id = $1
	`

	_, err := r.db.Exec(
		query,
		cabin.ID,
		cabin.AircraftID,
		cabin.SegmentID,
		cabin.Deck,
		cabin.FirstRow,
		cabin.LastRow,
		cabin.SeatColumns,
		cabin.UpdatedAt,
	)

	return err
}

// Delete deletes a cabin by its ID
func (r *CabinRepository) Delete(id string) error {
	query := `DELETE FROM cabins WHERE id = $1`
	_, err := r.db.Exec(query, id)
	return err
}
