package routes

import (
	"github.com/akbarsahata/alterra-agmc/day-1/controllers"
	"github.com/labstack/echo/v4"
)

func New() *echo.Echo {
	e := echo.New()

	userRoute := e.Group("/users")

	userRoute.GET("", controllers.GetAllUsers)

	userRoute.POST("", controllers.CreateUser)

	return e
}
