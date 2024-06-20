package mobile

import (
	"capstone/config"
	"capstone/externals/openai"
	"capstone/externals/redis"
	"capstone/handlers"
	"capstone/middlewares"
	"capstone/repositories"
	"capstone/usecases"

	"github.com/labstack/echo/v4"
)

func HomepageRouter(r *echo.Group) {
	r.Use(middlewares.JWTMiddleware)

	destinationRepo := repositories.NewDestinationRepository(config.DB)
	routeRepo := repositories.NewRouteRepository(config.DB)
	userPersonalizationRepo := repositories.NewUserPersonalizationRepository(config.DB)
	redisClient := redis.NewRedisClient(
		config.ENV.REDIS_ADDR,
		config.ENV.REDIS_USERNAME,
		config.ENV.REDIS_PASSWORD,
	)
	openAIClient := openai.NewOpenAIClient(config.ENV.OPENAI_API_KEY)

	usecase := usecases.NewHomepageUsecase(
		destinationRepo,
		routeRepo,
		userPersonalizationRepo,
		redisClient,
		openAIClient,
	)
	handler := handlers.NewHomePageHandler(usecase)

	r.GET("", handler.GetHomepageData)
}
