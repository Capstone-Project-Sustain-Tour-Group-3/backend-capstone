package repositories

import (
	"capstone/entities"
	"github.com/google/uuid"
	"github.com/stretchr/testify/mock"
	"gorm.io/gorm"
)

type MockDestinationFacilityRepository struct {
	mock.Mock
}

func (m *MockDestinationFacilityRepository) Create(destinationFacilities *[]entities.DestinationFacility, tx *gorm.DB) error {
	args := m.Called(destinationFacilities, tx)
	if args.Error(0) == nil {
		return nil
	}
	return args.Error(0)
}

func (m *MockDestinationFacilityRepository) FindAll() ([]entities.DestinationFacility, error) {
	args := m.Called()
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]entities.DestinationFacility), nil
}

func (m *MockDestinationFacilityRepository) FindById(id uuid.UUID) (*entities.DestinationFacility, error) {
	args := m.Called(id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*entities.DestinationFacility), nil
}

func (m *MockDestinationFacilityRepository) Update(destinationFacility *entities.DestinationFacility) error {
	args := m.Called(destinationFacility)
	if args.Error(0) == nil {
		return nil
	}
	return args.Error(0)
}

func (m *MockDestinationFacilityRepository) Delete(destinationFacility *entities.DestinationFacility) error {
	args := m.Called(destinationFacility)
	if args.Error(0) == nil {
		return nil
	}
	return args.Error(0)
}

func (m *MockDestinationFacilityRepository) DeleteMany(destinationFacilities *[]entities.DestinationFacility) error {
	args := m.Called(destinationFacilities)
	if args.Error(0) == nil {
		return nil
	}
	return args.Error(0)
}
