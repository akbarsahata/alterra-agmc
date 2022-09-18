package main

import (
	"net/http"
	"os"
	"sync"

	"github.com/akbarsahata/alterra-agmc/day-4/config"
	"github.com/akbarsahata/alterra-agmc/day-4/lib"
	"github.com/akbarsahata/alterra-agmc/day-4/routes"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

var once sync.Once

func init() {
	once.Do(func() {
		godotenv.Load()

		config.InitEnv()
		config.InitDB()

		lib.InitValidator()

		if os.Getenv("ENV") != "production" {
			config.InitMigration()
		}
	})
}

func main() {
	e := echo.New()
	db := config.GetDB()

	e.Use(middleware.CORS())
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "ok")
	})

	routes.Auths(e, db)

	routes.V1books(e)

	routes.V1users(e, db)

	e.Logger.Fatal(e.Start(":3001"))
}
