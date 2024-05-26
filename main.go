package main

import (
	"fmt"

	"capstone/config"
	"capstone/middlewares"
	"capstone/routers"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	config.LoadConfig()
	config.LoadDb()
	e := echo.New()
	e.Use(middleware.RateLimiter(middleware.NewRateLimiterMemoryStore(10)))
	middlewares.LogMiddleware(e)

	e.Use(middleware.CORS())
	e.Static("static", "static")
	e.File("/docs", "./static/index.html")
	e.GET("/docs/swagger.yaml", func(c echo.Context) error {
		return c.File("./docs/swagger.yaml")
	})

	routers.SetupRouter(e)
	e.Logger.Fatal(e.Start(fmt.Sprintf(":%s", config.ENV.PORT)))
}
