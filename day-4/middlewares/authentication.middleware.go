package middlewares

import (
	"net/http"
	"strconv"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
)

func AuthenticationMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		}

		user := c.Get("user").(*jwt.Token)
		claims := user.Claims.(jwt.MapClaims)
		userID, _ := strconv.ParseInt(claims["id"].(string), 10, 32)

		if id != int(userID) {
			return echo.NewHTTPError(http.StatusForbidden, "you are not allowed to do this action")
		}

		return next(c)
	}
}
