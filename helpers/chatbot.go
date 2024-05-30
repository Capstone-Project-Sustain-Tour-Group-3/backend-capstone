package helpers

import (
	"fmt"

	"github.com/google/generative-ai-go/genai"
)

func GetSystemInstruction() *genai.Content {
	prompt := "kamu adalah virtual assistant dengan karakteristik ceria dan tidak membosankan. kamu bisa memberikan rekomendasi tempat wisata yang ada di wilayah Indonesia (alam, seni & budaya, dan sejarah) dan menjawab pertanyaan seputar travelling dengan jawaban yang singkat, pdat, dan jelas. namun selain itu kamu tidak akan bisa menjawab pertanyaan tersebut."

	return &genai.Content{
		Parts: []genai.Part{genai.Text(prompt)},
	}
}

func ToChatbotHistory(history *[]string) []*genai.Content {
	var content []*genai.Content
	var role string

	if history == nil {
		return content
	}

	for i := len(*history) - 1; i >= 0; i-- {
		role = "user"
		if i%2 == 0 {
			role = "model"
		}

		content = append(content, &genai.Content{
			Parts: []genai.Part{genai.Text((*history)[i])},
			Role:  role,
		})
	}

	return content
}

func ToChatbotResponse(response *genai.GenerateContentResponse) string {
	var processedMessage string
	for _, cand := range response.Candidates {
		if cand.Content != nil {
			for _, part := range cand.Content.Parts {
				processedMessage += fmt.Sprintf("%s", part)
			}
		}
	}

	return processedMessage
}
