package mocks

import (
	"capstone/entities"

	"github.com/stretchr/testify/mock"
)

type MockAuthRepository struct {
	mock.Mock
}

func (m *MockAuthRepository) FindByEmail(email string) (*entities.User, error) {
	args := m.Called(email)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*entities.User), nil
}

func (m *MockAuthRepository) GetUserByRefreshToken(refreshToken string) (*entities.User, error) {
	args := m.Called(refreshToken)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*entities.User), nil
}

func (m *MockAuthRepository) Create(user *entities.User) error {
	args := m.Called(user)
	if args.Error(0) == nil {
		return nil
	}
	return args.Error(0)
}

func (m *MockAuthRepository) Update(user *entities.User) error {
	args := m.Called(user)
	if args.Error(0) == nil {
		return nil
	}
	return args.Error(0)
}

func (m *MockAuthRepository) FindByUsername(name string) (*entities.User, error) {
	args := m.Called(name)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*entities.User), nil
}
