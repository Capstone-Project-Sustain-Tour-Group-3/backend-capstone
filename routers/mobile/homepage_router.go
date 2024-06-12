package routers

import (
	"capstone/config"
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
	redisClient := redis.NewRedisClient(
		config.ENV.REDIS_ADDR,
		config.ENV.REDIS_USERNAME,
		config.ENV.REDIS_PASSWORD,
	)

	usecase := usecases.NewHomepageUsecase(destinationRepo, routeRepo, redisClient)
	handler := handlers.NewHomePageHandler(usecase)

	r.GET("", handler.GetHomepageData)
}
