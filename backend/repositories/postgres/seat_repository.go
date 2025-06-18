package postgres

import (
	"database/sql"
	"errors"

	"github.com/evaizee/seat-arrangements/backend/models"
	"github.com/evaizee/seat-arrangements/backend/repositories"
)

// SeatRepository is a PostgreSQL implementation of the SeatRepository interface
type SeatRepository struct {
	db *sql.DB
}

// NewSeatRepository creates a new SeatRepository
func NewSeatRepository(db *sql.DB) repositories.SeatRepository {
	return &SeatRepository{db: db}
}

// Create creates a new seat in the database
func (r *SeatRepository) Create(seat *models.Seat) error {
	query := `
		INSERT INTO seats (
			id, row_id, segment_id, storefront_slot_code, code, available, 
			entitled, fee_waived, free_of_charge, originally_selected, 
			entitled_rule_id, fee_waived_rule_id, refund_indicator, 
			seat_characteristics, raw_characteristics, created_at, updated_at
		)
		VALUES (
			$1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17
		)
	`

	_, err := r.db.Exec(
		query,
		seat.ID,
		seat.RowID,
		seat.SegmentID,
		seat.StorefrontSlotCode,
		seat.Code,
		seat.Available,
		seat.Entitled,
		seat.FeeWaived,
		seat.FreeOfCharge,
		seat.OriginallySelected,
		seat.EntitledRuleID,
		seat.FeeWaivedRuleID,
		seat.RefundIndicator,
		seat.SeatCharacteristics,
		seat.RawCharacteristics,
		seat.CreatedAt,
		seat.UpdatedAt,
	)

	return err
}

// CreatePrice creates a new seat price in the database
func (r *SeatRepository) CreatePrice(price *models.SeatPrice) error {
	query := `
		INSERT INTO seat_prices (id, seat_id, type, amount, currency)
		VALUES ($1, $2, $3, $4, $5)
	`

	_, err := r.db.Exec(
		query,
		price.ID,
		price.SeatID,
		price.Type,
		price.Amount,
		price.Currency,
	)

	return err
}

// GetByID retrieves a seat by ID
func (r *SeatRepository) GetByID(id string) (*models.Seat, error) {
	query := `
		SELECT 
			id, row_id, segment_id, storefront_slot_code, code, available, 
			entitled, fee_waived, free_of_charge, originally_selected, 
			entitled_rule_id, fee_waived_rule_id, refund_indicator, 
			seat_characteristics, raw_characteristics, created_at, updated_at
		FROM seats
		WHERE id = $1
	`

	seat := &models.Seat{}
	err := r.db.QueryRow(query, id).Scan(
		&seat.ID,
		&seat.RowID,
		&seat.SegmentID,
		&seat.StorefrontSlotCode,
		&seat.Code,
		&seat.Available,
		&seat.Entitled,
		&seat.FeeWaived,
		&seat.FreeOfCharge,
		&seat.OriginallySelected,
		&seat.EntitledRuleID,
		&seat.FeeWaivedRuleID,
		&seat.RefundIndicator,
		&seat.SeatCharacteristics,
		&seat.RawCharacteristics,
		&seat.CreatedAt,
		&seat.UpdatedAt,
	)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil // Seat not found
		}
		return nil, err
	}

	return seat, nil
}

