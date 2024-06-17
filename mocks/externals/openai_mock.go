package externals

import (
	"github.com/stretchr/testify/mock"
)

type MockOpenAIClient struct {
	mock.Mock
}

func (m *MockOpenAIClient) GenerateAnswer(prompt, instruction string) (string, error) {
	args := m.Called(prompt, instruction)
	if args.Get(0) == nil {
		return "", args.Error(1)
	}
	return args.String(0), args.Error(1)
}
