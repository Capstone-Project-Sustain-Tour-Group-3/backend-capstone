package routers

import (
	"capstone/config"
	"capstone/handlers"
	"capstone/middlewares"
	"capstone/repositories"
	"capstone/usecases"

	"github.com/labstack/echo/v4"
)

func RouteRouter(r *echo.Group) {
	r.Use(middlewares.JWTMiddleware)

	repository := repositories.NewRouteRepository(config.DB)
	routeRepo := repositories.NewRouteDetailRepository(config.DB)
	usecase := usecases.NewRouteUsecase(repository, routeRepo)
	handler := handlers.NewRouteHandler(usecase)
	r.POST("", handler.SaveRoute)
	r.DELETE("/:id", handler.DeleteRoute)
}
