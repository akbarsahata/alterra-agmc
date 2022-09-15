package routes

import (
	"github.com/akbarsahata/alterra-agmc/day-2/controllers"
	"github.com/labstack/echo/v4"
)

func V1books(e *echo.Echo) {
	bookController := controllers.NewBookController()

	v1bookRoutes := e.Group("/v1/books")

	v1bookRoutes.GET("", bookController.GetMany)

	v1bookRoutes.POST("", bookController.CreateOne)

	v1bookRoutes.GET("/:id", bookController.GetOneByID)

	v1bookRoutes.PUT("/:id", bookController.UpdateOneByID)

	v1bookRoutes.DELETE("/:id", bookController.DeleteOneByID)
}
