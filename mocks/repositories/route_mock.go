package repositories

import (
	"capstone/entities"

	"github.com/google/uuid"
	"github.com/stretchr/testify/mock"
	"gorm.io/gorm"
)

type MockRouteRepository struct {
	mock.Mock
}

func (m *MockRouteRepository) FindAll(page, limit int, searchQuery string) (*[]entities.Route, *int64, error) {
	args := m.Called(page, limit, searchQuery)
	if args.Get(0) == nil {
		return nil, nil, args.Error(2)
	}
	return args.Get(0).(*[]entities.Route), args.Get(1).(*int64), args.Error(2)
}

func (m *MockRouteRepository) FindById(id uuid.UUID) (*entities.Route, error) {
	args := m.Called(id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*entities.Route), args.Error(1)
}

func (m *MockRouteRepository) Delete(route *entities.Route) error {
	args := m.Called(route)
	if args.Error(0) == nil {
		return nil
	}
	return args.Error(0)
}

func (m *MockRouteRepository) FindVisitedByUserSubquery(uid uuid.UUID) *gorm.DB {
	args := m.Called(uid)
	if args.Get(0) == nil {
		return nil
	}
	return args.Get(0).(*gorm.DB)
}
