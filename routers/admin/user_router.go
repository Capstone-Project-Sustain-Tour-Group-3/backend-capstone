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

func UserRouter(r *echo.Group) {
	repository := repositories.NewUserRepository(config.DB)
	cloudinaryClient := cloudinary.NewCloudinaryClient(config.ENV.CLOUDINARY_URL)
	usecase := usecases.NewUserUsecase(repository, cloudinaryClient)
	handler := handlers.NewUserHandler(usecase)
	r.Use(middlewares.JWTMiddleware)
	r.Use(middlewares.RoleMiddleware("admin"))

	r.GET("", handler.FindAll)
	r.POST("", handler.Create)
	r.GET("/:id", handler.FindById)
	r.PUT("/:id", handler.Update)
	r.DELETE("/:id", handler.Delete)
}
