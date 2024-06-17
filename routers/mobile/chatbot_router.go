package mobile

import (
	"capstone/config"
	exgemini "capstone/externals/gemini"
	exredis "capstone/externals/redis"
	"capstone/handlers"
	"capstone/middlewares"
	"capstone/usecases"

	"github.com/gorilla/websocket"
	"github.com/labstack/echo/v4"
)

func ChatbotRouter(r *echo.Group) {
	r.Use(middlewares.JWTMiddleware)

	gemini := exgemini.NewGeminiClient(config.ENV.GEMINI_API_KEY)
	redis := exredis.NewRedisClient(
		config.ENV.REDIS_ADDR,
		config.ENV.REDIS_USERNAME,
		config.ENV.REDIS_PASSWORD,
	)

	usecase := usecases.NewChatbotUsecase(gemini, redis)
	upgrader := websocket.Upgrader{}

	handler := handlers.NewChatbotHandler(usecase, upgrader)

	r.GET("", handler.Chatbot)
}
