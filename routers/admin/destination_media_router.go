package admin

import (
	"capstone/config"
	"capstone/handlers"
	"capstone/repositories"
	"capstone/usecases"
	"github.com/labstack/echo/v4"
)

func DestinationMediaRouter(r *echo.Group) {
	//r.Use(middlewares.JWTMiddleware)

	destinationRepo := repositories.NewDestinationRepository(config.DB)
	destinationMediaRepo := repositories.NewDestinationMediaRepository(config.DB)

	usecase := usecases.NewDestinationMediaUsecase(destinationMediaRepo, destinationRepo)

	handler := handlers.NewDestinationMediaHandler(usecase)

	r.POST("", handler.CreateDestinationMedia)
	r.GET("", handler.AllDestinationMedia)
	r.GET("/:id", handler.DetailDestinationMedia)
	r.PUT("/:id", handler.UpdateDestinationMedia)
	r.DELETE("/:id", handler.DeleteDestinationMedia)
}
