package gemini

import (
	"context"
	"log"

	"capstone/helpers"

	"github.com/google/generative-ai-go/genai"
	"google.golang.org/api/option"
)

type GeminiClient interface {
	AnswerChat(request string, history []*genai.Content) (string, error)
}

type geminiClient struct {
	client *genai.Client
}

func NewGeminiClient(apiKey string) *geminiClient {
	ctx := context.Background()
	client, err := genai.NewClient(ctx, option.WithAPIKey(apiKey))
	if err != nil {
		log.Fatal("cannot create gemini client :: ", err)
	}

	return &geminiClient{client: client}
}

func (g *geminiClient) AnswerChat(request string, history []*genai.Content) (string, error) {
	ctx := context.Background()

	model := g.client.GenerativeModel("gemini-1.5-flash")
	model.SystemInstruction = helpers.GetSystemInstruction()

	cs := model.StartChat()
	cs.History = history

	ans, err := cs.SendMessage(ctx, genai.Text(request))
	if err != nil {
		return "Mohon maaf saya sedang tidak bisa merespon. Silakan coba kembali nanti.", err
	}

	res := helpers.ToChatbotResponse(ans)

	return res, nil
}
