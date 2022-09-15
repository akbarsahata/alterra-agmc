package routes

import (
	"github.com/akbarsahata/alterra-agmc/day-2/controllers"
	"github.com/akbarsahata/alterra-agmc/day-2/lib"
	"github.com/labstack/echo/v4"
)

func V1users(e *echo.Echo) {
	userController := controllers.NewUserController(lib.Validate)

	v1usersRoutes := e.Group("/v1/users")

	v1usersRoutes.GET("", userController.GetMany)

	v1usersRoutes.POST("", userController.CreateOne)

	v1usersRoutes.GET("/:id", userController.GetOneByID)

	v1usersRoutes.PUT("/:id", userController.UpdateOneByID)

	v1usersRoutes.DELETE("/:id", userController.DeleteOneByID)
}
