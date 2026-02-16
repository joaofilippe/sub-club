package middleware

import (
	"log/slog"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

// ConfigureLogger returns the Echo middleware configuration for structured logging
func ConfigureLogger() echo.MiddlewareFunc {
	return middleware.RequestLoggerWithConfig(middleware.RequestLoggerConfig{
		LogStatus: true,
		LogURI:    true,
		LogMethod: true,
		LogValuesFunc: func(c echo.Context, v middleware.RequestLoggerValues) error {
			slog.Info("request",
				"method", v.Method,
				"uri", v.URI,
				"status", v.Status,
			)
			return nil
		},
	})
}
