package user

import (
	"net/http"

	"github.com/joaofilippe/subclub/internal/application/usecase/user"
	"github.com/labstack/echo/v4"
)

type UserHandler struct {
	createUserUseCase *user.CreateUserUseCase
}

func NewUserHandler(createUserUseCase *user.CreateUserUseCase) *UserHandler {
	return &UserHandler{
		createUserUseCase: createUserUseCase,
	}
}

func (h *UserHandler) Create(c echo.Context) error {
	var input user.CreateUserInput
	if err := c.Bind(&input); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request payload"})
	}

	// Basic validation (can be extracted later)
	if input.Email == "" || input.Type == "" || input.Role == "" {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Missing required fields"})
	}

	id, err := h.createUserUseCase.Execute(c.Request().Context(), input)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to create user"})
	}

	return c.JSON(http.StatusCreated, map[string]string{"id": id})
}
