package repositories

import (
	"capstone/entities"

	"github.com/stretchr/testify/mock"
)

type MockUserRepository struct {
	mock.Mock
}

func (m *MockUserRepository) FindByEmail(email string) (*entities.User, error) {
	args := m.Called(email)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*entities.User), nil
}

func (m *MockUserRepository) GetUserByRefreshToken(refreshToken string) (*entities.User, error) {
	args := m.Called(refreshToken)
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

func (m *MockUserRepository) FindByUsername(name string) (*entities.User, error) {
	args := m.Called(name)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*entities.User), nil
}
