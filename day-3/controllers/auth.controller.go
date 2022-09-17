package controllers

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/akbarsahata/alterra-agmc/day-3/lib/auth"
	"github.com/akbarsahata/alterra-agmc/day-3/models"
	"github.com/akbarsahata/alterra-agmc/day-3/models/dao"
	"github.com/akbarsahata/alterra-agmc/day-3/models/dto"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type AuthController interface {
	Login(c echo.Context) error
}

type authController struct {
	validate *validator.Validate
	userDAO  dao.UserDAO
}

func NewAuthController(validate *validator.Validate, userDAO dao.UserDAO) AuthController {
	return &authController{validate, userDAO}
}

func (ac *authController) Login(c echo.Context) error {
	data := dto.AuthBodyDTO{}

	if err := c.Bind(&data); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if err := ac.validate.Struct(&data); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	user, err := ac.userDAO.GetOne(&models.User{Email: data.Email})
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		return echo.NewHTTPError(http.StatusNotFound, err.Error())
	} else if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	correctPass, err := auth.ComparePassword(user.Password, data.Password)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	if !correctPass {
		return echo.NewHTTPError(http.StatusUnauthorized, "incorrect email and password")
	}

	token, err := auth.GetToken(strconv.FormatUint(uint64(user.ID), 10))
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	response := map[string]string{
		"token": token,
	}

	return c.JSON(http.StatusOK, response)
}
