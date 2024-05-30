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
		DestinationRouter(destination)

		chatbot := v1Mobile.Group("/chatbot")
		mobile.ChatbotRouter(chatbot)
	}

	v1Admin := e.Group("/v1/admin")
	{
		auth := v1Admin.Group("/auth")
		AuthAdminRouter(auth)

		user := v1Admin.Group("/users")
		admin.UserRouter(user)
	}
}
