package routers

import (
	"github.com/labstack/echo/v4"
)

func SetupRouter(e *echo.Echo) {
	v1User := e.Group("/v1")
	{
		auth := v1User.Group("/auth")
		AuthUserRouter(auth)
	}

	v1Admin := e.Group("/v1/admin")
	{
		auth := v1Admin.Group("/auth")
		AuthAdminRouter(auth)
	}
}
