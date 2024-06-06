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

func ManageAdminRouter(r *echo.Group) {
	r.Use(middlewares.JWTMiddleware)
	r.Use(middlewares.RoleMiddleware("superadmin"))

	repository := repositories.NewAdminRepository(config.DB)
	cloudinaryClient := cloudinary.NewCloudinaryClient(config.ENV.CLOUDINARY_URL)

	usecase := usecases.NewAdminUsecase(repository, cloudinaryClient)

	handler := handlers.NewAdminHandler(usecase)

	r.GET("", handler.GetAllAdmins)
	r.POST("", handler.CreateAdmin)
	r.GET("/:id", handler.GetAdminDetail)
	r.PUT("/:id", handler.UpdateAdmin)
}
