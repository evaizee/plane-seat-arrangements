package postgres

import (
	"database/sql"

	"github.com/evaizee/seat-arrangements/backend/models"
	"github.com/evaizee/seat-arrangements/backend/repositories"
	"go.uber.org/zap"
)

// RowRepository is a PostgreSQL implementation of the RowRepository interface
type RowRepository struct {
	db *sql.DB
}

// NewRowRepository creates a new RowRepository
func NewRowRepository(db *sql.DB) repositories.RowRepository {
	return &RowRepository{
		db: db,
	}
}

// GetByCabinID retrieves seat rows by cabin ID
func (r *RowRepository) GetByCabinID(cabinID string) ([]*models.SeatRow, error) {
	query := `
		SELECT 
			id, cabin_id, row_number, seat_codes, created_at, updated_at
		FROM seat_rows
		WHERE cabin_id = $1
		ORDER BY row_number ASC
	`

	rows, err := r.db.Query(query, cabinID)
	if err != nil {
		zap.L().Error("Failed to get seat rows by cabin ID", zap.Error(err), zap.String("cabin_id", cabinID))
		return nil, err
	}
	defer rows.Close()

	seatRows := []*models.SeatRow{}
	for rows.Next() {
		seatRow := &models.SeatRow{}
		err := rows.Scan(
			&seatRow.ID,
			&seatRow.CabinID,
			&seatRow.RowNumber,
			&seatRow.SeatCodes,
			&seatRow.CreatedAt,
			&seatRow.UpdatedAt,
		)
		if err != nil {
			zap.L().Error("Failed to scan seat row", zap.Error(err))
			return nil, err
		}
		seatRows = append(seatRows, seatRow)
	}

	if err = rows.Err(); err != nil {
		zap.L().Error("Error iterating over seat rows", zap.Error(err))
		return nil, err
	}

	return seatRows, nil
}
