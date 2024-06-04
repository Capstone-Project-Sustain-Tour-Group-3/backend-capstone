package routers

import (
	"capstone/config"
	"capstone/handlers"
	"capstone/middlewares"
	"capstone/repositories"
	"capstone/usecases"

	"github.com/labstack/echo/v4"
)

func PersonalizationRoute(r *echo.Group) {
	r.Use(middlewares.JWTMiddleware)

	provinceRepo := repositories.NewProvinceRepository(config.DB)
	categoryRepo := repositories.NewCategoryRepository(config.DB)
	userPersonalizationRepo := repositories.NewUserPersonalizationRepository(config.DB)

	usecase := usecases.NewPersonalizationUsecase(
		provinceRepo,
		categoryRepo,
		userPersonalizationRepo,
	)

	handler := handlers.NewPersonalizationHandler(usecase)

	r.POST("", handler.CreatePersonalization)
	r.GET("/categories", handler.GetCategories)
	r.GET("/provinces", handler.GetProvinces)
}
