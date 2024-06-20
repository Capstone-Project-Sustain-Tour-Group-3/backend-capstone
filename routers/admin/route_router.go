package admin

import (
	"capstone/config"
	"capstone/externals/openai"
	"capstone/handlers"
	"capstone/middlewares"
	"capstone/repositories"
	"capstone/usecases"

	"github.com/labstack/echo/v4"
)

func RouteRouter(r *echo.Group) {
	r.Use(middlewares.JWTMiddleware)
	r.Use(middlewares.RoleMiddleware("admin", "superadmin"))

	repository := repositories.NewRouteRepository(config.DB)
	cityRepo := repositories.NewCityRepository(config.DB)
	routeRepo := repositories.NewRouteDetailRepository(config.DB)
	destinationRepo := repositories.NewDestinationRepository(config.DB)
	openAIClient := openai.NewOpenAIClient(config.ENV.OPENAI_API_KEY)

	usecase := usecases.NewRouteUsecase(cityRepo, destinationRepo, repository, routeRepo, openAIClient)
	handler := handlers.NewRouteHandler(usecase)
	r.GET("", handler.FindAll)
	r.GET("/:id", handler.FindById)
	r.DELETE("/:id", handler.DeleteRoute)
}
