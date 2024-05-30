package middlewares

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func RoleMiddleware(roles ...string) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			userRole, ok := c.Get("role").(string)
			if !ok {
				return echo.NewHTTPError(http.StatusInternalServerError, "Gagal mendapatkan role")
			}

			for _, role := range roles {
				if userRole == role {
					return next(c)
				}
			}

			return echo.NewHTTPError(http.StatusForbidden, "Anda tidak memiliki izin untuk mengakses")
		}
	}
}
