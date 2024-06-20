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

func DestinationMediaRouter(r *echo.Group) {
	r.Use(middlewares.JWTMiddleware)
	r.Use(middlewares.RoleMiddleware("admin", "superadmin"))

	destinationRepo := repositories.NewDestinationRepository(config.DB)
	destinationMediaRepo := repositories.NewDestinationMediaRepository(config.DB)
	cloudinaryClient := cloudinary.NewCloudinaryClient(config.ENV.CLOUDINARY_URL)

	usecase := usecases.NewDestinationMediaUsecase(destinationMediaRepo, destinationRepo, cloudinaryClient)

	handler := handlers.NewDestinationMediaHandler(usecase)

	r.POST("", handler.CreateDestinationMedia)
	r.POST("/image", handler.UploadImageFile)
	r.GET("", handler.AllDestinationMedia)
	r.GET("/:id", handler.DetailDestinationMedia)
	r.PUT("/:id", handler.UpdateDestinationMedia)
	r.PUT("/:id/image", handler.UpdateImageFile)
	r.DELETE("/:id", handler.DeleteDestinationMedia)
}
