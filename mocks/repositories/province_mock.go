package repositories

import (
	"capstone/entities"

	"github.com/stretchr/testify/mock"
)

type MockProvinceRepository struct {
	mock.Mock
}

func (m *MockProvinceRepository) Create(province *entities.Province) error {
	args := m.Called(province)
	if args.Error(0) == nil {
		return nil
	}
	return args.Error(0)
}

func (m *MockProvinceRepository) FindAll() ([]entities.Province, error) {
	args := m.Called()
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]entities.Province), nil
}

func (m *MockProvinceRepository) FindById(id string) (*entities.Province, error) {
	args := m.Called(id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*entities.Province), nil
}

func (m *MockProvinceRepository) Update(province *entities.Province) error {
	args := m.Called(province)
	if args.Error(0) == nil {
		return nil
	}
	return args.Error(0)
}

func (m *MockProvinceRepository) Delete(province *entities.Province) error {
	args := m.Called(province)
	if args.Error(0) == nil {
		return nil
	}
	return args.Error(0)
}
