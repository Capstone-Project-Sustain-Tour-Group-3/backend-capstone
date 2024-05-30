package usecases

import (
	"capstone/dto"
	"capstone/externals/gemini"
	"capstone/externals/redis"
	"capstone/helpers"

	"github.com/google/uuid"
	"github.com/labstack/gommon/log"
)

type ChatbotUsecase interface {
	GenerateAnswer(question *dto.ChatbotRequest, uid uuid.UUID) *dto.ChatbotResponse
	FlushChatHistory(uid uuid.UUID) error
}

type chatbotUsecase struct {
	gemini gemini.GeminiClient
	redis  redis.RedisClient
}

func NewChatbotUsecase(gemini gemini.GeminiClient, redis redis.RedisClient) *chatbotUsecase {
	return &chatbotUsecase{
		gemini: gemini,
		redis:  redis,
	}
}

func (u *chatbotUsecase) GenerateAnswer(request *dto.ChatbotRequest, uid uuid.UUID) *dto.ChatbotResponse {
	hashedKey := helpers.EncodeRedisKey(uid.String())
	rawHistory, err := u.redis.GetChatHistory(hashedKey)
	if err != nil {
		log.Error(err)
	}

	convertedHistory := helpers.ToChatbotHistory(rawHistory)
	ans, err := u.gemini.AnswerChat(request.Message, convertedHistory)
	if err != nil {
		log.Error(err)
	}

	err = u.redis.UpdateChatHistory(hashedKey, request.Message)
	if err != nil {
		log.Error(err)
	}

	err = u.redis.UpdateChatHistory(hashedKey, ans)
	if err != nil {
		log.Error(err)
	}

	return &dto.ChatbotResponse{Answer: ans}
}

func (u *chatbotUsecase) FlushChatHistory(uid uuid.UUID) error {
	return u.redis.DeleteChatHistory(helpers.EncodeRedisKey(uid.String()))
}
