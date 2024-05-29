package handlers

import (
	"capstone/externals/gemini"
)

type chatbotHandler struct {
	g gemini.GeminiClient
}

func NewChatbotHandler(g gemini.GeminiClient) *chatbotHandler {
	return &chatbotHandler{
		g: g,
	}
}
