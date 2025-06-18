package impl

import (
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/evaizee/seat-arrangements/backend/models"
	"github.com/evaizee/seat-arrangements/backend/repositories"
	"github.com/evaizee/seat-arrangements/backend/services"
	"go.uber.org/zap"
)

// SeatService is an implementation of the SeatService interface
type SeatService struct {
	seatRepository      repositories.SeatRepository
	flightRepository    repositories.FlightRepository
	aircraftRepository  repositories.AircraftRepository
	cabinRepository     repositories.CabinRepository
	passengerRepository repositories.PassengerRepository
	rowRepository       repositories.RowRepository
}

// NewSeatService creates a new SeatService
// Additional repositories are optional for backward compatibility
func NewSeatService(
	seatRepository repositories.SeatRepository,
	optionalRepos ...interface{},
) services.SeatService {
	service := &SeatService{
		seatRepository: seatRepository,
	}

	// Process optional repositories if provided
	for _, repo := range optionalRepos {
		switch r := repo.(type) {
		case repositories.FlightRepository:
			service.flightRepository = r
		case repositories.AircraftRepository:
			service.aircraftRepository = r
		case repositories.CabinRepository:
			service.cabinRepository = r
		case repositories.PassengerRepository:
			service.passengerRepository = r
		case repositories.RowRepository:
			service.rowRepository = r
		}
	}

	return service
}

// GetByID retrieves a seat with its price by ID
func (s *SeatService) GetByID(id string) (*models.SeatWithPrice, error) {
	// Get the seat
	seat, err := s.seatRepository.GetByID(id)
	if err != nil {
		zap.L().Error("Failed to get seat", zap.Error(err), zap.String("seat_id", id))
		return nil, err
	}

	if seat == nil {
		return nil, errors.New("seat not found")
	}

	// Get the seat price
	price, err := s.seatRepository.GetPriceBySeatID(id)
	if err != nil {
		zap.L().Error("Failed to get seat price", zap.Error(err), zap.String("seat_id", id))
		return nil, err
	}

	// Return the seat with its price
	return &models.SeatWithPrice{
		Seat:  seat,
		Price: price,
	}, nil
}

// GetByCabinID retrieves seats with their prices by cabin ID
func (s *SeatService) GetByRowID(rowID string) ([]*models.SeatWithPrice, error) {
	// Get the seats
	seats, err := s.seatRepository.GetByRowID(rowID)
	if err != nil {
		zap.L().Error("Failed to get seats by row ID", zap.Error(err), zap.String("row_id", rowID))
		return nil, err
	}

	// Get the prices for each seat
	result := make([]*models.SeatWithPrice, 0, len(seats))
	for _, seat := range seats {
		price, err := s.seatRepository.GetPriceBySeatID(seat.ID)
		if err != nil {
			zap.L().Error("Failed to get seat price", zap.Error(err), zap.String("seat_id", seat.ID))
			continue
		}

		result = append(result, &models.SeatWithPrice{
			Seat:  seat,
			Price: price,
		})
	}

	return result, nil
}

// GetByFlightID retrieves seats with their prices by flight ID
func (s *SeatService) GetByFlightID(flightID string) ([]*models.SeatWithPrice, error) {
	seats, err := s.seatRepository.GetByFlightID(flightID)
	if err != nil {
		zap.L().Error("Failed to get seats by flight ID", zap.Error(err), zap.String("flight_id", flightID))
		return nil, err
	}

	// Get the prices for each seat
	result := make([]*models.SeatWithPrice, 0, len(seats))
	for _, seat := range seats {
		price, err := s.seatRepository.GetPriceBySeatID(seat.ID)
		if err != nil {
			zap.L().Error("Failed to get seat price", zap.Error(err), zap.String("seat_id", seat.ID))
			continue
		}

		result = append(result, &models.SeatWithPrice{
			Seat:  seat,
			Price: price,
		})
	}

	return result, nil
}

// UpdateAvailability updates the availability of a seat
func (s *SeatService) UpdateAvailability(seatID string, available bool) error {
	// Get the seat
	seat, err := s.seatRepository.GetByID(seatID)
	if err != nil {
		zap.L().Error("Failed to get seat", zap.Error(err), zap.String("seat_id", seatID))
		return err
	}

	if seat == nil {
		return errors.New("seat not found")
	}

	// Update the seat
	seat.Available = available
	seat.UpdatedAt = time.Now()

	// Save the seat
	err = s.seatRepository.Update(seat)
	if err != nil {
		zap.L().Error("Failed to update seat", zap.Error(err), zap.String("seat_id", seatID))
		return err
	}

	return nil
}

