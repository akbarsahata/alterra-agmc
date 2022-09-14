package controllers

import (
	"net/http"
	"strconv"

	"github.com/akbarsahata/alterra-agmc/day-2/lib"
	"github.com/akbarsahata/alterra-agmc/day-2/lib/generator"
	"github.com/akbarsahata/alterra-agmc/day-2/models"
	"github.com/labstack/echo/v4"
)

type BookController interface {
	GetMany(c echo.Context) error
	GetOneByID(c echo.Context) error
	CreateOne(c echo.Context) error
	UpdateOneByID(c echo.Context) error
	DeleteOneByID(c echo.Context) error
}

type controller struct{}

func (ctrl *controller) GetMany(c echo.Context) error {
	data := generator.GenerateBooks()

	response := map[string]interface{}{
		"code": http.StatusOK,
		"data": data,
	}

	return c.JSON(http.StatusOK, response)
}

func (ctrl *controller) GetOneByID(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	data := generator.GenerateBook()

	data.ID = id

	response := map[string]interface{}{
		"code": http.StatusOK,
		"data": data,
	}

	return c.JSON(http.StatusOK, response)
}
func (ctrl *controller) CreateOne(c echo.Context) error {
	data := models.BookModel{}
	c.Bind(&data)

	if err := lib.Validate.Struct(&data); err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}

	data.ID = 1

	response := map[string]interface{}{
		"code": http.StatusOK,
		"data": data,
	}

	return c.JSON(http.StatusOK, response)
}
func (ctrl *controller) UpdateOneByID(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	data := models.BookModel{}
	c.Bind(&data)

	if err := lib.Validate.Struct(&data); err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}

	if err := lib.Validate.Var(id, "gt=0"); err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}

	data.ID = id

	response := map[string]interface{}{
		"code": http.StatusOK,
		"data": data,
	}

	return c.JSON(http.StatusOK, response)
}
func (ctrl *controller) DeleteOneByID(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	if err := lib.Validate.Var(id, "gt=0"); err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}

	return c.String(http.StatusNoContent, "")
}

func NewBookController() BookController {
	ctrl := new(controller)

	return ctrl
}
