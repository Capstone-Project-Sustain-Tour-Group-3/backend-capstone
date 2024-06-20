package repositories

import (
	"capstone/entities"

	"github.com/google/uuid"
	"github.com/stretchr/testify/mock"
	"gorm.io/gorm"
)

type MockDestinationMediaRepository struct {
	mock.Mock
}

func (m *MockDestinationMediaRepository) Create(destinationMedia *entities.DestinationMedia, tx *gorm.DB) error {
	args := m.Called(destinationMedia, tx)
	if args.Error(0) == nil {
		return nil
	}
	return args.Error(0)
}

func (m *MockDestinationMediaRepository) FindAll(page, limit int, searchQuery string) (*int64, []entities.DestinationMedia, error) {
	args := m.Called(page, limit, searchQuery)
	if args.Get(0) == nil {
		return nil, nil, args.Error(2)
	}
	return args.Get(0).(*int64), args.Get(1).([]entities.DestinationMedia), nil
}

func (m *MockDestinationMediaRepository) FindById(id uuid.UUID) (*entities.DestinationMedia, error) {
	args := m.Called(id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*entities.DestinationMedia), nil
}

func (m *MockDestinationMediaRepository) Update(destinationMedia *entities.DestinationMedia) error {
	args := m.Called(destinationMedia)
	if args.Error(0) == nil {
		return nil
	}
	return args.Error(0)
}

func (m *MockDestinationMediaRepository) Delete(destinationMedia *entities.DestinationMedia, tx *gorm.DB) error {
	args := m.Called(destinationMedia, tx)
	if args.Error(0) == nil {
		return nil
	}
	return args.Error(0)
}

func (m *MockDestinationMediaRepository) DeleteMany(destinationMedias *[]entities.DestinationMedia, tx *gorm.DB) error {
	args := m.Called(destinationMedias, tx)
	if args.Error(0) == nil {
		return nil
	}
	return args.Error(0)
}

func (m *MockDestinationMediaRepository) BeginTx() *gorm.DB {
	args := m.Called()
	if args.Get(0) == nil {
		return nil
	}
	return args.Get(0).(*gorm.DB)
}
