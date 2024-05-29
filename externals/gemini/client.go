package gemini

import (
	"capstone/dto"

	"github.com/google/generative-ai-go/genai"
)

type GeminiClient interface {
	AnswerQuestion(r *dto.ChatbotRequest) (string, error)
}

type geminiClient struct {
	c *genai.Client
}

func NewGeminiClient(c *genai.Client) *geminiClient {
	return &geminiClient{
		c: c,
	}
}

func (g *geminiClient) AnswerQuestion(r *dto.ChatbotRequest) (string, error) {
	panic("not implemented") // TODO: Implement
}
