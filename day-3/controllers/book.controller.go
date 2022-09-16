package controllers

import (
	"net/http"
	"strconv"

	"github.com/akbarsahata/alterra-agmc/day-3/lib"
	"github.com/akbarsahata/alterra-agmc/day-3/lib/generator"
	"github.com/akbarsahata/alterra-agmc/day-3/models"
	"github.com/labstack/echo/v4"
)

type BookController interface {
	GetMany(c echo.Context) error
	GetOneByID(c echo.Context) error
	CreateOne(c echo.Context) error
	UpdateOneByID(c echo.Context) error
	DeleteOneByID(c echo.Context) error
}

type bookController struct{}

func (bc *bookController) GetMany(c echo.Context) error {
	data := generator.GenerateBooks()

	response := map[string]interface{}{
		"code": http.StatusOK,
		"data": data,
	}

	return c.JSON(http.StatusOK, response)
}

func (bc *bookController) GetOneByID(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	data := generator.GenerateBook()

	data.ID = id

	response := map[string]interface{}{
		"code": http.StatusOK,
		"data": data,
	}

	return c.JSON(http.StatusOK, response)
}
func (bc *bookController) CreateOne(c echo.Context) error {
	data := models.BookModel{}
	c.Bind(&data)

	if err := lib.Validate.Struct(&data); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	data.ID = 1

	response := map[string]interface{}{
		"code": http.StatusOK,
		"data": data,
	}

	return c.JSON(http.StatusOK, response)
}
func (bc *bookController) UpdateOneByID(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	data := models.BookModel{}
	c.Bind(&data)

	if err := lib.Validate.Struct(&data); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if err := lib.Validate.Var(id, "gt=0"); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	data.ID = id

	response := map[string]interface{}{
		"code": http.StatusOK,
		"data": data,
	}

	return c.JSON(http.StatusOK, response)
}
func (bc *bookController) DeleteOneByID(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	if err := lib.Validate.Var(id, "gt=0"); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.String(http.StatusNoContent, "")
}

func NewBookController() BookController {
	ctrl := new(bookController)

	return ctrl
}
