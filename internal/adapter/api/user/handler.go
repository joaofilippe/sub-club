package userhandler

import (
	"net/http"

	"github.com/joaofilippe/subclub/internal/adapter/api/common"
	"github.com/joaofilippe/subclub/internal/domain/user"
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
		return common.Error(c, http.StatusBadRequest, "Invalid request payload")
	}

	// Basic validation (can be extracted later)
	if input.Email == "" || input.Type == "" || input.Role == "" {
		return common.Error(c, http.StatusBadRequest, "Missing required fields")
	}

	id, err := h.createUserUseCase.Execute(c.Request().Context(), input)
	if err != nil {
		return common.Error(c, http.StatusInternalServerError, "Failed to create user")
	}

	return common.Success(c, http.StatusCreated, "User created successfully", map[string]string{"id": id})
}
