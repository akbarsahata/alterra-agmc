package controllers_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/akbarsahata/alterra-agmc/day-4/controllers"
	"github.com/akbarsahata/alterra-agmc/day-4/lib"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"

	"github.com/joho/godotenv"
)

func TestGetBooks(t *testing.T) {
	godotenv.Load()

	lib.InitValidator()

	e := echo.New()

	req := httptest.NewRequest(http.MethodGet, "/v1/books", nil)
	rec := httptest.NewRecorder()

	c := e.NewContext(req, rec)
	controller := controllers.NewBookController()

	if assert.NoError(t, controller.GetMany(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
	}
}

func TestGetBookByID(t *testing.T) {
	godotenv.Load()

	lib.InitValidator()

	e := echo.New()

	req := httptest.NewRequest(http.MethodGet, "/v1/books/1", nil)
	rec := httptest.NewRecorder()

	c := e.NewContext(req, rec)
	controller := controllers.NewBookController()

	if assert.NoError(t, controller.GetOneByID(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
	}
}
