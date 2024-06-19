package repositories

import (
	"capstone/entities"

	"github.com/stretchr/testify/mock"
)

type MockRouteDetailRepository struct {
	mock.Mock
}

func (m *MockRouteDetailRepository) Create(routeDetail *entities.RouteDetail) error {
	args := m.Called(routeDetail)
	if args.Error(0) == nil {
		return nil
	}
	return args.Error(0)
}

func (m *MockRouteDetailRepository) DeleteMany(routeDetail *[]entities.RouteDetail) error {
	args := m.Called(routeDetail)
	if args.Error(0) == nil {
		return nil
	}
	return args.Error(0)
}
