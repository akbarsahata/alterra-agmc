package controllers

import (
	"net/http"
	"strconv"

	"github.com/akbarsahata/alterra-agmc/day-2/lib/database"
	"github.com/akbarsahata/alterra-agmc/day-2/models/dto"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

type UserController interface {
	GetMany(c echo.Context) error
	GetOneByID(c echo.Context) error
	CreateOne(c echo.Context) error
	UpdateOneByID(c echo.Context) error
	DeleteOneByID(c echo.Context) error
}

type userController struct {
	validate *validator.Validate
}

func (uc *userController) GetMany(c echo.Context) error {
	users, err := database.GetUsers()
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	response := []dto.UserResponseDTO{}

	for _, u := range users {
		response = append(response, u.ToResponseDTO())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"code": http.StatusOK,
		"data": response,
	})
}

func (uc *userController) GetOneByID(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	data, err := database.GetUserByID(uint(id))
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"code": http.StatusOK,
		"data": data.ToResponseDTO(),
	})
}

func (uc *userController) CreateOne(c echo.Context) error {
	data := dto.CreateUserBodyDTO{}

	if err := c.Bind(&data); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if err := uc.validate.Struct(&data); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if user, err := database.CreateUser(data); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	} else {
		return c.JSON(http.StatusOK, map[string]interface{}{
			"code": http.StatusOK,
			"data": user.ToResponseDTO(),
		})
	}
}

func (uc *userController) UpdateOneByID(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	data := dto.UpdateUserBodyDTO{}

	if err := c.Bind(&data); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if err := uc.validate.Struct(&data); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if user, err := database.UpdateUser(uint(id), data); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	} else {
		return c.JSON(http.StatusOK, map[string]interface{}{
			"code": http.StatusOK,
			"data": user.ToResponseDTO(),
		})
	}
}

func (uc *userController) DeleteOneByID(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if err := database.DeleteUser(uint(id)); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.String(http.StatusNoContent, "")
}

func NewUserController(validate *validator.Validate) UserController {
	ctrl := userController{validate: validate}

	return &ctrl
}
