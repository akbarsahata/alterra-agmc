package main

import (
	"net/http"
	"os"

	"github.com/akbarsahata/alterra-agmc/day-3/config"
	"github.com/akbarsahata/alterra-agmc/day-3/lib"
	"github.com/akbarsahata/alterra-agmc/day-3/routes"
	"github.com/labstack/echo/v4"
)

func init() {
	lib.InitValidator()

	config.InitDB()

	if os.Getenv("ENV") != "production" {
		config.InitMigration()
	}
}

func main() {
	e := echo.New()

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "ok")
	})

	routes.V1books(e)

	routes.V1users(e)

	e.Logger.Fatal(e.Start(":3001"))
}
