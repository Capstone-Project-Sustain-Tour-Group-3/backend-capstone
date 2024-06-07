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
	usecase := usecases.NewRouteUsecase(repository)
	handler := handlers.NewRouteHandler(usecase)
	r.GET("", handler.FindAll)
	r.GET("/:id", handler.FindById)
}
