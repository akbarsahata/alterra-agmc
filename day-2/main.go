package main

import (
	"net/http"
	"os"

	"github.com/akbarsahata/alterra-agmc/day-2/config"
	"github.com/akbarsahata/alterra-agmc/day-2/lib"
	"github.com/akbarsahata/alterra-agmc/day-2/routes"
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

	e.Logger.Fatal(e.Start(":3001"))
}
