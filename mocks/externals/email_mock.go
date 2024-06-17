package externals

import "github.com/stretchr/testify/mock"

type MockEmailClient struct {
	mock.Mock
}

func (m *MockEmailClient) SendOTP(to, subject, body string) error {
	args := m.Called(to, subject, body)
	return args.Error(0)
}
