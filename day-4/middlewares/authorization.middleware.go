package middlewares

import (
	"github.com/akbarsahata/alterra-agmc/day-4/config"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func AuthorizationMiddleware() echo.MiddlewareFunc {
	return middleware.JWTWithConfig(middleware.JWTConfig{
		SigningKey: []byte(config.Env["JWT_SECRET"]),
	})
}
