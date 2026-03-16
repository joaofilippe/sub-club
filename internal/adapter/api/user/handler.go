package user

import (
	"net/http"

	"github.com/joaofilippe/subclub/internal/adapter/api/response"
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
		return response.Error(c, http.StatusBadRequest, "Invalid request payload")
	}

	// Basic validation (can be extracted later)
	if input.Email == "" || input.Type == "" || input.Role == "" {
		return response.Error(c, http.StatusBadRequest, "Missing required fields")
	}

	id, err := h.createUserUseCase.Execute(c.Request().Context(), input)
	if err != nil {
		return response.Error(c, http.StatusInternalServerError, "Failed to create user")
	}

	return response.Success(c, http.StatusCreated, "User created successfully", map[string]string{"id": id})
}
