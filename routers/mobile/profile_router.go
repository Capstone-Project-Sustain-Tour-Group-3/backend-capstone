package mobile

import (
	"capstone/config"
	"capstone/externals/cloudinary"
	"capstone/handlers"
	"capstone/helpers"
	"capstone/middlewares"
	"capstone/repositories"
	"capstone/usecases"

	"github.com/labstack/echo/v4"
)

func ProfileRouter(r *echo.Group) {
	r.Use(middlewares.JWTMiddleware)
	repository := repositories.NewUserRepository(config.DB)
	cloudinaryClient := cloudinary.NewCloudinaryClient(config.ENV.CLOUDINARY_URL)
	passwordHelper := helpers.NewPasswordHelper()
	iVal := helpers.NewValidationHelper()
	usecase := usecases.NewProfileUsecase(repository, cloudinaryClient, passwordHelper, iVal)
	handler := handlers.NewProfileHandler(usecase)
	r.GET("", handler.GetDetailUser)
	r.PUT("", handler.InsertUserDetail)
	r.PUT("/change-password", handler.ChangePassword)
}
