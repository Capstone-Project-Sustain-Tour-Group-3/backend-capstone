package routers

import (
	"capstone/config"
	"capstone/handlers"
	"capstone/repositories"
	"capstone/usecases"

	"github.com/labstack/echo/v4"
)

func DestinationRouter(r *echo.Group) {
	// r.Use(middlewares.JWTMiddleware)
	destinationRepo := repositories.NewDestinationRepository(config.DB)
	usecase := usecases.NewDestinationUsecase(destinationRepo)
	handler := handlers.NewDestinationHandler(usecase)
	r.GET("", handler.SearchDestinations)
	r.GET("/:id", handler.DetailDestination)
}
