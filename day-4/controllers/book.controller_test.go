package controllers_test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/akbarsahata/alterra-agmc/day-4/controllers"
	"github.com/akbarsahata/alterra-agmc/day-4/lib"
	"github.com/akbarsahata/alterra-agmc/day-4/models"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"

	"github.com/joho/godotenv"
)

func setupTest() {
	godotenv.Load()

	lib.InitValidator()
}

func TestGetBooks(t *testing.T) {
	setupTest()

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
	setupTest()

	e := echo.New()

	req := httptest.NewRequest(http.MethodGet, "/v1/books/1", nil)
	rec := httptest.NewRecorder()

	c := e.NewContext(req, rec)
	controller := controllers.NewBookController()

	if assert.NoError(t, controller.GetOneByID(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
	}
}

func TestCreateBook_400_InvalidPayload(t *testing.T) {
	setupTest()
	e := echo.New()

	req := httptest.NewRequest(http.MethodPost, "/v1/books/", nil)
	rec := httptest.NewRecorder()

	c := e.NewContext(req, rec)
	controller := controllers.NewBookController()

	err := controller.CreateOne(c)

	if assert.Error(t, err) {
		assert.NotNil(t, err)

		httpError, ok := err.(*echo.HTTPError)
		if ok {
			assert.Equal(t, http.StatusBadRequest, httpError.Code)
		}
	}
}

func TestCreateBook(t *testing.T) {
	setupTest()

	reqData := models.BookModel{
		Title:       "Bumi Manusia",
		Author:      "Pramoedya Ananta Toer",
		Description: "Begitukah jadi pribumi?",
	}

	resData := models.BookModel{
		ID:          1,
		Title:       "Bumi Manusia",
		Author:      "Pramoedya Ananta Toer",
		Description: "Begitukah jadi pribumi?",
	}

	response := map[string]interface{}{
		"code": http.StatusOK,
		"data": resData,
	}

	reqPayload, _ := json.Marshal(reqData)
	resPayload, _ := json.Marshal(response)

	e := echo.New()

	req := httptest.NewRequest(http.MethodPost, "/v1/books/", strings.NewReader(string(reqPayload)))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()

	c := e.NewContext(req, rec)
	controller := controllers.NewBookController()

	if assert.NoError(t, controller.CreateOne(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, string(resPayload)+"\n", rec.Body.String())
	}
}

func TestUpdateBook_400_InvalidPayload(t *testing.T) {
	setupTest()

	e := echo.New()

	req := httptest.NewRequest(http.MethodPut, "/v1/books/1", nil)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()

	c := e.NewContext(req, rec)
	controller := controllers.NewBookController()

	err := controller.UpdateOneByID(c)

	if assert.Error(t, err) {
		assert.NotNil(t, err)

		httpError, ok := err.(*echo.HTTPError)
		if ok {
			assert.Equal(t, http.StatusBadRequest, httpError.Code)
		}
	}
}

func TestUpdateBook_400_InvalidID(t *testing.T) {
	setupTest()

	reqData := models.BookModel{
		Title:       "Bumi Manusia",
		Author:      "Pramoedya Ananta Toer",
		Description: "Begitukah jadi pribumi?",
	}

	reqPayload, _ := json.Marshal(reqData)

	e := echo.New()

	req := httptest.NewRequest(http.MethodPut, "/v1/books/", strings.NewReader(string(reqPayload)))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()

	c := e.NewContext(req, rec)
	c.SetPath("/:id")
	c.SetParamNames("id")
	c.SetParamValues("ab")
	controller := controllers.NewBookController()

	err := controller.UpdateOneByID(c)

	if assert.Error(t, err) {
		assert.NotNil(t, err)

		httpError, ok := err.(*echo.HTTPError)
		if ok {
			assert.Equal(t, http.StatusBadRequest, httpError.Code)
		}
	}
}

func TestUpdateBook(t *testing.T) {
	setupTest()

	reqData := models.BookModel{
		Title:       "Bumi Manusia",
		Author:      "Pramoedya Ananta Toer",
		Description: "Begitukah jadi pribumi?",
	}

	resData := models.BookModel{
		ID:          1,
		Title:       "Bumi Manusia",
		Author:      "Pramoedya Ananta Toer",
		Description: "Begitukah jadi pribumi?",
	}

	response := map[string]interface{}{
		"code": http.StatusOK,
		"data": resData,
	}

	reqPayload, _ := json.Marshal(reqData)
	resPayload, _ := json.Marshal(response)

	e := echo.New()

	req := httptest.NewRequest(http.MethodPut, "/v1/books/", strings.NewReader(string(reqPayload)))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()

	c := e.NewContext(req, rec)
	c.SetPath("/:id")
	c.SetParamNames("id")
	c.SetParamValues("1")
	controller := controllers.NewBookController()

	if assert.NoError(t, controller.UpdateOneByID(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, string(resPayload)+"\n", rec.Body.String())
	}
}

func TestDeleteBook_400_InvalidID(t *testing.T) {
	setupTest()

	e := echo.New()

	req := httptest.NewRequest(http.MethodDelete, "/v1/books/", nil)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()

	c := e.NewContext(req, rec)
	c.SetPath("/:id")
	c.SetParamNames("id")
	c.SetParamValues("ab")
	controller := controllers.NewBookController()

	err := controller.DeleteOneByID(c)

	if assert.Error(t, err) {
		assert.NotNil(t, err)

		httpError, ok := err.(*echo.HTTPError)
		if ok {
			assert.Equal(t, http.StatusBadRequest, httpError.Code)
		}
	}
}

func TestDeleteBook(t *testing.T) {
	setupTest()

	e := echo.New()

	req := httptest.NewRequest(http.MethodDelete, "/v1/books/", nil)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()

	c := e.NewContext(req, rec)
	c.SetPath("/:id")
	c.SetParamNames("id")
	c.SetParamValues("1")
	controller := controllers.NewBookController()

	if assert.NoError(t, controller.DeleteOneByID(c)) {
		assert.Equal(t, http.StatusNoContent, rec.Code)
	}
}