// getMockSeatMap provides sample seat map data for development
func (s *SeatService) getMockSeatMap(flightID string, passengerID string) (*models.SeatMapResponse, error) {
	// Get passenger info if available
	var passenger *models.Passenger
	var err error
	if s.passengerRepository != nil && passengerID != "" {
		passenger, err = s.passengerRepository.GetByID(passengerID)
		if err != nil {
			zap.L().Warn("Failed to get passenger, using default",
				zap.Error(err),
				zap.String("passenger_id", passengerID))
		}
	}

	// Create a mock seat map response
	response := &models.SeatMapResponse{
		SeatsItineraryParts: []models.ItineraryPart{
			{
				SegmentSeatMaps: []models.SegmentSeatMap{
					{
						PassengerSeatMaps: []models.PassengerSeatMap{
							{
								SeatSelectionEnabledForPax: true,
								SeatMap: models.SeatMap{
									RowsDisabledCauses: []string{},
									Aircraft:           "B737",
									Cabins: []models.CabinMap{
										{
											Deck:        "MAIN",
											SeatColumns: []string{"A", "B", "C", "D", "E", "F"},
											FirstRow:    1,
											LastRow:     30,
											SeatRows:    generateMockSeatRows(1, 30),
										},
									},
								},
							},
						},
						Segment: models.Segment{
							Type: "FlightSegment",
							SegmentOfferInformation: models.SegmentOfferInformation{
								FlightsMiles: 500,
							},
						},
					},
				},
			},
		},
		SelectedSeats: []string{},
	}

	// Add passenger info if available
	if passenger != nil {
		// Convert from Passenger model to PassengerInfo for the response
		var gender, email, phone string
		if passenger.Gender != nil {
			gender = *passenger.Gender
		}
		if passenger.Email != nil {
			email = *passenger.Email
		}
		if passenger.Phone != nil {
			phone = *passenger.Phone
		}

		response.SeatsItineraryParts[0].SegmentSeatMaps[0].PassengerSeatMaps[0].Passenger = models.PassengerInfo{
			PassengerIndex:      passenger.PassengerIndex,
			PassengerNameNumber: passenger.PassengerNameNumber,
			PassengerDetails: models.PassengerDetails{
				FirstName: passenger.FirstName,
				LastName:  passenger.LastName,
				Gender:    gender,
				Email:     email,
				Phone:     phone,
			},
		}
	} else {
		// Add default passenger info
		response.SeatsItineraryParts[0].SegmentSeatMaps[0].PassengerSeatMaps[0].Passenger = models.PassengerInfo{
			PassengerIndex:      1,
			PassengerNameNumber: "01.01",
			PassengerDetails: models.PassengerDetails{
				FirstName: "John",
				LastName:  "Doe",
			},
		}
	}

	return response, nil
}

// Helper function to generate mock seat rows
func generateMockSeatRows(firstRow, lastRow int) []models.SeatMapRow {
	rows := make([]models.SeatMapRow, 0, lastRow-firstRow+1)
	columns := []string{"A", "B", "C", "D", "E", "F"}

	for i := firstRow; i <= lastRow; i++ {
		seatCodes := make([]string, 0, len(columns))
		seats := make([]models.SeatMapItem, 0, len(columns))

		for _, col := range columns {
			seatCode := fmt.Sprintf("%d%s", i, col)
			seatCodes = append(seatCodes, seatCode)

			// Create seat with different statuses for variety
			seatStatus := "AVAILABLE"
			seatType := "STANDARD"

			// Make some seats unavailable or premium
			if i < 5 {
				seatType = "PREMIUM"
			}
			if i == 10 || i == 20 {
				seatStatus = "OCCUPIED"
			}
			if i == 15 && (col == "A" || col == "F") {
				seatStatus = "BLOCKED"
			}

			// Create seat with correct field names
			available := seatStatus == "AVAILABLE"
			designations := []string{}
			if seatType == "PREMIUM" {
				designations = append(designations, "PREMIUM")
			}

			seats = append(seats, models.SeatMapItem{
				StorefrontSlotCode:  seatCode,
				Code:                seatCode,
				Available:           available,
				Entitled:            true,
				FeeWaived:           false,
				FreeOfCharge:        seatType == "STANDARD",
				OriginallySelected:  false,
				Designations:        designations,
				SlotCharacteristics: []string{},
			})
		}

		rows = append(rows, models.SeatMapRow{
			RowNumber: i,
			SeatCodes: seatCodes,
			Seats:     seats,
		})
	}

	return rows
}

