package routers

import (
	"capstone/config"
	"capstone/handlers"
	"capstone/repositories"
	"capstone/usecases"
	"github.com/labstack/echo/v4"
)

func AuthUserRouter(r *echo.Group) {
	userRepo := repositories.NewUserRepository(config.DB)
	cacheRepo := repositories.NewCacheRepository()
	usecase := usecases.NewUserUsecase(userRepo, cacheRepo)
	handler := handlers.NewUserHandler(usecase)
	r.POST("/register", handler.Register)
	r.POST("/resend-otp", handler.ResendOTP)
	r.POST("/verify", handler.VerifyEmail)
}

func AuthAdminRouter(r *echo.Group) {
}
