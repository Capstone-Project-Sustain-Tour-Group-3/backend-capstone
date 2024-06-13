package admin

import (
	"capstone/config"
	"capstone/handlers"
	"capstone/middlewares"
	"capstone/repositories"
	"capstone/usecases"
	"github.com/labstack/echo/v4"
)

func DashboardRouter(r *echo.Group) {
	r.Use(middlewares.JWTMiddleware)
	r.Use(middlewares.RoleMiddleware("admin", "superadmin"))

	repository := repositories.NewDashboardRepository(config.DB)
	usecase := usecases.NewDashboardUsecase(repository)
	handler := handlers.NewDashboardHandler(usecase)
	r.GET("", handler.GetAllData)
}
