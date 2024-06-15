package externals

import "github.com/stretchr/testify/mock"

type MockTokenHelper struct {
	mock.Mock
}

func (m *MockTokenHelper) GenerateAccessToken(user interface{}) (string, error) {
	args := m.Called(user)
	return args.String(0), args.Error(1)
}

func (m *MockTokenHelper) GenerateRefreshToken(user interface{}) (string, error) {
	args := m.Called(user)
	return args.String(0), args.Error(1)
}
