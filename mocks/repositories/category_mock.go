package repositories

import (
	"capstone/entities"

	"github.com/google/uuid"
	"github.com/stretchr/testify/mock"
)

type MockCategoryRepository struct {
	mock.Mock
}

func (m *MockCategoryRepository) Create(category *entities.Category) error {
	args := m.Called(category)
	if args.Error(0) == nil {
		return nil
	}
	return args.Error(0)
}

func (m *MockCategoryRepository) FindAll() ([]entities.Category, error) {
	args := m.Called()
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]entities.Category), nil
}

func (m *MockCategoryRepository) FindById(id uuid.UUID) (*entities.Category, error) {
	args := m.Called(id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*entities.Category), nil
}

func (m *MockCategoryRepository) Update(category *entities.Category) error {
	args := m.Called(category)
	if args.Error(0) == nil {
		return nil
	}
	return args.Error(0)
}

func (m *MockCategoryRepository) Delete(category *entities.Category) error {
	args := m.Called(category)
	if args.Error(0) == nil {
		return nil
	}
	return args.Error(0)
}