// GetSeatMap generates a seat map for a flight and passenger
func (s *SeatService) GetSeatMap(flightID string, passengerID string) (*models.SeatMapResponse, error) {
	// Check if we have all required repositories
	if s.flightRepository == nil || s.aircraftRepository == nil || s.cabinRepository == nil {
		// If repositories are missing, return mock data for development
		zap.L().Warn("Using mock data for seat map as some repositories are not initialized")
		return s.getMockSeatMap(flightID, passengerID)
	} else {
		zap.L().Warn("Using real data for seat map, cool")
	}

	// Get flight details
	flight, err := s.flightRepository.GetByID(flightID)
	if err != nil {
		zap.L().Error("Failed to get flight", zap.Error(err), zap.String("flight_id", flightID))
		return nil, err
	}

	if flight == nil {
		return nil, errors.New("flight not found")
	}

	// Get aircraft details using the equipment field from flight
	aircraft, err := s.aircraftRepository.GetByCode(flight.Equipment)
	if err != nil {
		zap.L().Error("Failed to get aircraft", zap.Error(err), zap.String("aircraft_code", flight.Equipment))
		return nil, err
	}

	if aircraft == nil {
		return nil, errors.New("aircraft not found for flight")
	}

	// Get cabins for this flight
	cabins, err := s.cabinRepository.GetBySegmentID(flightID)
	if err != nil {
		zap.L().Error("Failed to get cabins", zap.Error(err), zap.String("flight_id", flightID))
		return nil, err
	}

	// Get seat rows for each cabin
	if s.rowRepository == nil {
		zap.L().Warn("Row repository not initialized, skipping row retrieval")
	}

	// Get all seats for this flight
	seats, err := s.GetByFlightID(flightID)
	if err != nil {
		zap.L().Error("Failed to get seats", zap.Error(err), zap.String("flight_id", flightID))
		return nil, err
	}

	// Get passenger details if provided
	var passenger *models.Passenger
	if passengerID != "" {
		passenger, err = s.passengerRepository.GetByID(passengerID)
		if err != nil {
			zap.L().Error("Failed to get passenger", zap.Error(err), zap.String("passenger_id", passengerID))
			// Continue without passenger info
		}
	}

	// Build the seat map response
	seatMapResponse := &models.SeatMapResponse{
		SeatsItineraryParts: []models.ItineraryPart{
			{
				SegmentSeatMaps: []models.SegmentSeatMap{
					{
						PassengerSeatMaps: []models.PassengerSeatMap{},
						Segment: models.Segment{
							Type: "Segment",
							SegmentOfferInformation: models.SegmentOfferInformation{
								FlightsMiles: 0, // Set this based on your data
							},
						},
					},
				},
			},
		},
		SelectedSeats: []string{},
	}

	// Create a passenger seat map
	passengerSeatMap := models.PassengerSeatMap{
		SeatSelectionEnabledForPax: true,
		SeatMap: models.SeatMap{
			RowsDisabledCauses: []string{},
			Aircraft:           aircraft.Code,
			Cabins:             []models.CabinMap{},
		},
	}

	// Add passenger info if available
	if passenger != nil {
		// Convert from Passenger model to PassengerInfo for the response
		var gender, email, phone string
		if passenger.Gender != nil {
			gender = *passenger.Gender
		}
		if passenger.Email != nil {
			email = *passenger.Email
		}
		if passenger.Phone != nil {
			phone = *passenger.Phone
		}

		passengerSeatMap.Passenger = models.PassengerInfo{
			PassengerIndex:      passenger.PassengerIndex,
			PassengerNameNumber: passenger.PassengerNameNumber,
			PassengerDetails: models.PassengerDetails{
				FirstName: passenger.FirstName,
				LastName:  passenger.LastName,
				Gender:    gender,
				Email:     email,
				Phone:     phone,
			},
		}
	} else {
		// Add default passenger info
		passengerSeatMap.Passenger = models.PassengerInfo{
			PassengerIndex:      1,
			PassengerNameNumber: "01.01",
			PassengerDetails:    models.PassengerDetails{},
		}
	}

	// Process each cabin
	for _, cabin := range cabins {
		// Parse seat columns from the stored string
		columns := strings.Split(cabin.SeatColumns, ",")

		// Create cabin map
		cabinMap := models.CabinMap{
			Deck:        cabin.Deck,
			SeatColumns: append([]string{"LEFT_SIDE"}, append(columns, "RIGHT_SIDE")...),
			SeatRows:    []models.SeatMapRow{},
			FirstRow:    cabin.FirstRow,
			LastRow:     cabin.LastRow,
		}

		// Get rows for this cabin
		var cabinRows []*models.SeatRow
		if s.rowRepository != nil {
			cabinRows, err = s.rowRepository.GetByCabinID(cabin.ID)
			if err != nil {
				zap.L().Error("Failed to get rows for cabin", zap.Error(err), zap.String("cabin_id", cabin.ID))
				// Continue with empty rows
				cabinRows = []*models.SeatRow{}
			}
		}

		// Group seats by row
		seatsByRow := make(map[int][]models.SeatWithPrice)
		for _, seat := range seats {
			// Check if this seat belongs to a row in this cabin
			belongsToCabin := false
			for _, row := range cabinRows {
				if seat.Seat.RowID == row.ID {
					belongsToCabin = true
					break
				}
			}

			if belongsToCabin || len(cabinRows) == 0 { // If no rows found, include all seats as fallback
				// Extract row number from seat code (e.g., "4A" -> 4)
				rowNum := 0
				fmt.Sscanf(seat.Seat.Code, "%d", &rowNum)

				if _, ok := seatsByRow[rowNum]; !ok {
					seatsByRow[rowNum] = []models.SeatWithPrice{}
				}
				seatsByRow[rowNum] = append(seatsByRow[rowNum], *seat)
			}
		}

		// Create seat rows
		for rowNum := cabin.FirstRow; rowNum <= cabin.LastRow; rowNum++ {
			seatRow := models.SeatMapRow{
				RowNumber: rowNum,
				SeatCodes: []string{},
				Seats:     []models.SeatMapItem{},
			}

			// Add left side placeholder
			seatRow.Seats = append(seatRow.Seats, models.SeatMapItem{
				SlotCharacteristics: []string{"LEFT_SIDE"},
				StorefrontSlotCode:  "BLANK",
				Available:           false,
				Entitled:            false,
				FeeWaived:           false,
				FreeOfCharge:        true,
				OriginallySelected:  false,
			})

			// Add seats for this row
			rowSeats, hasSeats := seatsByRow[rowNum]
			if hasSeats {
				// Add actual seats
				for _, col := range columns {
					seatCode := fmt.Sprintf("%d%s", rowNum, col)
					found := false

					// Find the seat with this code
					for _, s := range rowSeats {
						if s.Seat.Code == seatCode {
							// Parse seat characteristics
							characteristics := []string{}
							if s.Seat.SeatCharacteristics != "" {
								characteristics = strings.Split(s.Seat.SeatCharacteristics, ",")
							}

							// Add the seat
							seatRow.Seats = append(seatRow.Seats, models.SeatMapItem{
								StorefrontSlotCode:  s.Seat.StorefrontSlotCode,
								Available:           s.Seat.Available,
								Code:                s.Seat.Code,
								Entitled:            s.Seat.Entitled,
								FeeWaived:           s.Seat.FeeWaived,
								FreeOfCharge:        s.Seat.FreeOfCharge,
								OriginallySelected:  s.Seat.OriginallySelected,
								SlotCharacteristics: characteristics,
								Designations:        []string{},
							})
							found = true
							break
						}
					}

					// If seat not found, add a blank seat
					if !found {
						seatRow.Seats = append(seatRow.Seats, models.SeatMapItem{
							StorefrontSlotCode: "BLANK",
							Available:          false,
							Entitled:           false,
							FeeWaived:          false,
							FreeOfCharge:       true,
							OriginallySelected: false,
						})
					}
				}
			} else {
				// No seats for this row, add blanks
				for range columns {
					seatRow.Seats = append(seatRow.Seats, models.SeatMapItem{
						StorefrontSlotCode: "BLANK",
						Available:          false,
						Entitled:           false,
						FeeWaived:          false,
						FreeOfCharge:       true,
						OriginallySelected: false,
					})
				}
			}

			// Add right side placeholder
			seatRow.Seats = append(seatRow.Seats, models.SeatMapItem{
				SlotCharacteristics: []string{"RIGHT_SIDE"},
				StorefrontSlotCode:  "BLANK",
				Available:           false,
				Entitled:            false,
				FeeWaived:           false,
				FreeOfCharge:        true,
				OriginallySelected:  false,
			})

			// Add the row to the cabin
			cabinMap.SeatRows = append(cabinMap.SeatRows, seatRow)
		}

		// Add the cabin to the seat map
		passengerSeatMap.SeatMap.Cabins = append(passengerSeatMap.SeatMap.Cabins, cabinMap)
	}

	// Add the passenger seat map to the response
	seatMapResponse.SeatsItineraryParts[0].SegmentSeatMaps[0].PassengerSeatMaps = append(
		seatMapResponse.SeatsItineraryParts[0].SegmentSeatMaps[0].PassengerSeatMaps,
		passengerSeatMap,
	)

	return seatMapResponse, nil
}
