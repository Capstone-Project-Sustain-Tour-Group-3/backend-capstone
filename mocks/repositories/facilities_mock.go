package repositories

import (
	"capstone/entities"

	"github.com/google/uuid"
	"github.com/stretchr/testify/mock"
)

type MockFacilityRepository struct {
	mock.Mock
}

func (m *MockFacilityRepository) Create(facility *entities.Facility) error {
	args := m.Called(facility)
	if args.Error(0) == nil {
		return nil
	}
	return args.Error(0)
}

func (m *MockFacilityRepository) FindAll() ([]entities.Facility, error) {
	args := m.Called()
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]entities.Facility), nil
}

func (m *MockFacilityRepository) FindById(id uuid.UUID) (*entities.Facility, error) {
	args := m.Called(id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*entities.Facility), nil
}

func (m *MockFacilityRepository) Update(facility *entities.Facility) error {
	args := m.Called(facility)
	if args.Error(0) == nil {
		return nil
	}
	return args.Error(0)
}

func (m *MockFacilityRepository) Delete(facility *entities.Facility) error {
	args := m.Called(facility)
	if args.Error(0) == nil {
		return nil
	}
	return args.Error(0)
}
