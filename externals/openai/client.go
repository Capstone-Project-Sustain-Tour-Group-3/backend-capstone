package openai

import (
	"context"

	oai "github.com/sashabaranov/go-openai"
)

type OpenAIClient interface {
	GenerateAnswer(prompt, instruction string) (string, error)
}

type openaiClient struct {
	client *oai.Client
}

func NewOpenAIClient(apiKey string) *openaiClient {
	return &openaiClient{
		client: oai.NewClient(apiKey),
	}
}

func (o *openaiClient) GenerateAnswer(prompt, instruction string) (string, error) {
	res, err := o.client.CreateChatCompletion(
		context.Background(),
		oai.ChatCompletionRequest{
			Model: oai.GPT3Dot5Turbo,
			Messages: []oai.ChatCompletionMessage{
				{
					Role:    oai.ChatMessageRoleSystem,
					Content: instruction,
				},
				{
					Role:    oai.ChatMessageRoleUser,
					Content: prompt,
				},
			},
			Temperature: 0,
		},
	)
	if err != nil {
		return "", err
	}

	return res.Choices[0].Message.Content, nil
}
