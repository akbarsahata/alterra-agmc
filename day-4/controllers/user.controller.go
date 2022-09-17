package controllers

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/akbarsahata/alterra-agmc/day-4/models"
	"github.com/akbarsahata/alterra-agmc/day-4/models/dao"
	"github.com/akbarsahata/alterra-agmc/day-4/models/dto"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
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
	userDAO  dao.UserDAO
}

func NewUserController(validate *validator.Validate, userDAO dao.UserDAO) UserController {
	return &userController{validate: validate, userDAO: userDAO}
}

func (uc *userController) GetMany(c echo.Context) error {
	users, err := uc.userDAO.GetMany()
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

	data, err := uc.userDAO.GetOneByID(uint(id))
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

	existingUser, err := uc.userDAO.GetOne(&models.User{Email: data.Email})
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	} else if existingUser != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("%s already been used", data.Email))
	}

	if user, err := uc.userDAO.CreateOne(data); err != nil {
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

	if user, err := uc.userDAO.UpdateOne(uint(id), data); err != nil {
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

	if err := uc.userDAO.DeleteOne(uint(id)); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.String(http.StatusNoContent, "")
}
