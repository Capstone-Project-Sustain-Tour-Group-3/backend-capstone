package externals

import "github.com/stretchr/testify/mock"

type MockPasswordHelper struct {
	mock.Mock
}

func (m *MockPasswordHelper) HashPassword(password string) (string, error) {
	args := m.Called(password)
	return args.String(0), args.Error(1)
}

func (m *MockPasswordHelper) VerifyPassword(hashedPassword, password string) error {
	args := m.Called(hashedPassword, password)
	return args.Error(0)
}
