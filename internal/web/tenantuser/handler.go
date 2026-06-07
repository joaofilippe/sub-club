package tenantuser

import (
	"errors"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
	domain "github.com/joaofilippe/subclub/internal/domain/tenantuser"
	"github.com/joaofilippe/subclub/internal/domain/tenantuser/model"
	"github.com/joaofilippe/subclub/internal/web/common"
)

type TenantUserHandler struct {
	service domain.Service
}

func NewTenantUserHandler(service domain.Service) *TenantUserHandler {
	return &TenantUserHandler{service: service}
}

// Create godoc
// @Summary      Create a tenant user
// @Description  Creates a new user (admin, operations, etc.) for the current tenant. Requires tenant admin role.
// @Tags         users
// @Accept       json
// @Produce      json
// @Param        user  body      TenantUserInputDTO  true  "User data"
// @Success      201   {object}  TenantUserDTO
// @Failure      400   {object}  common.Response
// @Failure      403   {object}  common.Response
// @Failure      500   {object}  common.Response
// @Router       /users [post]
func (h *TenantUserHandler) Create(c echo.Context) error {
	var input TenantUserInputDTO
	if err := c.Bind(&input); err != nil {
		return c.JSON(http.StatusBadRequest, common.Response{Message: "invalid request body"})
	}
	if input.Name == "" || input.Email == "" || input.Password == "" || input.Role == "" {
		return c.JSON(http.StatusBadRequest, common.Response{Message: "name, email, password and role are required"})
	}

	result, err := h.service.Create(c.Request().Context(), model.CreateTenantUserInput{
		Name:     input.Name,
		Email:    input.Email,
		Password: input.Password,
		Role:     model.Role(input.Role),
	})
	if err != nil {
		return c.JSON(http.StatusInternalServerError, common.Response{Message: err.Error()})
	}
	return c.JSON(http.StatusCreated, mapToDTO(result))
}

// Get godoc
// @Summary      Get a tenant user
// @Description  Returns a tenant user by ID. Requires tenant admin role.
// @Tags         users
// @Produce      json
// @Param        id  path      string  true  "User ID"
// @Success      200 {object}  TenantUserDTO
// @Failure      403 {object}  common.Response
// @Failure      404 {object}  common.Response
// @Router       /users/{id} [get]
func (h *TenantUserHandler) Get(c echo.Context) error {
	result, err := h.service.GetByID(c.Request().Context(), c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusNotFound, common.Response{Message: "user not found"})
	}
	return c.JSON(http.StatusOK, mapToDTO(result))
}

// List godoc
// @Summary      List tenant users
// @Description  Returns all active users for the current tenant. Requires tenant admin role.
// @Tags         users
// @Produce      json
// @Success      200 {array}   TenantUserDTO
// @Failure      403 {object}  common.Response
// @Failure      500 {object}  common.Response
// @Router       /users [get]
func (h *TenantUserHandler) List(c echo.Context) error {
	results, err := h.service.List(c.Request().Context())
	if err != nil {
		return c.JSON(http.StatusInternalServerError, common.Response{Message: err.Error()})
	}
	dtos := make([]TenantUserDTO, len(results))
	for i, u := range results {
		dtos[i] = mapToDTO(u)
	}
	return c.JSON(http.StatusOK, dtos)
}

// Update godoc
// @Summary      Update a tenant user
// @Description  Updates name and/or role of a tenant user. Requires tenant admin role.
// @Tags         users
// @Accept       json
// @Produce      json
// @Param        id    path      string                    true  "User ID"
// @Param        user  body      UpdateTenantUserInputDTO  true  "Updated user data"
// @Success      200   {object}  TenantUserDTO
// @Failure      400   {object}  common.Response
// @Failure      403   {object}  common.Response
// @Failure      404   {object}  common.Response
// @Failure      500   {object}  common.Response
// @Router       /users/{id} [put]
func (h *TenantUserHandler) Update(c echo.Context) error {
	var input UpdateTenantUserInputDTO
	if err := c.Bind(&input); err != nil {
		return c.JSON(http.StatusBadRequest, common.Response{Message: "invalid request body"})
	}

	result, err := h.service.Update(c.Request().Context(), model.UpdateTenantUserInput{
		ID:   c.Param("id"),
		Name: input.Name,
		Role: model.Role(input.Role),
	})
	if err != nil {
		if errors.Is(err, model.ErrNotFound) {
			return c.JSON(http.StatusNotFound, common.Response{Message: "user not found"})
		}
		return c.JSON(http.StatusInternalServerError, common.Response{Message: err.Error()})
	}
	return c.JSON(http.StatusOK, mapToDTO(result))
}

// Delete godoc
// @Summary      Delete a tenant user
// @Description  Soft deletes a tenant user by ID. Requires tenant admin role.
// @Tags         users
// @Produce      json
// @Param        id  path      string  true  "User ID"
// @Success      200 {string}  string  "OK"
// @Failure      403 {object}  common.Response
// @Failure      500 {object}  common.Response
// @Router       /users/{id} [delete]
func (h *TenantUserHandler) Delete(c echo.Context) error {
	if err := h.service.Delete(c.Request().Context(), c.Param("id")); err != nil {
		return c.JSON(http.StatusInternalServerError, common.Response{Message: err.Error()})
	}
	return c.NoContent(http.StatusOK)
}

func mapToDTO(u *model.TenantUser) TenantUserDTO {
	return TenantUserDTO{
		ID:        u.ID,
		Name:      u.Name,
		Email:     u.Email,
		Role:      string(u.Role),
		CreatedAt: u.CreatedAt.Format(time.RFC3339),
		UpdatedAt: u.UpdatedAt.Format(time.RFC3339),
	}
}
