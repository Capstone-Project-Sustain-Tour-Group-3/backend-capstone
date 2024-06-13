package admin

import (
	"capstone/config"
	"capstone/externals/cloudinary"
	"capstone/handlers"
	"capstone/middlewares"
	"capstone/repositories"
	"capstone/usecases"

	"github.com/labstack/echo/v4"
)

func DestinationRouter(r *echo.Group) {
	r.Use(middlewares.JWTMiddleware)
	r.Use(middlewares.RoleMiddleware("admin", "superadmin"))

	destinationRepo := repositories.NewDestinationRepository(config.DB)
	destinationFacilityRepo := repositories.NewDestinationFacilityRepository(config.DB)
	destinationMediaRepo := repositories.NewDestinationMediaRepository(config.DB)
	destinationAddressRepo := repositories.NewDestinationAddressRepository(config.DB)
	categoryRepo := repositories.NewCategoryRepository(config.DB)
	facilityRepo := repositories.NewFacilityRepository(config.DB)
	provinceRepo := repositories.NewProvinceRepository(config.DB)
	cityRepo := repositories.NewCityRepository(config.DB)
	subdistrictRepo := repositories.NewSubdistrictRepository(config.DB)
	cloudinaryClient := cloudinary.NewCloudinaryClient(config.ENV.CLOUDINARY_URL)

	usecase := usecases.NewDestinationUsecase(
		destinationRepo,
		destinationFacilityRepo,
		destinationMediaRepo,
		destinationAddressRepo,
		categoryRepo,
		facilityRepo,
		provinceRepo,
		cityRepo,
		subdistrictRepo,
		cloudinaryClient,
	)

	handler := handlers.NewDestinationHandler(usecase)
	r.POST("", handler.CreateDestination)
	r.GET("", handler.GetAllDestinations)
	r.GET("/:id", handler.GetDestinationById)
	r.PUT("/:id", handler.UpdateDestination)
	r.DELETE("/:id", handler.DeleteDestination)
}
