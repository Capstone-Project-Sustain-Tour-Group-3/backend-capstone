package repositories

import (
	"capstone/entities"

	"github.com/google/uuid"
	"github.com/stretchr/testify/mock"
)

type MockUserRepository struct {
	mock.Mock
}

func (m *MockUserRepository) FindAll(page, limit int, searchQuery string) (*[]entities.User, *int64, error) {
	args := m.Called(page, limit, searchQuery)
	if args.Get(0) == nil {
		return nil, nil, args.Error(2)
	}
	return args.Get(0).(*[]entities.User), args.Get(1).(*int64), nil
}

func (m *MockUserRepository) FindByUsername(username string) (*entities.User, error) {
	args := m.Called(username)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*entities.User), nil
}

func (m *MockUserRepository) FindByEmail(email string) (*entities.User, error) {
	args := m.Called(email)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*entities.User), nil
}

func (m *MockUserRepository) FindById(id uuid.UUID) (*entities.User, error) {
	args := m.Called(id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*entities.User), nil
}

func (m *MockUserRepository) Create(user *entities.User) error {
	args := m.Called(user)
	if args.Error(0) == nil {
		return nil
	}
	return args.Error(0)
}

func (m *MockUserRepository) Update(user *entities.User) error {
	args := m.Called(user)
	if args.Error(0) == nil {
		return nil
	}
	return args.Error(0)
}

func (m *MockUserRepository) Delete(user *entities.User) error {
	args := m.Called(user)
	if args.Error(0) == nil {
		return nil
	}
	return args.Error(0)
}
