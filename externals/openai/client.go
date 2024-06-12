package openai

import (
	"capstone/helpers"
	"context"

	oai "github.com/sashabaranov/go-openai"
)

type OpenAIClient interface {
	GetRecommendation(prompt string) (string, error)
}

type openaiClient struct {
	client *oai.Client
}

func NewOpenAIClient(apiKey string) *openaiClient {
	return &openaiClient{
		client: oai.NewClient(apiKey),
	}
}

func (o *openaiClient) GetRecommendation(prompt string) (string, error) {
	res, err := o.client.CreateChatCompletion(
		context.Background(),
		oai.ChatCompletionRequest{
			Model: oai.GPT3Dot5Turbo,
			Messages: []oai.ChatCompletionMessage{
				{
					Role:    oai.ChatMessageRoleSystem,
					Content: helpers.GetRecommendationSystemInstruction(),
				},
				{
					Role:    oai.ChatMessageRoleUser,
					Content: prompt,
				},
			},
		},
	)

	if err != nil {
		return "", err
	}

	return res.Choices[0].Message.Content, nil
}
