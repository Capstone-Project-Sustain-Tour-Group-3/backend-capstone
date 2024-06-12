package routers

import (
	admin "capstone/routers/admin"
	mobile "capstone/routers/mobile"

	"github.com/labstack/echo/v4"
)

func SetupRouter(e *echo.Echo) {
	v1Mobile := e.Group("/v1/mobile")
	{
		auth := v1Mobile.Group("/auth")
		AuthUserRouter(auth)

		destination := v1Mobile.Group("/destinations")
		mobile.DestinationRouter(destination)

		chatbot := v1Mobile.Group("/chatbot")
		mobile.ChatbotRouter(chatbot)

		profile := v1Mobile.Group("/profile")
		mobile.ProfileRouter(profile)

		recommendation := v1Mobile.Group("/personalization")
		mobile.PersonalizationRoute(recommendation)

		homepage := v1Mobile.Group("/home")
		mobile.HomepageRouter(homepage)
	}

	v1Admin := e.Group("/v1/admin")
	{
		auth := v1Admin.Group("/auth")
		AuthAdminRouter(auth)

		user := v1Admin.Group("/users")
		admin.UserRouter(user)

		destination := v1Admin.Group("/destinations")
		admin.DestinationRouter(destination)

		route := v1Admin.Group("/routes")
		admin.RouteRouter(route)

		admins := v1Admin.Group("/admins")
		admin.ManageAdminRouter(admins)

		destinationMedia := v1Admin.Group("/destination-media")
		admin.DestinationMediaRouter(destinationMedia)
	}
}
