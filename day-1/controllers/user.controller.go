package controllers

import (
	"net/http"

	"github.com/akbarsahata/alterra-agmc/day-1/lib/database"
	"github.com/akbarsahata/alterra-agmc/day-1/lib/dto"
	"github.com/labstack/echo/v4"
)

func GetAllUsers(c echo.Context) error {
	data := []dto.UserDTO{}

	if users, err := database.GetUsers(); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	} else {
		for _, u := range users {
			data = append(data, u.DTO())
		}

		return c.JSON(http.StatusOK, map[string]interface{}{
			"code": http.StatusOK,
			"data": data,
		})
	}
}

func CreateUser(c echo.Context) error {
	data := dto.UserDTO{}

	c.Bind(&data)

	if user, err := database.CreateUser(data); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	} else {
		return c.JSON(http.StatusOK, map[string]interface{}{
			"code": http.StatusOK,
			"data": user.DTO(),
		})
	}
}
