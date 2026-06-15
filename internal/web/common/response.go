package common

import (
	"github.com/labstack/echo/v4"
)

// ErrorDetail holds the detail of an error response.
type ErrorDetail struct {
	Message string `json:"message"`
}

// Response represents a standard API response format.
type Response struct {
	Message string       `json:"message"`
	Data    any          `json:"data,omitempty"`
	Error   *ErrorDetail `json:"error,omitempty"`
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
		Error:   &ErrorDetail{Message: message},
	})
}
