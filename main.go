package main

import (
	"flag"
	"fmt"

	"capstone/config"
	"capstone/middlewares"
	"capstone/routers"
	"capstone/seeds"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	config.LoadConfig()
	config.LoadDb()
	e := echo.New()
	e.Use(middleware.RateLimiter(middleware.NewRateLimiterMemoryStore(10)))
	middlewares.LogMiddleware(e)

	seed := flag.Bool("seed", false, "Run database seeder")
	flag.Parse()

	if *seed {
		seeds.RunSeeder()
	}

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins:     []string{"http://localhost:5173"},
		AllowMethods:     []string{echo.GET, echo.POST, echo.PUT, echo.DELETE},
		AllowHeaders:     []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept, echo.HeaderAuthorization},
		AllowCredentials: true,
	}))

	e.Static("static", "static")
	e.File("/docs", "./static/index.html")
	e.GET("/docs/swagger.yaml", func(c echo.Context) error {
		return c.File("./docs/swagger.yaml")
	})

	routers.SetupRouter(e)
	e.Logger.Fatal(e.Start(fmt.Sprintf(":%s", config.ENV.PORT)))
}
