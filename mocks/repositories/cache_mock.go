package repositories

import "github.com/stretchr/testify/mock"

type MockCacheRepository struct {
	mock.Mock
}

func (m *MockCacheRepository) Set(name, val string) {
	m.Called(name, val)
}

func (m *MockCacheRepository) Get(name string) (string, bool) {
	args := m.Called(name)
	if args.Get(0) == nil {
		return "", args.Bool(1)
	}
	return args.Get(0).(string), args.Bool(1)
}
