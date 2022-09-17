package routes

import (
	"github.com/akbarsahata/alterra-agmc/day-3/controllers"
	"github.com/akbarsahata/alterra-agmc/day-3/lib"
	"github.com/akbarsahata/alterra-agmc/day-3/models/dao"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func Auths(e *echo.Echo, db *gorm.DB) {
	userDAO := dao.NewUserDAO(db)
	authController := controllers.NewAuthController(lib.Validate, userDAO)

	e.POST("/login", authController.Login)
}
