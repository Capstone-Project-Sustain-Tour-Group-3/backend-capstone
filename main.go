package main

import (
	"flag"
	"fmt"
	"log"
	"time"

	"capstone/config"
	"capstone/middlewares"
	"capstone/routers"
	"capstone/seeds"

	"github.com/devfeel/mapper"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	config.LoadConfig()

	config.LoadDb()
	sqlDB, err := config.DB.DB()
	if err != nil {
		log.Fatal(err)
	}
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(45)
	sqlDB.SetConnMaxIdleTime(30 * time.Minute)
	defer sqlDB.Close()

	mapper.SetEnabledJsonTag(false)

	e := echo.New()
	e.Use(middleware.RateLimiter(middleware.NewRateLimiterMemoryStore(10)))
	middlewares.LogMiddleware(e)

	seed := flag.Bool("seed", false, "Run database seeder")
	flag.Parse()

	if *seed {
		seeds.RunSeeder()
	}

	e.Use(middleware.CORS())

	e.Static("static", "static")
	e.File("/docs", "./static/index.html")
	e.GET("/docs/swagger.yaml", func(c echo.Context) error {
		return c.File("./docs/swagger.yaml")
	})

	routers.SetupRouter(e)
	e.Logger.Fatal(e.Start(fmt.Sprintf(":%s", config.ENV.PORT)))
}
