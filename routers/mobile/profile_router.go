package routers

import (
	"capstone/config"
	"capstone/handlers"
	"capstone/middlewares"
	"capstone/repositories"
	"capstone/usecases"
	"github.com/labstack/echo/v4"
)

func ProfileRouter(r *echo.Group) {
	r.Use(middlewares.JWTMiddleware)
	repository := repositories.NewUserRepository(config.DB)
	usecase := usecases.NewProfileUsecase(repository)
	handler := handlers.NewProfileHandler(usecase)
	r.GET("", handler.GetDetailUser)
}
