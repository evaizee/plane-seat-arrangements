package impl

import (
	"fmt"

	"github.com/evaizee/seat-arrangements/backend/models"
	"github.com/evaizee/seat-arrangements/backend/repositories"
	"github.com/evaizee/seat-arrangements/backend/services"
)

// FrequentFlyerService is an implementation of the FrequentFlyerService interface
type FrequentFlyerService struct {
	frequentFlyerRepository repositories.FrequentFlyerRepository
}

// NewFrequentFlyerService creates a new frequent flyer service
func NewFrequentFlyerService(
	frequentFlyerRepository repositories.FrequentFlyerRepository,
) services.FrequentFlyerService {
	return &FrequentFlyerService{
		frequentFlyerRepository: frequentFlyerRepository,
	}
}

// CreateFrequentFlyer creates a new frequent flyer record
func (s *FrequentFlyerService) CreateFrequentFlyer(passengerID int, airline, number string) (*models.FrequentFlyer, error) {
	frequentFlyer := models.NewFrequentFlyer(passengerID, airline, number)

	err := s.frequentFlyerRepository.Create(frequentFlyer)
	if err != nil {
		return nil, fmt.Errorf("error creating frequent flyer: %w", err)
	}

	return frequentFlyer, nil
}

// GetByID retrieves a frequent flyer by ID
func (s *FrequentFlyerService) GetByID(id string) (*models.FrequentFlyer, error) {
	return s.frequentFlyerRepository.GetByID(id)
}

// GetByPassengerID retrieves frequent flyers by passenger ID
func (s *FrequentFlyerService) GetByPassengerID(passengerID string) ([]*models.FrequentFlyer, error) {
	return s.frequentFlyerRepository.GetByPassengerID(passengerID)
}

// UpdateFrequentFlyer updates a frequent flyer
func (s *FrequentFlyerService) UpdateFrequentFlyer(frequentFlyer *models.FrequentFlyer) error {
	return s.frequentFlyerRepository.Update(frequentFlyer)
}