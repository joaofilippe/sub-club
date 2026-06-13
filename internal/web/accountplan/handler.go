package accountplan

import (
	"errors"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
	domain "github.com/joaofilippe/subclub/internal/domain/accountplan"
	"github.com/joaofilippe/subclub/internal/domain/accountplan/model"
	"github.com/joaofilippe/subclub/internal/web/common"
)

type AccountPlanHandler struct {
	service domain.Service
}

func NewAccountPlanHandler(service domain.Service) *AccountPlanHandler {
	return &AccountPlanHandler{service: service}
}

// Create godoc
// @Summary      Create an account plan
// @Description  Creates a new SubClub subscription plan
// @Tags         account-plans
// @Accept       json
// @Produce      json
// @Param        plan  body      AccountPlanInputDTO  true  "Account plan data"
// @Success      201   {object}  AccountPlanDTO
// @Failure      400   {object}  common.Response
// @Failure      500   {object}  common.Response
// @Router       /api/v1/account-plans [post]
func (h *AccountPlanHandler) Create(c echo.Context) error {
	var input AccountPlanInputDTO
	if err := c.Bind(&input); err != nil {
		return c.JSON(http.StatusBadRequest, common.Response{Message: "invalid request body"})
	}

	result, err := h.service.Create(c.Request().Context(), model.CreateAccountPlanInput{
		Name:         input.Name,
		Description:  input.Description,
		Price:        input.Price,
		MaxCustomers: input.MaxCustomers,
		MaxPlans:     input.MaxPlans,
		MaxProducts:  input.MaxProducts,
	})
	if err != nil {
		return c.JSON(http.StatusInternalServerError, common.Response{Message: err.Error()})
	}
	return c.JSON(http.StatusCreated, mapDomainToDTO(result))
}

// Get godoc
// @Summary      Get an account plan by ID
// @Description  Retrieves a SubClub subscription plan by UUID
// @Tags         account-plans
// @Produce      json
// @Param        id   path      string  true  "Account Plan ID"
// @Success      200  {object}  AccountPlanDTO
// @Failure      404  {object}  common.Response
// @Router       /api/v1/account-plans/{id} [get]
func (h *AccountPlanHandler) Get(c echo.Context) error {
	result, err := h.service.GetByID(c.Request().Context(), c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusNotFound, common.Response{Message: "account plan not found"})
	}
	return c.JSON(http.StatusOK, mapDomainToDTO(result))
}

// Update godoc
// @Summary      Update an account plan
// @Description  Updates an existing SubClub subscription plan
// @Tags         account-plans
// @Accept       json
// @Produce      json
// @Param        id    path      string               true  "Account Plan ID"
// @Param        plan  body      AccountPlanInputDTO  true  "Updated plan data"
// @Success      200   {object}  AccountPlanDTO
// @Failure      400   {object}  common.Response
// @Failure      404   {object}  common.Response
// @Failure      500   {object}  common.Response
// @Router       /api/v1/account-plans/{id} [put]
func (h *AccountPlanHandler) Update(c echo.Context) error {
	var input AccountPlanInputDTO
	if err := c.Bind(&input); err != nil {
		return c.JSON(http.StatusBadRequest, common.Response{Message: "invalid request body"})
	}

	result, err := h.service.Update(c.Request().Context(), model.UpdateAccountPlanInput{
		ID:           c.Param("id"),
		Name:         input.Name,
		Description:  input.Description,
		Price:        input.Price,
		MaxCustomers: input.MaxCustomers,
		MaxPlans:     input.MaxPlans,
		MaxProducts:  input.MaxProducts,
		Active:       input.Active,
	})
	if err != nil {
		if errors.Is(err, model.ErrNotFound) {
			return c.JSON(http.StatusNotFound, common.Response{Message: "account plan not found"})
		}
		return c.JSON(http.StatusInternalServerError, common.Response{Message: err.Error()})
	}
	return c.JSON(http.StatusOK, mapDomainToDTO(result))
}

// Delete godoc
// @Summary      Delete an account plan
// @Description  Soft deletes an account plan by ID
// @Tags         account-plans
// @Produce      json
// @Param        id   path      string  true  "Account Plan ID"
// @Success      200  {string}  string  "OK"
// @Failure      500  {object}  common.Response
// @Router       /api/v1/account-plans/{id} [delete]
func (h *AccountPlanHandler) Delete(c echo.Context) error {
	if err := h.service.Delete(c.Request().Context(), c.Param("id")); err != nil {
		return c.JSON(http.StatusInternalServerError, common.Response{Message: err.Error()})
	}
	return c.NoContent(http.StatusOK)
}

// List godoc
// @Summary      List account plans
// @Description  Returns all active SubClub subscription plans
// @Tags         account-plans
// @Produce      json
// @Success      200  {array}   AccountPlanDTO
// @Failure      500  {object}  common.Response
// @Router       /api/v1/account-plans [get]
func (h *AccountPlanHandler) List(c echo.Context) error {
	results, err := h.service.List(c.Request().Context())
	if err != nil {
		return c.JSON(http.StatusInternalServerError, common.Response{Message: err.Error()})
	}
	dtos := make([]AccountPlanDTO, 0, len(results))
	for _, p := range results {
		dtos = append(dtos, mapDomainToDTO(p))
	}
	return c.JSON(http.StatusOK, dtos)
}

func mapDomainToDTO(p *model.AccountPlan) AccountPlanDTO {
	return AccountPlanDTO{
		ID:           p.ID,
		Name:         p.Name,
		Description:  p.Description,
		Price:        p.Price,
		MaxCustomers: p.MaxCustomers,
		MaxPlans:     p.MaxPlans,
		MaxProducts:  p.MaxProducts,
		Active:       p.Active,
		CreatedAt:    p.CreatedAt.Format(time.RFC3339),
	}
}