// GetByRowID retrieves seats by row ID
func (r *SeatRepository) GetByRowID(rowID string) ([]*models.Seat, error) {
	query := `
		SELECT 
			id, row_id, segment_id, storefront_slot_code, code, available, 
			entitled, fee_waived, free_of_charge, originally_selected, 
			entitled_rule_id, fee_waived_rule_id, refund_indicator, 
			seat_characteristics, raw_characteristics, created_at, updated_at
		FROM seats
		WHERE row_id = $1
	`

	rows, err := r.db.Query(query, rowID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	seats := []*models.Seat{}
	for rows.Next() {
		seat := &models.Seat{}
		err := rows.Scan(
			&seat.ID,
			&seat.RowID,
			&seat.SegmentID,
			&seat.StorefrontSlotCode,
			&seat.Code,
			&seat.Available,
			&seat.Entitled,
			&seat.FeeWaived,
			&seat.FreeOfCharge,
			&seat.OriginallySelected,
			&seat.EntitledRuleID,
			&seat.FeeWaivedRuleID,
			&seat.RefundIndicator,
			&seat.SeatCharacteristics,
			&seat.RawCharacteristics,
			&seat.CreatedAt,
			&seat.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		seats = append(seats, seat)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return seats, nil
}

// GetBySegmentID retrieves seats by segment ID
func (r *SeatRepository) GetBySegmentID(segmentID string) ([]*models.Seat, error) {
	query := `
		SELECT 
			id, row_id, segment_id, storefront_slot_code, code, available, 
			entitled, fee_waived, free_of_charge, originally_selected, 
			entitled_rule_id, fee_waived_rule_id, refund_indicator, 
			seat_characteristics, raw_characteristics, created_at, updated_at
		FROM seats
		WHERE segment_id = $1
	`

	rows, err := r.db.Query(query, segmentID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	seats := []*models.Seat{}
	for rows.Next() {
		seat := &models.Seat{}
		err := rows.Scan(
			&seat.ID,
			&seat.RowID,
			&seat.SegmentID,
			&seat.StorefrontSlotCode,
			&seat.Code,
			&seat.Available,
			&seat.Entitled,
			&seat.FeeWaived,
			&seat.FreeOfCharge,
			&seat.OriginallySelected,
			&seat.EntitledRuleID,
			&seat.FeeWaivedRuleID,
			&seat.RefundIndicator,
			&seat.SeatCharacteristics,
			&seat.RawCharacteristics,
			&seat.CreatedAt,
			&seat.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		seats = append(seats, seat)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return seats, nil
}

// GetPriceBySeatID retrieves a seat price by seat ID
func (r *SeatRepository) GetPriceBySeatID(seatID string) (*models.SeatPrice, error) {
	query := `
		SELECT id, seat_id, type, amount, currency
		FROM seat_prices
		WHERE seat_id = $1
	`

	price := &models.SeatPrice{}
	err := r.db.QueryRow(query, seatID).Scan(
		&price.ID,
		&price.SeatID,
		&price.Type,
		&price.Amount,
		&price.Currency,
	)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil // Price not found
		}
		return nil, err
	}

	return price, nil
}

// GetAll retrieves all seats
func (r *SeatRepository) GetAll() ([]*models.Seat, error) {
	query := `
		SELECT 
			id, row_id, segment_id, storefront_slot_code, code, available, 
			entitled, fee_waived, free_of_charge, originally_selected, 
			entitled_rule_id, fee_waived_rule_id, refund_indicator, 
			seat_characteristics, raw_characteristics, created_at, updated_at
		FROM seats
	`

	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	seats := []*models.Seat{}
	for rows.Next() {
		seat := &models.Seat{}
		err := rows.Scan(
			&seat.ID,
			&seat.RowID,
			&seat.SegmentID,
			&seat.StorefrontSlotCode,
			&seat.Code,
			&seat.Available,
			&seat.Entitled,
			&seat.FeeWaived,
			&seat.FreeOfCharge,
			&seat.OriginallySelected,
			&seat.EntitledRuleID,
			&seat.FeeWaivedRuleID,
			&seat.RefundIndicator,
			&seat.SeatCharacteristics,
			&seat.RawCharacteristics,
			&seat.CreatedAt,
			&seat.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		seats = append(seats, seat)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return seats, nil
}

// Update updates a seat in the database
func (r *SeatRepository) Update(seat *models.Seat) error {
	query := `
		UPDATE seats
		SET 
			row_id = $2,
			segment_id = $3,
			storefront_slot_code = $4,
			code = $5,
			available = $6,
			entitled = $7,
			fee_waived = $8,
			free_of_charge = $9,
			originally_selected = $10,
			entitled_rule_id = $11,
			fee_waived_rule_id = $12,
			refund_indicator = $13,
			seat_characteristics = $14,
			raw_characteristics = $15,
			updated_at = $16
		WHERE id = $1
	`

	_, err := r.db.Exec(
		query,
		seat.ID,
		seat.RowID,
		seat.SegmentID,
		seat.StorefrontSlotCode,
		seat.Code,
		seat.Available,
		seat.Entitled,
		seat.FeeWaived,
		seat.FreeOfCharge,
		seat.OriginallySelected,
		seat.EntitledRuleID,
		seat.FeeWaivedRuleID,
		seat.RefundIndicator,
		seat.SeatCharacteristics,
		seat.RawCharacteristics,
		seat.UpdatedAt,
	)

	return err
}

// UpdatePrice updates a seat price in the database
func (r *SeatRepository) UpdatePrice(price *models.SeatPrice) error {
	query := `
		UPDATE seat_prices
		SET seat_id = $2, type = $3, amount = $4, currency = $5
		WHERE id = $1
	`

	_, err := r.db.Exec(
		query,
		price.ID,
		price.SeatID,
		price.Type,
		price.Amount,
		price.Currency,
	)

	return err
}

// Delete deletes a seat from the database
func (r *SeatRepository) Delete(id string) error {
	query := `DELETE FROM seats WHERE id = $1`
	_, err := r.db.Exec(query, id)
	return err
}

// DeletePrice deletes a seat price from the database
func (r *SeatRepository) DeletePrice(id string) error {
	query := `DELETE FROM seat_prices WHERE id = $1`
	_, err := r.db.Exec(query, id)
	return err
}

// GetByFlightID retrieves seats by flight ID (segment ID)
func (r *SeatRepository) GetByFlightID(flightID string) ([]*models.Seat, error) {
	query := `
		SELECT 
			s.id, s.row_id, s.segment_id, s.storefront_slot_code, COALESCE(s.code, ''), s.available, 
			s.entitled, s.fee_waived, s.free_of_charge, s.originally_selected, 
			COALESCE(s.entitled_rule_id, ''), COALESCE(s.fee_waived_rule_id, ''), COALESCE(s.refund_indicator, ''), 
			COALESCE(s.seat_characteristics, ''), COALESCE(s.raw_characteristics, ''), s.created_at, s.updated_at
		FROM seats s
		WHERE s.segment_id = $1
	`

	rows, err := r.db.Query(query, flightID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	seats := []*models.Seat{}
	for rows.Next() {
		seat := &models.Seat{}
		err := rows.Scan(
			&seat.ID,
			&seat.RowID,
			&seat.SegmentID,
			&seat.StorefrontSlotCode,
			&seat.Code,
			&seat.Available,
			&seat.Entitled,
			&seat.FeeWaived,
			&seat.FreeOfCharge,
			&seat.OriginallySelected,
			&seat.EntitledRuleID,
			&seat.FeeWaivedRuleID,
			&seat.RefundIndicator,
			&seat.SeatCharacteristics,
			&seat.RawCharacteristics,
			&seat.CreatedAt,
			&seat.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		seats = append(seats, seat)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return seats, nil
}