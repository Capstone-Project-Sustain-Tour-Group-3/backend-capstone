package admin

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

func UserRouter(r *echo.Group) {
	r.Use(middlewares.JWTMiddleware)
	r.Use(middlewares.RoleMiddleware("admin", "superadmin"))

	repository := repositories.NewUserRepository(config.DB)
	cloudinaryClient := cloudinary.NewCloudinaryClient(config.ENV.CLOUDINARY_URL)
	passwordHelper := helpers.NewPasswordHelper()
	iVal := helpers.NewValidationHelper()
	usecase := usecases.NewUserUsecase(repository, cloudinaryClient, passwordHelper, iVal)
	handler := handlers.NewUserHandler(usecase)
	r.Use(middlewares.JWTMiddleware)
	r.Use(middlewares.RoleMiddleware("admin"))

	r.GET("", handler.FindAll)
	r.POST("", handler.Create)
	r.GET("/:id", handler.FindById)
	r.PUT("/:id", handler.Update)
	r.DELETE("/:id", handler.Delete)
}
