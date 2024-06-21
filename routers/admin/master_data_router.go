package admin

import (
	"capstone/config"
	"capstone/handlers"
	"capstone/middlewares"
	"capstone/repositories"
	"capstone/usecases"

	"github.com/labstack/echo/v4"
)

func MasterDataRouter(r *echo.Group) {
	r.Use(middlewares.JWTMiddleware)
	r.Use(middlewares.RoleMiddleware("superadmin"))

	categoryRepo := repositories.NewCategoryRepository(config.DB)
	facilityRepo := repositories.NewFacilityRepository(config.DB)
	provinceRepo := repositories.NewProvinceRepository(config.DB)
	cityRepo := repositories.NewCityRepository(config.DB)
	subdistrictRepo := repositories.NewSubdistrictRepository(config.DB)

	usecase := usecases.NewMasterDataUsecase(provinceRepo, cityRepo, subdistrictRepo, categoryRepo, facilityRepo)
	handler := handlers.NewMasterDataHandler(usecase)

	r.GET("/categories", handler.GetCategories)
	r.GET("/facilities", handler.GetFacilities)
	r.GET("/provinces", handler.GetProvinces)
	r.GET("/cities", handler.GetCities)
	r.GET("/subdistricts", handler.GetSubdistricts)
}
