package userhandler

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/joaofilippe/subclub/internal/adapter/api/common"
	domain "github.com/joaofilippe/subclub/internal/domain/user"
	"github.com/joaofilippe/subclub/internal/domain/user/model"
)

type UserHandler struct {
	service domain.Service
}

func NewUserHandler(service domain.Service) *UserHandler {
	return &UserHandler{
		service: service,
	}
}

// Create godoc
// @Summary      Create a user
// @Description  Creates a new administrative user
// @Tags         users
// @Accept       json
// @Produce      json
// @Param        user  body      model.CreateUserInput  true  "User input"
// @Success      201   {object}  map[string]interface{}
// @Failure      400   {object}  map[string]string
// @Failure      500   {object}  map[string]string
// @Router       /users [post]
func (h *UserHandler) Create(c echo.Context) error {
	var input model.CreateUserInput
	if err := c.Bind(&input); err != nil {
		return common.Error(c, http.StatusBadRequest, "Invalid request payload")
	}

	if input.Email == "" || input.Type == "" || input.Role == "" {
		return common.Error(c, http.StatusBadRequest, "Missing required fields")
	}

	id, err := h.service.Create(c.Request().Context(), input)
	if err != nil {
		return common.Error(c, http.StatusInternalServerError, "Failed to create user")
	}

	return common.Success(c, http.StatusCreated, "User created successfully", map[string]string{"id": id})
}
