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
	cacheRepo := repositories.NewCacheRepository()
	usecase := usecases.NewProfileUsecase(repository, cacheRepo)
	handler := handlers.NewProfileHandler(usecase)
	r.GET("", handler.GetDetailUser)
	r.PUT("", handler.InsertUserDetail)
}
