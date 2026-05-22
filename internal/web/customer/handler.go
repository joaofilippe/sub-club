package customer

import (
	"errors"
	"net/http"
	"strconv"
	"time"

	"github.com/labstack/echo/v4"
	domain "github.com/joaofilippe/subclub/internal/domain/customer"
	"github.com/joaofilippe/subclub/internal/domain/customer/model"
	"github.com/joaofilippe/subclub/internal/web/common"
)

type CustomerHandler struct {
	service domain.Service
}

func NewCustomerHandler(customerService domain.Service) *CustomerHandler {
	return &CustomerHandler{service: customerService}
}

// Create godoc
// @Summary      Create a customer
// @Description  Creates a new customer in the system
// @Tags         customers
// @Accept       json
// @Produce      json
// @Param        customer  body      CustomerInputDTO  true  "Customer data"
// @Success      201       {object}  CustomerDTO
// @Failure      400       {object}  common.Response
// @Failure      500       {object}  common.Response
// @Router       /customers [post]
func (h *CustomerHandler) Create(c echo.Context) error {
	var input CustomerInputDTO
	if err := c.Bind(&input); err != nil {
		return c.JSON(http.StatusBadRequest, common.Response{Message: "invalid request body"})
	}

	ucInput := model.CreateCustomerInput{
		Name:     input.Name,
		Email:    input.Email,
		Phone:    input.Phone,
		Document: input.Document,
		Active:   input.Active,
		Address:  mapAddressDTOToDomain(input.Address),
	}

	cust, err := h.service.Create(c.Request().Context(), ucInput)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, common.Response{Message: err.Error()})
	}

	return c.JSON(http.StatusCreated, mapDomainToDTO(cust))
}

// Get godoc
// @Summary      Get a customer by ID
// @Description  Retrieves a customer by their UUID
// @Tags         customers
// @Produce      json
// @Param        id      path      string  true  "Customer ID"
// @Success      200     {object}  CustomerDTO
// @Failure      404     {object}  common.Response
// @Router       /customers/{id} [get]
func (h *CustomerHandler) Get(c echo.Context) error {
	cust, err := h.service.GetByID(c.Request().Context(), c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusNotFound, common.Response{Message: "customer not found"})
	}
	return c.JSON(http.StatusOK, mapDomainToDTO(cust))
}

// Update godoc
// @Summary      Update a customer
// @Description  Updates an existing customer
// @Tags         customers
// @Accept       json
// @Produce      json
// @Param        id        path      string            true  "Customer ID"
// @Param        customer  body      CustomerInputDTO  true  "Updated customer data"
// @Success      200       {object}  CustomerDTO
// @Failure      400       {object}  common.Response
// @Failure      404       {object}  common.Response
// @Failure      500       {object}  common.Response
// @Router       /customers/{id} [put]
func (h *CustomerHandler) Update(c echo.Context) error {
	var input CustomerInputDTO
	if err := c.Bind(&input); err != nil {
		return c.JSON(http.StatusBadRequest, common.Response{Message: "invalid request body"})
	}

	ucInput := model.UpdateCustomerInput{
		ID:       c.Param("id"),
		Name:     input.Name,
		Email:    input.Email,
		Phone:    input.Phone,
		Document: input.Document,
		Active:   input.Active,
		Address:  mapAddressDTOToDomain(input.Address),
	}

	result, err := h.service.Update(c.Request().Context(), ucInput)
	if err != nil {
		if errors.Is(err, model.ErrNotFound) {
			return c.JSON(http.StatusNotFound, common.Response{Message: "customer not found"})
		}
		return c.JSON(http.StatusInternalServerError, common.Response{Message: err.Error()})
	}

	return c.JSON(http.StatusOK, mapDomainToDTO(result))
}

// Delete godoc
// @Summary      Delete a customer
// @Description  Soft deletes a customer by ID
// @Tags         customers
// @Produce      json
// @Param        id   path      string  true  "Customer ID"
// @Success      200  {string}  string  "OK"
// @Failure      500  {object}  common.Response
// @Router       /customers/{id} [delete]
func (h *CustomerHandler) Delete(c echo.Context) error {
	if err := h.service.Delete(c.Request().Context(), c.Param("id")); err != nil {
		return c.JSON(http.StatusInternalServerError, common.Response{Message: err.Error()})
	}
	return c.NoContent(http.StatusOK)
}

// List godoc
// @Summary      List customers
// @Description  Get a paginated list of customers
// @Tags         customers
// @Produce      json
// @Param        search    query     string  false  "Search by name or email"
// @Param        active    query     bool    false  "Filter by active status"
// @Param        page      query     int     false  "Page number"
// @Param        pageSize  query     int     false  "Page size"
// @Success      200       {object}  PaginatedCustomerResponse
// @Failure      500       {object}  common.Response
// @Router       /customers [get]
func (h *CustomerHandler) List(c echo.Context) error {
	filter := model.Filter{}

	if s := c.QueryParam("search"); s != "" {
		filter.Search = &s
	}
	if act := c.QueryParam("active"); act != "" {
		active := act == "true"
		filter.IsActive = &active
	}

	page := 1
	if p, err := strconv.Atoi(c.QueryParam("page")); err == nil && p > 0 {
		page = p
	}
	pageSize := 10
	if ps, err := strconv.Atoi(c.QueryParam("pageSize")); err == nil && ps > 0 {
		pageSize = ps
	}
	filter.Page = page
	filter.PageSize = pageSize

	paginatedList, err := h.service.List(c.Request().Context(), filter)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, common.Response{Message: err.Error()})
	}

	response := PaginatedCustomerResponse{
		Items:      make([]CustomerDTO, 0, len(paginatedList.Items)),
		TotalCount: paginatedList.TotalCount,
	}
	for _, item := range paginatedList.Items {
		response.Items = append(response.Items, mapDomainToDTO(item))
	}
	return c.JSON(http.StatusOK, response)
}

func mapAddressDTOToDomain(a *AddressDTO) *model.Address {
	if a == nil {
		return nil
	}
	return &model.Address{
		ZipCode:      a.ZipCode,
		Street:       a.Street,
		Number:       a.Number,
		Complement:   a.Complement,
		Neighborhood: a.Neighborhood,
		City:         a.City,
		State:        a.State,
	}
}

func mapDomainToDTO(c *model.Customer) CustomerDTO {
	var addressDTO *AddressDTO
	if c.Address != nil {
		addressDTO = &AddressDTO{
			ZipCode:      c.Address.ZipCode,
			Street:       c.Address.Street,
			Number:       c.Address.Number,
			Complement:   c.Address.Complement,
			Neighborhood: c.Address.Neighborhood,
			City:         c.Address.City,
			State:        c.Address.State,
		}
	}
	return CustomerDTO{
		ID:        c.ID,
		Name:      c.Name,
		Email:     c.Email,
		Phone:     c.Phone,
		Document:  c.Document,
		Active:    c.Active,
		CreatedAt: c.CreatedAt.Format(time.RFC3339),
		Address:   addressDTO,
	}
}
