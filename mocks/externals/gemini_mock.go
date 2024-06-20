package externals

import (
	"github.com/google/generative-ai-go/genai"
	"github.com/stretchr/testify/mock"
)

type MockGeminiClient struct {
	mock.Mock
}

func (m *MockGeminiClient) AnswerChat(request string, history []*genai.Content) (string, error) {
	args := m.Called(request, history)
	if args.Error(1) != nil {
		return "Sorry i can't answer your question", args.Error(1)
	}
	return args.Get(0).(string), args.Error(1)
}
