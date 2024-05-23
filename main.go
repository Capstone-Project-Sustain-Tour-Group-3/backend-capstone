package main

import (
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

	routers.SetupRouter(e)
	e.Logger.Fatal(e.Start(":1323"))
}
