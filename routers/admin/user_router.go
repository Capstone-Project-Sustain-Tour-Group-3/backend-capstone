package admin

import (
	"capstone/config"
	"capstone/handlers"
	"capstone/repositories"
	"capstone/usecases"

	"github.com/labstack/echo/v4"
)

func UserRouter(r *echo.Group) {
	repository := repositories.NewUserRepository(config.DB)
	usecase := usecases.NewUserUsecase(repository)
	handler := handlers.NewUserHandler(usecase)

	r.GET("", handler.FindAll)
	r.POST("", handler.Create)
	r.GET("/:id", handler.FindById)
	r.PUT("/:id", handler.Update)
	r.DELETE("/:id", handler.Delete)
}
