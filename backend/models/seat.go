package models

import (
	"time"
)

// Seat represents a seat in an aircraft cabin
type Seat struct {
	ID                  string    `json:"id"`
	RowID               string    `json:"row_id"`
	SegmentID           string    `json:"segment_id"`
	StorefrontSlotCode  string    `json:"storefront_slot_code"`
	Code                string    `json:"code"`
	Available           bool      `json:"available"`
	Entitled            bool      `json:"entitled"`
	FeeWaived           bool      `json:"fee_waived"`
	FreeOfCharge        bool      `json:"free_of_charge"`
	OriginallySelected  bool      `json:"originally_selected"`
	EntitledRuleID      string    `json:"entitled_rule_id"`
	FeeWaivedRuleID     string    `json:"fee_waived_rule_id"`
	RefundIndicator     string    `json:"refund_indicator"`
	SeatCharacteristics string    `json:"seat_characteristics"`
	RawCharacteristics  string    `json:"raw_characteristics"`
	CreatedAt           time.Time `json:"created_at"`
	UpdatedAt           time.Time `json:"updated_at"`
}

// SeatPrice represents the price of a seat
type SeatPrice struct {
	ID       string  `json:"id"`
	SeatID   string  `json:"seat_id"`
	Type     string  `json:"type"`
	Amount   float64 `json:"amount"`
	Currency string  `json:"currency"`
}

// SeatWithPrice represents a seat with its price
type SeatWithPrice struct {
	Seat  *Seat      `json:"seat"`
	Price *SeatPrice `json:"price"`
}