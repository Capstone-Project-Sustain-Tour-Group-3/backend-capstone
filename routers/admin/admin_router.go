package admin

import (
	"capstone/config"
	"capstone/handlers"
	"capstone/middlewares"
	"capstone/repositories"
	"capstone/usecases"

	"github.com/labstack/echo/v4"
)

func ManageAdminRouter(r *echo.Group) {
	r.Use(middlewares.JWTMiddleware)
	r.Use(middlewares.RoleMiddleware("superadmin"))

	repository := repositories.NewAdminRepository(config.DB)
	usecase := usecases.NewAdminUsecase(repository)
	handler := handlers.NewAdminHandler(usecase)

	r.GET("", handler.GetAllAdmins)
}
