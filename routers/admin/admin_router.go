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

func ManageAdminRouter(r *echo.Group) {
	r.Use(middlewares.JWTMiddleware)
	r.Use(middlewares.RoleMiddleware("superadmin"))

	repository := repositories.NewAdminRepository(config.DB)
	passwordHelper := helpers.NewPasswordHelper()
	tokenHelper := helpers.NewTokenHelper()
	cloudinaryClient := cloudinary.NewCloudinaryClient(config.ENV.CLOUDINARY_URL)

	usecase := usecases.NewAdminUsecase(repository, cloudinaryClient, passwordHelper, tokenHelper)

	handler := handlers.NewAdminHandler(usecase)

	r.GET("", handler.GetAllAdmins)
	r.POST("", handler.CreateAdmin)
	r.GET("/:id", handler.GetAdminDetail)
	r.PUT("/:id", handler.UpdateAdmin)
	r.DELETE("/:id", handler.DeleteAdmin)
}
