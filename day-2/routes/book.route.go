package routes

import (
	"github.com/akbarsahata/alterra-agmc/day-2/controllers"
	"github.com/labstack/echo/v4"
)

func V1books(e *echo.Echo) {
	bookController := controllers.NewBookController()

	bookRoutes := e.Group("/v1/books")

	bookRoutes.GET("", bookController.GetMany)

	bookRoutes.POST("", bookController.CreateOne)

	bookRoutes.GET("/:id", bookController.GetOneByID)

	bookRoutes.PUT("/:id", bookController.UpdateOneByID)

	bookRoutes.DELETE("/:id", bookController.DeleteOneByID)
}
