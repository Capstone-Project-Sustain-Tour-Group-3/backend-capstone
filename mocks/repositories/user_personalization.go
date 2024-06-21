package repositories

import (
	"capstone/entities"

	"github.com/google/uuid"
	"github.com/stretchr/testify/mock"
)

type MockUserPersonalizationRepository struct {
	mock.Mock
}

func (m *MockUserPersonalizationRepository) Create(userPersonalization *entities.UserPersonalization) error {
	args := m.Called(userPersonalization)
	if args.Error(0) == nil {
		return nil
	}
	return args.Error(0)
}

func (m *MockUserPersonalizationRepository) FindByUserId(uid uuid.UUID) (*entities.UserPersonalization, error) {
	args := m.Called(uid)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*entities.UserPersonalization), nil
}
