package admin

import (
	"capstone/config"
	"capstone/handlers"
	"capstone/repositories"
	"capstone/usecases"

	"github.com/labstack/echo/v4"
)

func RouteRouter(r *echo.Group) {
	repository := repositories.NewRouteRepository(config.DB)
	routeRepo := repositories.NewRouteDetailRepository(config.DB)
	usecase := usecases.NewRouteUsecase(repository, routeRepo)
	handler := handlers.NewRouteHandler(usecase)
	r.GET("", handler.FindAll)
	r.GET("/:id", handler.FindById)
	r.DELETE("/:id", handler.DeleteRoute)
}
