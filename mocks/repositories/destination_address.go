package repositories

import (
	"capstone/entities"
	"github.com/google/uuid"
	"github.com/stretchr/testify/mock"
	"gorm.io/gorm"
)

type MockDestinationAddressRepository struct {
	mock.Mock
}

func (m *MockDestinationAddressRepository) Create(destinationAddress *entities.DestinationAddress, tx *gorm.DB) error {
	args := m.Called(destinationAddress, tx)
	if args.Error(0) == nil {
		return nil
	}
	return args.Error(0)
}

func (m *MockDestinationAddressRepository) FindAll() ([]entities.DestinationAddress, error) {
	args := m.Called()
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]entities.DestinationAddress), nil
}

func (m *MockDestinationAddressRepository) FindById(id uuid.UUID) (*entities.DestinationAddress, error) {
	args := m.Called(id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*entities.DestinationAddress), nil
}

func (m *MockDestinationAddressRepository) Update(destinationAddress *entities.DestinationAddress) error {
	args := m.Called(destinationAddress)
	if args.Error(0) == nil {
		return nil
	}
	return args.Error(0)
}

func (m *MockDestinationAddressRepository) Delete(destinationAddress *entities.DestinationAddress) error {
	args := m.Called(destinationAddress)
	if args.Error(0) == nil {
		return nil
	}
	return args.Error(0)
}
