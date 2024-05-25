package routers

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func PingRouter(r *echo.Group) {
	r.GET("/ping", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello World")
	})
}
