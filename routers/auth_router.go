package routers

import (
	"capstone/config"
	"capstone/handlers"
	"capstone/middlewares"
	"capstone/repositories"
	"capstone/usecases"

	"github.com/labstack/echo/v4"
)

func AuthUserRouter(r *echo.Group) {
	userRepo := repositories.NewAuthRepository(config.DB)
	cacheRepo := repositories.NewCacheRepository()
	usecase := usecases.NewAuthUsecase(userRepo, cacheRepo)
	handler := handlers.NewAuthHandler(usecase)

	r.POST("/register", handler.Register)
	r.POST("/resend-otp", handler.ResendOTP)
	r.POST("/verify", handler.VerifyEmail)
	r.POST("/login", handler.Login)
	r.PUT("/forgot-password", handler.ForgotPassword)
	r.DELETE("/logout", handler.Logout)
	r.POST("/token", handler.GetNewAccessToken)

	// Routes that require JWT middleware
	r.Use(middlewares.JWTMiddleware)
	r.GET("/pong", handler.Pong)
}

func AuthAdminRouter(r *echo.Group) {
	repository := repositories.NewAdminRepository(config.DB)
	usecase := usecases.NewAdminUsecase(repository)
	handler := handlers.NewAdminHandler(usecase)

	r.POST("/login", handler.Login)
	r.DELETE("/logout", handler.Logout)
	r.POST("/token", handler.GetNewAccessToken)
}
