package controllers_test

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/akbarsahata/alterra-agmc/day-4/controllers"
	"github.com/akbarsahata/alterra-agmc/day-4/lib"
	mocks "github.com/akbarsahata/alterra-agmc/day-4/mocks/models/dao"
	"github.com/akbarsahata/alterra-agmc/day-4/models"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestGetUsers_500_UserDAOFailedToGetMany(t *testing.T) {
	setupTest()

	userDAOmock := mocks.NewUserDAO(t)

	userDAOmock.On("GetMany").Return(
		func() []models.User {
			return nil
		},
		func() error {
			return errors.New("failed")
		},
	)

	e := echo.New()

	req := httptest.NewRequest(http.MethodGet, "/v1/users", nil)
	rec := httptest.NewRecorder()

	c := e.NewContext(req, rec)
	controller := controllers.NewUserController(lib.Validate, userDAOmock)

	err := controller.GetMany(c)

	if assert.Error(t, err) {
		assert.NotNil(t, err)

		httpError, ok := err.(*echo.HTTPError)
		if ok {
			assert.Equal(t, http.StatusInternalServerError, httpError.Code)
		}
	}
}

func TestGetUsers(t *testing.T) {
	setupTest()

	userModel := models.User{
		Name:     "test",
		Email:    "test@test.com",
		Password: "test",
	}
	userModel.ID = 1
	userModel.CreatedAt = time.Now()
	userModel.UpdatedAt = time.Now()

	getManyReturn := []models.User{}
	getManyReturn = append(getManyReturn, userModel)

	userDAOmock := mocks.NewUserDAO(t)

	userDAOmock.On("GetMany").Return(
		func() []models.User {
			return getManyReturn
		},
		func() error {
			return nil
		},
	)

	e := echo.New()

	req := httptest.NewRequest(http.MethodGet, "/v1/users", nil)
	rec := httptest.NewRecorder()

	c := e.NewContext(req, rec)
	controller := controllers.NewUserController(lib.Validate, userDAOmock)

	if assert.NoError(t, controller.GetMany(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		fmt.Println(rec.Body.String())
	}
}

func TestGetUserByID_400_InvalidID(t *testing.T) {
	setupTest()

	userDAOmock := mocks.NewUserDAO(t)

	e := echo.New()

	req := httptest.NewRequest(http.MethodGet, "/v1/users", nil)
	rec := httptest.NewRecorder()

	c := e.NewContext(req, rec)
	c.SetPath("/:id")
	c.SetParamNames("id")
	c.SetParamValues("ab")
	controller := controllers.NewUserController(lib.Validate, userDAOmock)

	err := controller.GetOneByID(c)

	if assert.Error(t, err) {
		assert.NotNil(t, err)

		httpError, ok := err.(*echo.HTTPError)
		if ok {
			assert.Equal(t, http.StatusBadRequest, httpError.Code)
		}
	}
}

func TestGetUserByID_500_UserDAOFailedToGetOneByID(t *testing.T) {
	setupTest()

	userDAOmock := mocks.NewUserDAO(t)

	userDAOmock.On("GetOneByID", mock.Anything).Return(
		func(userID uint) *models.User {
			return nil
		},
		func(userID uint) error {
			return errors.New("failed")
		},
	)

	e := echo.New()

	req := httptest.NewRequest(http.MethodGet, "/v1/users", nil)
	rec := httptest.NewRecorder()

	c := e.NewContext(req, rec)
	c.SetPath("/:id")
	c.SetParamNames("id")
	c.SetParamValues("1")
	controller := controllers.NewUserController(lib.Validate, userDAOmock)

	err := controller.GetOneByID(c)

	if assert.Error(t, err) {
		assert.NotNil(t, err)

		httpError, ok := err.(*echo.HTTPError)
		if ok {
			assert.Equal(t, http.StatusInternalServerError, httpError.Code)
		}
	}
}

func TestGetUserByID(t *testing.T) {
	setupTest()

	userModel := models.User{
		Name:     "test",
		Email:    "test@test.com",
		Password: "test",
	}
	userModel.ID = 1
	userModel.CreatedAt = time.Now()
	userModel.UpdatedAt = time.Now()

	response := map[string]interface{}{
		"code": http.StatusOK,
		"data": userModel.ToResponseDTO(),
	}

	resPayload, _ := json.Marshal(response)

	userDAOmock := mocks.NewUserDAO(t)

	userDAOmock.On("GetOneByID", mock.Anything).Return(
		func(userID uint) *models.User {
			return &userModel
		},
		func(userID uint) error {
			return nil
		},
	)

	e := echo.New()

	req := httptest.NewRequest(http.MethodGet, "/v1/users", nil)
	rec := httptest.NewRecorder()

	c := e.NewContext(req, rec)
	c.SetPath("/:id")
	c.SetParamNames("id")
	c.SetParamValues("1")
	controller := controllers.NewUserController(lib.Validate, userDAOmock)

	if assert.NoError(t, controller.GetOneByID(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, string(resPayload)+"\n", rec.Body.String())
	}
}

func TestCreateOne_400_InvalidJSON(t *testing.T) {

}

func TestCreateOne_400_InvalidRequestPayload(t *testing.T) {

}

func TestCreateOne_500_UserDAOFailedToGetOne(t *testing.T) {

}

func TestCreateOne_400_EmailAlreadyUsed(t *testing.T) {

}

func TestCreateOne_500_UserDAOFailedToCreateOne(t *testing.T) {

}

func TestCreateOne(t *testing.T) {

}
