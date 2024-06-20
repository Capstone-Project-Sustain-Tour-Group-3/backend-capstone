package repositories

import (
	"capstone/dto"

	"github.com/stretchr/testify/mock"
)

type MockDashboardRepository struct {
	mock.Mock
}

func (m *MockDashboardRepository) GetTotalAdmin() (int64, error) {
	args := m.Called()
	if args.Get(0) == nil {
		return 0, args.Error(1)
	}
	return args.Get(0).(int64), args.Error(1)
}

func (m *MockDashboardRepository) GetTotalUser() (int64, error) {
	args := m.Called()
	if args.Get(0) == nil {
		return 0, args.Error(1)
	}
	return args.Get(0).(int64), args.Error(1)
}

func (m *MockDashboardRepository) GetTotalRoute() (int64, error) {
	args := m.Called()
	if args.Get(0) == nil {
		return 0, args.Error(1)
	}
	return args.Get(0).(int64), args.Error(1)
}

func (m *MockDashboardRepository) GetTotalDestination() (int64, error) {
	args := m.Called()
	if args.Get(0) == nil {
		return 0, args.Error(1)
	}
	return args.Get(0).(int64), args.Error(1)
}

func (m *MockDashboardRepository) GetTotalVidContent() (int64, error) {
	args := m.Called()
	if args.Get(0) == nil {
		return 0, args.Error(1)
	}
	return args.Get(0).(int64), args.Error(1)
}

func (m *MockDashboardRepository) GetTotalDestinationWithCategory() ([]dto.TotalDestinationationWithCategory, error) {
	args := m.Called()
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]dto.TotalDestinationationWithCategory), args.Error(1)
}

func (m *MockDashboardRepository) GetMonthlyUsers() ([]dto.MonthlyUser, error) {
	args := m.Called()
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]dto.MonthlyUser), args.Error(1)
}

func (m *MockDashboardRepository) GetRouteUser() ([]dto.RouteUser, error) {
	args := m.Called()
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]dto.RouteUser), args.Error(1)
}
