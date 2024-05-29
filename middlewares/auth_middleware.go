package middlewares

import (
	"net/http"
	"strings"

	"capstone/helpers"

	"github.com/labstack/echo/v4"
)

func JWTMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		authHeader := c.Request().Header.Get("Authorization")
		if authHeader == "" {
			return echo.NewHTTPError(http.StatusUnauthorized, "Token tidak terisi")
		}

		if !strings.HasPrefix(authHeader, "Bearer ") {
			return echo.NewHTTPError(http.StatusBadRequest, "Format token tidak valid, gunakan bearer token")
		}
		tokenStr := strings.TrimPrefix(authHeader, "Bearer ")

		claims, err := helpers.ParseJWT(tokenStr)
		id := &claims.Id
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, err)
		}

		c.Set("userId", id)
		return next(c)
	}
}
