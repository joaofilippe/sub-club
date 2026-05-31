package module

import (
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
	domain "github.com/joaofilippe/subclub/internal/domain/module"
	"github.com/joaofilippe/subclub/internal/domain/module/model"
	"github.com/joaofilippe/subclub/internal/web/common"
)

type ModuleHandler struct {
	service domain.Service
}

func NewModuleHandler(service domain.Service) *ModuleHandler {
	return &ModuleHandler{service: service}
}

// Create godoc
// @Summary      Create a module
// @Description  Creates a new platform module
// @Tags         modules
// @Accept       json
// @Produce      json
// @Param        module  body      ModuleInputDTO  true  "Module data"
// @Success      201     {object}  ModuleDTO
// @Failure      400     {object}  common.Response
// @Failure      500     {object}  common.Response
// @Router       /modules [post]
func (h *ModuleHandler) Create(c echo.Context) error {
	var input ModuleInputDTO
	if err := c.Bind(&input); err != nil {
		return common.Error(c, http.StatusBadRequest, "invalid request body")
	}

	m, err := h.service.Create(c.Request().Context(), model.CreateModuleInput{Name: input.Name})
	if err != nil {
		return common.Error(c, http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusCreated, mapToDTO(m))
}

// Get godoc
// @Summary      Get a module by ID
// @Description  Retrieves a module by its UUID
// @Tags         modules
// @Produce      json
// @Param        id   path      string  true  "Module ID"
// @Success      200  {object}  ModuleDTO
// @Failure      404  {object}  common.Response
// @Router       /modules/{id} [get]
func (h *ModuleHandler) Get(c echo.Context) error {
	m, err := h.service.GetByID(c.Request().Context(), c.Param("id"))
	if err != nil {
		return common.Error(c, http.StatusNotFound, "module not found")
	}
	return c.JSON(http.StatusOK, mapToDTO(m))
}

// Update godoc
// @Summary      Update a module
// @Description  Updates an existing module
// @Tags         modules
// @Accept       json
// @Produce      json
// @Param        id      path      string                true  "Module ID"
// @Param        module  body      UpdateModuleInputDTO  true  "Updated module data"
// @Success      200     {object}  ModuleDTO
// @Failure      400     {object}  common.Response
// @Failure      404     {object}  common.Response
// @Failure      500     {object}  common.Response
// @Router       /modules/{id} [put]
func (h *ModuleHandler) Update(c echo.Context) error {
	var input UpdateModuleInputDTO
	if err := c.Bind(&input); err != nil {
		return common.Error(c, http.StatusBadRequest, "invalid request body")
	}

	m, err := h.service.Update(c.Request().Context(), model.UpdateModuleInput{
		ID:     c.Param("id"),
		Name:   input.Name,
		Active: input.Active,
	})
	if err != nil {
		return common.Error(c, http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, mapToDTO(m))
}

// Delete godoc
// @Summary      Delete a module
// @Description  Soft deletes a module by ID
// @Tags         modules
// @Produce      json
// @Param        id   path      string  true  "Module ID"
// @Success      204  "No Content"
// @Failure      500  {object}  common.Response
// @Router       /modules/{id} [delete]
func (h *ModuleHandler) Delete(c echo.Context) error {
	if err := h.service.Delete(c.Request().Context(), c.Param("id")); err != nil {
		return common.Error(c, http.StatusInternalServerError, err.Error())
	}
	return c.NoContent(http.StatusNoContent)
}

// List godoc
// @Summary      List modules
// @Description  Returns all active modules
// @Tags         modules
// @Produce      json
// @Success      200  {array}   ModuleDTO
// @Failure      500  {object}  common.Response
// @Router       /modules [get]
func (h *ModuleHandler) List(c echo.Context) error {
	modules, err := h.service.List(c.Request().Context())
	if err != nil {
		return common.Error(c, http.StatusInternalServerError, err.Error())
	}

	result := make([]ModuleDTO, 0, len(modules))
	for _, m := range modules {
		result = append(result, mapToDTO(m))
	}
	return c.JSON(http.StatusOK, result)
}

func mapToDTO(m *model.Module) ModuleDTO {
	return ModuleDTO{
		ID:        m.ID,
		Name:      m.Name,
		Active:    m.Active,
		CreatedAt: m.CreatedAt.Format(time.RFC3339),
	}
}
