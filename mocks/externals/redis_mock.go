package externals

import "github.com/stretchr/testify/mock"

type MockRedisClient struct {
	mock.Mock
}

func (m *MockRedisClient) UpdateChatHistory(key string, value string) error {
	args := m.Called(key, value)
	if args.Get(0) == nil {
		return nil
	}
	return args.Error(0)
}

func (m *MockRedisClient) GetChatHistory(key string) (*[]string, error) {
	args := m.Called(key)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*[]string), args.Error(1)
}

func (m *MockRedisClient) DeleteChatHistory(key string) error {
	args := m.Called(key)
	if args.Get(0) == nil {
		return nil
	}
	return args.Error(0)
}

func (m *MockRedisClient) GetRecommendedDestinationsIds(key string) (*[]string, error) {
	args := m.Called(key)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*[]string), args.Error(1)
}

func (m *MockRedisClient) SetRecommendedDestinationsIds(key string, values []string) error {
	args := m.Called(key, values)
	if args.Get(0) == nil {
		return nil
	}
	return args.Error(0)
}
