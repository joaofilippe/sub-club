package response

import (
	"github.com/labstack/echo/v4"
)

// Response represents a standard API response format.
type Response struct {
	Message string `json:"message"`
	Data    any    `json:"data,omitempty"`
}

// Success is a helper function to return a successful response.
func Success(c echo.Context, statusCode int, message string, data any) error {
	return c.JSON(statusCode, Response{
		Message: message,
		Data:    data,
	})
}

// Error is a helper function to return an error response.
func Error(c echo.Context, statusCode int, message string) error {
	return c.JSON(statusCode, Response{
		Message: message,
	})
}
