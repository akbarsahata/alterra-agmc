package routes

import (
	"github.com/akbarsahata/alterra-agmc/day-3/controllers"
	"github.com/akbarsahata/alterra-agmc/day-3/middlewares"
	"github.com/labstack/echo/v4"
)

func V1books(e *echo.Echo) {
	bookController := controllers.NewBookController()

	v1bookRoutes := e.Group("/v1/books")

	v1bookRoutes.GET("", bookController.GetMany)

	v1bookRoutes.GET("/:id", bookController.GetOneByID)

	v1bookRoutes.POST("", bookController.CreateOne, middlewares.AuthorizationMiddleware())

	v1bookRoutes.PUT("/:id", bookController.UpdateOneByID, middlewares.AuthorizationMiddleware())

	v1bookRoutes.DELETE("/:id", bookController.DeleteOneByID, middlewares.AuthorizationMiddleware())
}
