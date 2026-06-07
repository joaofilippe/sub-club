package account

import (
	"errors"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
	domain "github.com/joaofilippe/subclub/internal/domain/account"
	"github.com/joaofilippe/subclub/internal/domain/account/model"
	"github.com/joaofilippe/subclub/internal/web/common"
)

type AccountHandler struct {
	service domain.Service
}

func NewAccountHandler(service domain.Service) *AccountHandler {
	return &AccountHandler{service: service}
}

// Create godoc
// @Summary      Create an account
// @Description  Registers a new account (SubClub client)
// @Tags         accounts
// @Accept       json
// @Produce      json
// @Param        account  body      AccountInputDTO  true  "Account data"
// @Success      201      {object}  AccountDTO
// @Failure      400      {object}  common.Response
// @Failure      500      {object}  common.Response
// @Router       /accounts [post]
func (h *AccountHandler) Create(c echo.Context) error {
	var input AccountInputDTO
	if err := c.Bind(&input); err != nil {
		return c.JSON(http.StatusBadRequest, common.Response{Message: "invalid request body"})
	}

	result, ownerPassword, err := h.service.Create(c.Request().Context(), model.CreateAccountInput{
		Name:          input.Name,
		Email:         input.Email,
		Document:      input.Document,
		Slug:          input.Slug,
		AccountPlanID: input.AccountPlanID,
	})
	if err != nil {
		return c.JSON(http.StatusInternalServerError, common.Response{Message: err.Error()})
	}
	dto := mapDomainToDTO(result)
	dto.OwnerPassword = &ownerPassword
	return c.JSON(http.StatusCreated, dto)
}

// Get godoc
// @Summary      Get an account by ID
// @Description  Retrieves an account by UUID
// @Tags         accounts
// @Produce      json
// @Param        id   path      string  true  "Account ID"
// @Success      200  {object}  AccountDTO
// @Failure      404  {object}  common.Response
// @Router       /accounts/{id} [get]
func (h *AccountHandler) Get(c echo.Context) error {
	result, err := h.service.GetByID(c.Request().Context(), c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusNotFound, common.Response{Message: "account not found"})
	}
	return c.JSON(http.StatusOK, mapDomainToDTO(result))
}

// Update godoc
// @Summary      Update an account
// @Description  Updates an existing account
// @Tags         accounts
// @Accept       json
// @Produce      json
// @Param        id       path      string                true  "Account ID"
// @Param        account  body      UpdateAccountInputDTO  true  "Updated account data"
// @Success      200      {object}  AccountDTO
// @Failure      400      {object}  common.Response
// @Failure      404      {object}  common.Response
// @Failure      500      {object}  common.Response
// @Router       /accounts/{id} [put]
func (h *AccountHandler) Update(c echo.Context) error {
	var input UpdateAccountInputDTO
	if err := c.Bind(&input); err != nil {
		return c.JSON(http.StatusBadRequest, common.Response{Message: "invalid request body"})
	}

	ucInput := model.UpdateAccountInput{
		ID:                 c.Param("id"),
		Name:               input.Name,
		Email:              input.Email,
		Document:           input.Document,
		AccountPlanID:      input.AccountPlanID,
		SubscriptionStatus: model.SubscriptionStatus(input.SubscriptionStatus),
		Active:             input.Active,
	}
	if input.SubscriptionExpiresAt != nil {
		if t, err := time.Parse(time.RFC3339, *input.SubscriptionExpiresAt); err == nil {
			ucInput.SubscriptionExpiresAt = &t
		}
	}

	result, err := h.service.Update(c.Request().Context(), ucInput)
	if err != nil {
		if errors.Is(err, model.ErrNotFound) {
			return c.JSON(http.StatusNotFound, common.Response{Message: "account not found"})
		}
		return c.JSON(http.StatusInternalServerError, common.Response{Message: err.Error()})
	}
	return c.JSON(http.StatusOK, mapDomainToDTO(result))
}

// Delete godoc
// @Summary      Delete an account
// @Description  Soft deletes an account by ID
// @Tags         accounts
// @Produce      json
// @Param        id   path      string  true  "Account ID"
// @Success      200  {string}  string  "OK"
// @Failure      500  {object}  common.Response
// @Router       /accounts/{id} [delete]
func (h *AccountHandler) Delete(c echo.Context) error {
	if err := h.service.Delete(c.Request().Context(), c.Param("id")); err != nil {
		return c.JSON(http.StatusInternalServerError, common.Response{Message: err.Error()})
	}
	return c.NoContent(http.StatusOK)
}

// List godoc
// @Summary      List accounts
// @Description  Returns all accounts
// @Tags         accounts
// @Produce      json
// @Success      200  {array}   AccountDTO
// @Failure      500  {object}  common.Response
// @Router       /accounts [get]
func (h *AccountHandler) List(c echo.Context) error {
	results, err := h.service.List(c.Request().Context())
	if err != nil {
		return c.JSON(http.StatusInternalServerError, common.Response{Message: err.Error()})
	}
	dtos := make([]AccountDTO, 0, len(results))
	for _, a := range results {
		dtos = append(dtos, mapDomainToDTO(a))
	}
	return c.JSON(http.StatusOK, dtos)
}

func mapDomainToDTO(a *model.Account) AccountDTO {
	dto := AccountDTO{
		ID:                 a.ID,
		Name:               a.Name,
		Email:              a.Email,
		Document:           a.Document,
		Slug:               a.Slug,
		AccountPlanID:      a.AccountPlanID,
		SubscriptionStatus: string(a.SubscriptionStatus),
		Active:             a.Active,
		CreatedAt:          a.CreatedAt.Format(time.RFC3339),
	}
	if a.SubscriptionExpiresAt != nil {
		s := a.SubscriptionExpiresAt.Format(time.RFC3339)
		dto.SubscriptionExpiresAt = &s
	}
	return dto
}
