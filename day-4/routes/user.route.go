package routes

import (
	"github.com/akbarsahata/alterra-agmc/day-4/controllers"
	"github.com/akbarsahata/alterra-agmc/day-4/lib"
	"github.com/akbarsahata/alterra-agmc/day-4/middlewares"
	"github.com/akbarsahata/alterra-agmc/day-4/models/dao"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func V1users(e *echo.Echo, db *gorm.DB) {
	userDAO := dao.NewUserDAO(db)
	userController := controllers.NewUserController(lib.Validate, userDAO)

	v1usersRoutes := e.Group("/v1/users")

	v1usersRoutes.GET("", userController.GetMany, middlewares.AuthorizationMiddleware())

	v1usersRoutes.POST("", userController.CreateOne)

	v1usersRoutes.GET("/:id", userController.GetOneByID, middlewares.AuthorizationMiddleware())

	v1usersRoutes.PUT("/:id", userController.UpdateOneByID, middlewares.AuthorizationMiddleware(), middlewares.AuthenticationMiddleware)

	v1usersRoutes.DELETE("/:id", userController.DeleteOneByID, middlewares.AuthorizationMiddleware(), middlewares.AuthenticationMiddleware)
}
