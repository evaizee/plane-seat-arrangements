package impl

import (
	"fmt"
	"strconv"

	"github.com/evaizee/seat-arrangements/backend/models"
	"github.com/evaizee/seat-arrangements/backend/repositories"
	"github.com/evaizee/seat-arrangements/backend/services"
)

// PassengerService is an implementation of the PassengerService interface
type PassengerService struct {
	passengerRepository     repositories.PassengerRepository
	frequentFlyerRepository repositories.FrequentFlyerRepository
}

// NewPassengerService creates a new passenger service
func NewPassengerService(
	passengerRepository repositories.PassengerRepository,
	frequentFlyerRepository repositories.FrequentFlyerRepository,
) services.PassengerService {
	return &PassengerService{
		passengerRepository:     passengerRepository,
		frequentFlyerRepository: frequentFlyerRepository,
	}
}

// CreatePassenger creates a new passenger
func (s *PassengerService) CreatePassenger(segmentID int, passengerIndex int, passengerNameNumber, firstName, lastName string) (*models.Passenger, error) {
	passenger := models.NewPassenger(segmentID, passengerIndex, passengerNameNumber, firstName, lastName)

	err := s.passengerRepository.Create(passenger)
	if err != nil {
		return nil, fmt.Errorf("error creating passenger: %w", err)
	}

	return passenger, nil
}

// GetByID retrieves a passenger by ID
func (s *PassengerService) GetByID(id string) (*models.Passenger, error) {
	return s.passengerRepository.GetByID(id)
}

// GetBySegmentID retrieves passengers by segment ID
func (s *PassengerService) GetBySegmentID(segmentID string) ([]*models.Passenger, error) {
	return s.passengerRepository.GetBySegmentID(segmentID)
}

// GetWithFrequentFlyers retrieves a passenger with their frequent flyer details
func (s *PassengerService) GetWithFrequentFlyers(id string) (*models.PassengerWithDetails, error) {
	passenger, err := s.passengerRepository.GetByID(id)
	if err != nil {
		return nil, err
	}

	frequentFlyers, err := s.frequentFlyerRepository.GetByPassengerID(strconv.Itoa(passenger.ID))
	if err != nil {
		return nil, fmt.Errorf("error getting frequent flyers: %w", err)
	}

	return &models.PassengerWithDetails{
		Passenger:      *passenger,
		FrequentFlyers: frequentFlyers,
	}, nil
}

// UpdatePassenger updates a passenger
func (s *PassengerService) UpdatePassenger(passenger *models.Passenger) error {
	return s.passengerRepository.Update(passenger)
}