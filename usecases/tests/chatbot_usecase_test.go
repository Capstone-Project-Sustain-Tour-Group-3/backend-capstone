package tests

import (
	"testing"

	"capstone/dto"
	"capstone/errorHandlers"
	"capstone/helpers"
	"capstone/mocks/externals"
	"capstone/usecases"

	"github.com/google/generative-ai-go/genai"
	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
)

type generateAnswerTestCase struct {
	name          string
	mockSetup     func(redisClient *externals.MockRedisClient, geminiClient *externals.MockGeminiClient)
	expectedResp  *dto.ChatbotResponse
	expectedError error
}

func TestGenerateAnswer(t *testing.T) {
	uid := uuid.New()
	hashedKey := helpers.EncodeRedisKey(uid.String())
	message := "Hello, how are you?"
	expectedAnswer := "I am fine, thank you!"
	expectedErrorAnswer := "Sorry i can't answer your question"
	request := &dto.ChatbotRequest{Message: message}
	history := []*genai.Content(nil)

	testCases := []generateAnswerTestCase{
		{
			name: "Success generate answer",
			mockSetup: func(redisClient *externals.MockRedisClient, geminiClient *externals.MockGeminiClient) {
				redisClient.On("GetChatHistory", hashedKey).Return(&[]string{}, nil)
				geminiClient.On("AnswerChat", message, history).Return(expectedAnswer, nil)
				redisClient.On("UpdateChatHistory", hashedKey, message).Return(nil)
				redisClient.On("UpdateChatHistory", hashedKey, expectedAnswer).Return(nil)
			},
			expectedResp: &dto.ChatbotResponse{Answer: expectedAnswer},
		},
		{
			name: "Failed to get chat history",
			mockSetup: func(redisClient *externals.MockRedisClient, geminiClient *externals.MockGeminiClient) {
				redisClient.On("GetChatHistory", hashedKey).Return(nil, &errorHandlers.InternalServerError{Message: "Failed to get chat history"})
				geminiClient.On("AnswerChat", message, history).Return(expectedAnswer, nil)
				redisClient.On("UpdateChatHistory", hashedKey, message).Return(nil)
				redisClient.On("UpdateChatHistory", hashedKey, expectedAnswer).Return(nil)
			},
			expectedResp:  &dto.ChatbotResponse{Answer: expectedAnswer},
			expectedError: &errorHandlers.InternalServerError{Message: "Failed to get chat history"},
		},
		{
			name: "Failed to get answer from Gemini",
			mockSetup: func(redisClient *externals.MockRedisClient, geminiClient *externals.MockGeminiClient) {
				redisClient.On("GetChatHistory", hashedKey).Return(&[]string{}, nil)
				geminiClient.On("AnswerChat", message, history).Return(expectedErrorAnswer, &errorHandlers.InternalServerError{Message: "Failed to get answer from Gemini"})
				redisClient.On("UpdateChatHistory", hashedKey, message).Return(nil)
				redisClient.On("UpdateChatHistory", hashedKey, expectedErrorAnswer).Return(nil)
			},
			expectedResp:  &dto.ChatbotResponse{Answer: expectedErrorAnswer},
			expectedError: &errorHandlers.InternalServerError{Message: "Failed to get answer from Gemini"},
		},
		{
			name: "Failed to update chat history with user message",
			mockSetup: func(redisClient *externals.MockRedisClient, geminiClient *externals.MockGeminiClient) {
				redisClient.On("GetChatHistory", hashedKey).Return(&[]string{}, nil)
				geminiClient.On("AnswerChat", message, history).Return(expectedAnswer, nil)
				redisClient.On("UpdateChatHistory", hashedKey, message).Return(&errorHandlers.InternalServerError{Message: "Failed to update chat history with user message"})
				redisClient.On("UpdateChatHistory", hashedKey, expectedAnswer).Return(nil)
			},
			expectedResp:  &dto.ChatbotResponse{Answer: expectedAnswer},
			expectedError: &errorHandlers.InternalServerError{Message: "Failed to update chat history with user message"},
		},
		{
			name: "Failed to update chat history with answer",
			mockSetup: func(redisClient *externals.MockRedisClient, geminiClient *externals.MockGeminiClient) {
				redisClient.On("GetChatHistory", hashedKey).Return(&[]string{}, nil)
				geminiClient.On("AnswerChat", message, history).Return(expectedAnswer, nil)
				redisClient.On("UpdateChatHistory", hashedKey, message).Return(nil)
				redisClient.On("UpdateChatHistory", hashedKey, expectedAnswer).Return(&errorHandlers.InternalServerError{Message: "Failed to update chat history with answer"})
			},
			expectedResp:  &dto.ChatbotResponse{Answer: expectedAnswer},
			expectedError: &errorHandlers.InternalServerError{Message: "Failed to update chat history with answer"},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			redisClient := new(externals.MockRedisClient)
			geminiClient := new(externals.MockGeminiClient)
			tc.mockSetup(redisClient, geminiClient)

			uc := usecases.NewChatbotUsecase(geminiClient, redisClient)
			resp := uc.GenerateAnswer(request, uid)

			require.NotNil(t, resp)
			require.Equal(t, tc.expectedResp, resp)
		})
	}
}

func TestFlushChatHistory(t *testing.T) {
	redisClient := new(externals.MockRedisClient)
	geminiClient := new(externals.MockGeminiClient)

	uid := uuid.New()
	encodedKey := helpers.EncodeRedisKey(uid.String())

	redisClient.On("DeleteChatHistory", encodedKey).Return(nil)

	uc := usecases.NewChatbotUsecase(geminiClient, redisClient)
	err := uc.FlushChatHistory(uid)

	require.NoError(t, err)
}
