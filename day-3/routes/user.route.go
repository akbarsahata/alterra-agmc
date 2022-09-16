package routes

import (
	"github.com/akbarsahata/alterra-agmc/day-3/config"
	"github.com/akbarsahata/alterra-agmc/day-3/controllers"
	"github.com/akbarsahata/alterra-agmc/day-3/lib"
	"github.com/akbarsahata/alterra-agmc/day-3/models/dao"
	"github.com/labstack/echo/v4"
)

func V1users(e *echo.Echo) {
	userDAO := dao.NewUserDAO(config.DB)
	userController := controllers.NewUserController(lib.Validate, userDAO)

	v1usersRoutes := e.Group("/v1/users")

	v1usersRoutes.GET("", userController.GetMany)

	v1usersRoutes.POST("", userController.CreateOne)

	v1usersRoutes.GET("/:id", userController.GetOneByID)

	v1usersRoutes.PUT("/:id", userController.UpdateOneByID)

	v1usersRoutes.DELETE("/:id", userController.DeleteOneByID)
}
