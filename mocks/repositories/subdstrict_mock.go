package repositories

import (
	"capstone/entities"

	"github.com/stretchr/testify/mock"
)

type MockSubdistrictRepository struct {
	mock.Mock
}

func (m *MockSubdistrictRepository) Create(subdistrict *entities.Subdistrict) error {
	args := m.Called(subdistrict)
	if args.Error(0) == nil {
		return nil
	}
	return args.Error(0)
}

func (m *MockSubdistrictRepository) FindAll() ([]entities.Subdistrict, error) {
	args := m.Called()
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]entities.Subdistrict), nil
}

func (m *MockSubdistrictRepository) FindById(id string) (*entities.Subdistrict, error) {
	args := m.Called(id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*entities.Subdistrict), nil
}

func (m *MockSubdistrictRepository) Update(subdistrict *entities.Subdistrict) error {
	args := m.Called(subdistrict)
	if args.Error(0) == nil {
		return nil
	}
	return args.Error(0)
}

func (m *MockSubdistrictRepository) Delete(subdistrict *entities.Subdistrict) error {
	args := m.Called(subdistrict)
	if args.Error(0) == nil {
		return nil
	}
	return args.Error(0)
}
