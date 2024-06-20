package mobile

import (
	"capstone/config"
	"capstone/externals/cloudinary"
	"capstone/externals/openai"
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

	repositoryRoute := repositories.NewRouteRepository(config.DB)
	cityRepo := repositories.NewCityRepository(config.DB)
	routeRepo := repositories.NewRouteDetailRepository(config.DB)
	destinationRepo := repositories.NewDestinationRepository(config.DB)
	openAIClient := openai.NewOpenAIClient(config.ENV.OPENAI_API_KEY)

	usecaseRoute := usecases.NewRouteUsecase(cityRepo, destinationRepo, repositoryRoute, routeRepo, openAIClient)
	handlerRoute := handlers.NewRouteHandler(usecaseRoute)

	r.GET("", handler.GetDetailUser)
	r.PUT("", handler.InsertUserDetail)
	r.PUT("/change-password", handler.ChangePassword)
	r.GET("/routes", handlerRoute.FindAllByCurrentUser)
	r.GET("/routes/:id", handlerRoute.FindRouteByIdCurrentUser)
	r.DELETE("/routes/:id", handlerRoute.DeleteRoute)
}
