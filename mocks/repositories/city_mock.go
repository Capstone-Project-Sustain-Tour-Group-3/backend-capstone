package repositories

import (
	"capstone/entities"

	"github.com/stretchr/testify/mock"
)

type MockCityRepository struct {
	mock.Mock
}

func (m *MockCityRepository) Create(city *entities.City) error {
	args := m.Called(city)
	if args.Error(0) == nil {
		return nil
	}
	return args.Error(0)
}

func (m *MockCityRepository) FindAll() ([]entities.City, error) {
	args := m.Called()
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]entities.City), nil
}

func (m *MockCityRepository) FindById(id string) (*entities.City, error) {
	args := m.Called(id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*entities.City), nil
}

func (m *MockCityRepository) Update(city *entities.City) error {
	args := m.Called(city)
	if args.Error(0) == nil {
		return nil
	}
	return args.Error(0)
}

func (m *MockCityRepository) Delete(city *entities.City) error {
	args := m.Called(city)
	if args.Error(0) == nil {
		return nil
	}
	return args.Error(0)
}

func (m *MockCityRepository) GetCitiesWithDestinations() ([]entities.City, error) {
	args := m.Called()
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]entities.City), nil
}
