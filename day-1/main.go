package main

import (
	"net/http"
	"os"

	"github.com/akbarsahata/alterra-agmc/day-1/config"
	"github.com/akbarsahata/alterra-agmc/day-1/routes"
	"github.com/labstack/echo/v4"
)

func init() {
	config.InitDB()

	if os.Getenv("ENV") != "production" {
		config.InitMigration()
	}
}

func main() {
	e := routes.New()

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "ok")
	})

	e.Logger.Fatal(e.Start(":3001"))
}
