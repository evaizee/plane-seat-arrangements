package models

// SeatMapResponse represents the complete seat map response
type SeatMapResponse struct {
	SeatsItineraryParts []ItineraryPart `json:"seatsItineraryParts"`
	SelectedSeats       []string        `json:"selectedSeats"`
}

// ItineraryPart represents a part of the itinerary
type ItineraryPart struct {
	SegmentSeatMaps []SegmentSeatMap `json:"segmentSeatMaps"`
}

// SegmentSeatMap represents a seat map for a segment
type SegmentSeatMap struct {
	PassengerSeatMaps []PassengerSeatMap `json:"passengerSeatMaps"`
	Segment           Segment            `json:"segment"`
}

// Segment represents a flight segment
type Segment struct {
	Type                    string                  `json:"@type"`
	SegmentOfferInformation SegmentOfferInformation `json:"segmentOfferInformation"`
}

// SegmentOfferInformation contains information about the segment offer
type SegmentOfferInformation struct {
	FlightsMiles int `json:"flightsMiles"`
	// Add other fields as needed
}

// PassengerSeatMap represents a seat map for a passenger
type PassengerSeatMap struct {
	SeatSelectionEnabledForPax bool          `json:"seatSelectionEnabledForPax"`
	SeatMap                    SeatMap       `json:"seatMap"`
	Passenger                  PassengerInfo `json:"passenger"`
}

// SeatMap represents the actual seat map
type SeatMap struct {
	RowsDisabledCauses []string   `json:"rowsDisabledCauses"`
	Aircraft           string     `json:"aircraft"`
	Cabins             []CabinMap `json:"cabins"`
}

// CabinMap represents a cabin in the seat map
type CabinMap struct {
	Deck        string       `json:"deck"`
	SeatColumns []string     `json:"seatColumns"`
	SeatRows    []SeatMapRow `json:"seatRows"`
	FirstRow    int          `json:"firstRow"`
	LastRow     int          `json:"lastRow"`
}

// SeatMapRow represents a row of seats in the seat map response
type SeatMapRow struct {
	RowNumber int           `json:"rowNumber"`
	SeatCodes []string      `json:"seatCodes"`
	Seats     []SeatMapItem `json:"seats"`
}

// SeatMapItem represents a seat in the seat map
type SeatMapItem struct {
	SlotCharacteristics []string `json:"slotCharacteristics,omitempty"`
	StorefrontSlotCode  string   `json:"storefrontSlotCode"`
	Available           bool     `json:"available"`
	Code                string   `json:"code,omitempty"`
	Entitled            bool     `json:"entitled"`
	FeeWaived           bool     `json:"feeWaived"`
	FreeOfCharge        bool     `json:"freeOfCharge"`
	OriginallySelected  bool     `json:"originallySelected"`
	Designations        []string `json:"designations,omitempty"`
}

// PassengerInfo represents information about a passenger
type PassengerInfo struct {
	PassengerIndex      int              `json:"passengerIndex"`
	PassengerNameNumber string           `json:"passengerNameNumber"`
	PassengerDetails    PassengerDetails `json:"passengerDetails"`
}

// PassengerDetails contains details about a passenger
type PassengerDetails struct {
	FirstName string `json:"firstName,omitempty"`
	LastName  string `json:"lastName,omitempty"`
	Gender    string `json:"gender,omitempty"`
	Email     string `json:"email,omitempty"`
	Phone     string `json:"phone,omitempty"`
}
