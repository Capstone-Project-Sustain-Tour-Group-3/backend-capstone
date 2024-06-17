package repositories

import (
	"capstone/entities"

	"github.com/google/uuid"
	"github.com/stretchr/testify/mock"
)

type MockAdminRepository struct {
	mock.Mock
}

func (m *MockAdminRepository) FindAll(offset, limit int, search string) (*[]entities.Admin, *int64, error) {
	args := m.Called(offset, limit, search)
	if args.Get(0) == nil {
		return nil, nil, args.Error(2)
	}
	return args.Get(0).(*[]entities.Admin), args.Get(1).(*int64), args.Error(2)
}

func (m *MockAdminRepository) FindById(id uuid.UUID) (*entities.Admin, error) {
	args := m.Called(id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*entities.Admin), args.Error(1)
}

func (r *MockAdminRepository) FindByUsername(username string) (*entities.Admin, error) {
	args := r.Called(username)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*entities.Admin), args.Error(1)
}

func (r *MockAdminRepository) GetUserByRefreshToken(refreshToken string) (*entities.Admin, error) {
	args := r.Called(refreshToken)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*entities.Admin), args.Error(1)
}

func (m *MockAdminRepository) Update(admin *entities.Admin) error {
	args := m.Called(admin)
	if args.Error(0) == nil {
		return nil
	}
	return args.Error(0)
}

func (m *MockAdminRepository) Create(admin *entities.Admin) error {
	args := m.Called(admin)
	if args.Error(0) == nil {
		return nil
	}
	return args.Error(0)
}

func (m *MockAdminRepository) Delete(admin *entities.Admin) error {
	args := m.Called(admin)
	if args.Error(0) == nil {
		return nil
	}
	return args.Error(0)
}
