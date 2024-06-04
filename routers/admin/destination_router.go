package admin

import (
	"capstone/config"
	"capstone/externals/cloudinary"
	"capstone/handlers"
	"capstone/repositories"
	"capstone/usecases"

	"github.com/labstack/echo/v4"
)

func DestinationRouter(r *echo.Group) {
	// r.Use(middlewares.JWTMiddleware)

	destinationRepo := repositories.NewDestinationRepository(config.DB)
	destinationFacilityRepo := repositories.NewDestinationFacilityRepository(config.DB)
	destinationMediaRepo := repositories.NewDestinationMediaRepository(config.DB)
	destinationAddressRepo := repositories.NewDestinationAddressRepository(config.DB)
	categoryRepo := repositories.NewCategoryRepository(config.DB)
	cloudinaryClient := cloudinary.NewCloudinaryClient(config.ENV.CLOUDINARY_URL)

	usecase := usecases.NewDestinationUsecase(
		destinationRepo,
		destinationFacilityRepo,
		destinationMediaRepo,
		destinationAddressRepo,
		categoryRepo,
		cloudinaryClient,
	)

	handler := handlers.NewDestinationHandler(usecase)
	r.POST("", handler.CreateDestination)
}
