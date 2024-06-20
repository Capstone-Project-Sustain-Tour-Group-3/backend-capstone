package repositories

import (
	"capstone/entities"

	"github.com/google/uuid"
	"github.com/stretchr/testify/mock"
	"gorm.io/gorm"
)

type MockDestinationRepository struct {
	mock.Mock
}

func (m *MockDestinationRepository) BeginTx() *gorm.DB {
	args := m.Called()
	if args.Get(0) == nil {
		return nil
	}
	return args.Get(0).(*gorm.DB)
}

func (m *MockDestinationRepository) FindById(id uuid.UUID) (*entities.Destination, error) {
	args := m.Called(id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*entities.Destination), args.Error(1)
}

func (m *MockDestinationRepository) FindByIdInCityId(id uuid.UUID, cityId string) (*entities.Destination, error) {
	args := m.Called(id, cityId)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*entities.Destination), args.Error(1)
}

func (m *MockDestinationRepository) FindAll(page, limit int, searchQuery, sortQuery, filterQuery string) (string, *int64, []entities.Destination, error) {
	args := m.Called(page, limit, searchQuery, sortQuery, filterQuery)
	if args.Get(0) == "" {
		return "", nil, nil, args.Error(3)
	}
	return args.String(0), args.Get(1).(*int64), args.Get(2).([]entities.Destination), args.Error(3)
}

func (m *MockDestinationRepository) FindByCategoryId(id uuid.UUID) ([]entities.Destination, error) {
	args := m.Called(id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]entities.Destination), args.Error(1)
}

func (m *MockDestinationRepository) Create(destination *entities.Destination, tx *gorm.DB) error {
	args := m.Called(destination, tx)
	if args.Get(0) == nil {
		return args.Error(1)
	}
	return args.Error(1)
}

func (m *MockDestinationRepository) Update(destination *entities.Destination, tx *gorm.DB) error {
	args := m.Called(destination, tx)
	if args.Get(0) == nil {
		return args.Error(0)
	}
	return args.Error(0)
}

func (m *MockDestinationRepository) Delete(destination *entities.Destination, tx *gorm.DB) error {
	args := m.Called(destination, tx)
	if args.Get(0) == nil {
		return args.Error(1)
	}
	return args.Error(1)
}

func (m *MockDestinationRepository) FindDestinationByCityId(id string) ([]entities.Destination, error) {
	args := m.Called(id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]entities.Destination), args.Error(1)
}

func (m *MockDestinationRepository) FindByProvinceName(name string, limit int) (*[]entities.Destination, error) {
	args := m.Called(name, limit)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*[]entities.Destination), args.Error(1)
}

func (m *MockDestinationRepository) FindPopularDestinationVideos() (*[]entities.Destination, error) {
	args := m.Called()
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*[]entities.Destination), args.Error(1)
}

func (m *MockDestinationRepository) FindBySubqueryRoute(subquery *gorm.DB, isVisited bool, provinceIds []string, categoryIds []uuid.UUID) (*[]entities.Destination, error) {
	args := m.Called(subquery, isVisited, provinceIds, categoryIds)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*[]entities.Destination), args.Error(1)
}

func (m *MockDestinationRepository) FindByManyIds(ids []string) (*[]entities.Destination, error) {
	args := m.Called(ids)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*[]entities.Destination), args.Error(1)
}
