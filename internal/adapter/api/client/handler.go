package client

import (
	"net/http"
	"strconv"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/joaofilippe/subclub/internal/adapter/api/common"
	domain "github.com/joaofilippe/subclub/internal/domain/client"
)

type ClientHandler struct {
	service domain.Service
}

func NewClientHandler(clientService domain.Service) *ClientHandler {
	return &ClientHandler{service: clientService}
}

// Create godoc
// @Summary      Create a client
// @Description  Creates a new client in the system
// @Tags         clients
// @Accept       json
// @Produce      json
// @Param        client  body      ClientInputDTO  true  "Client data"
// @Success      201     {object}  ClientDTO
// @Failure      400     {object}  common.Response
// @Failure      500     {object}  common.Response
// @Router       /clients [post]
func (h *ClientHandler) Create(c echo.Context) error {
	var input ClientInputDTO
	if err := c.Bind(&input); err != nil {
		return c.JSON(http.StatusBadRequest, common.Response{Message: "invalid request body"})
	}

	var address *domain.Address
	if input.Address != nil {
		address = &domain.Address{
			ZipCode:      input.Address.ZipCode,
			Street:       input.Address.Street,
			Number:       input.Address.Number,
			Complement:   input.Address.Complement,
			Neighborhood: input.Address.Neighborhood,
			City:         input.Address.City,
			State:        input.Address.State,
		}
	}

	client := &domain.Client{
		Name:     input.Name,
		Email:    input.Email,
		Phone:    input.Phone,
		Document: input.Document,
		Active:   input.Active,
		Address:  address,
	}

	if err := h.service.Create(c.Request().Context(), client); err != nil {
		return c.JSON(http.StatusInternalServerError, common.Response{Message: err.Error()})
	}

	return c.JSON(http.StatusCreated, mapDomainToDTO(client))
}

// Get godoc
// @Summary      Get a client by ID
// @Description  Retrieves a client by their UUID
// @Tags         clients
// @Produce      json
// @Param        id      path      string  true  "Client ID"
// @Success      200     {object}  ClientDTO
// @Failure      404     {object}  common.Response
// @Router       /clients/{id} [get]
func (h *ClientHandler) Get(c echo.Context) error {
	id := c.Param("id")
	client, err := h.service.GetByID(c.Request().Context(), id)
	if err != nil {
		return c.JSON(http.StatusNotFound, common.Response{Message: "client not found"})
	}

	return c.JSON(http.StatusOK, mapDomainToDTO(client))
}

// Update godoc
// @Summary      Update a client
// @Description  Updates an existing client
// @Tags         clients
// @Accept       json
// @Produce      json
// @Param        id      path      string          true  "Client ID"
// @Param        client  body      ClientInputDTO  true  "Updated client data"
// @Success      200     {object}  ClientDTO
// @Failure      400     {object}  common.Response
// @Failure      404     {object}  common.Response
// @Failure      500     {object}  common.Response
// @Router       /clients/{id} [put]
func (h *ClientHandler) Update(c echo.Context) error {
	id := c.Param("id")
	
	// Ensure exists
	existing, err := h.service.GetByID(c.Request().Context(), id)
	if err != nil {
		return c.JSON(http.StatusNotFound, common.Response{Message: "client not found"})
	}

	var input ClientInputDTO
	if err := c.Bind(&input); err != nil {
		return c.JSON(http.StatusBadRequest, common.Response{Message: "invalid request body"})
	}

	var address *domain.Address
	if input.Address != nil {
		address = &domain.Address{
			ZipCode:      input.Address.ZipCode,
			Street:       input.Address.Street,
			Number:       input.Address.Number,
			Complement:   input.Address.Complement,
			Neighborhood: input.Address.Neighborhood,
			City:         input.Address.City,
			State:        input.Address.State,
		}
	}

	existing.Name = input.Name
	existing.Email = input.Email
	existing.Phone = input.Phone
	existing.Document = input.Document
	existing.Active = input.Active
	existing.Address = address

	if err := h.service.Update(c.Request().Context(), existing); err != nil {
		return c.JSON(http.StatusInternalServerError, common.Response{Message: err.Error()})
	}

	return c.JSON(http.StatusOK, mapDomainToDTO(existing))
}

// Delete godoc
// @Summary      Delete a client
// @Description  Soft deletes a client by ID
// @Tags         clients
// @Produce      json
// @Param        id      path      string  true  "Client ID"
// @Success      200     {string}  string  "OK"
// @Failure      500     {object}  common.Response
// @Router       /clients/{id} [delete]
func (h *ClientHandler) Delete(c echo.Context) error {
	id := c.Param("id")
	if err := h.service.Delete(c.Request().Context(), id); err != nil {
		return c.JSON(http.StatusInternalServerError, common.Response{Message: err.Error()})
	}
	return c.NoContent(http.StatusOK) // or http.StatusNoContent
}

// List godoc
// @Summary      List clients
// @Description  Get a paginated list of clients
// @Tags         clients
// @Produce      json
// @Param        search    query     string  false  "Search by name or email"
// @Param        active    query     bool    false  "Filter by active status"
// @Param        page      query     int     false  "Page number"
// @Param        pageSize  query     int     false  "Page size"
// @Success      200       {object}  PaginatedClientResponse
// @Failure      500       {object}  common.Response
// @Router       /clients [get]
func (h *ClientHandler) List(c echo.Context) error {
	search := c.QueryParam("search")
	activeStr := c.QueryParam("active")
	pageStr := c.QueryParam("page")
	pageSizeStr := c.QueryParam("pageSize")

	filter := domain.Filter{}
	if search != "" {
		filter.Search = &search
	}
	if activeStr != "" {
		active := activeStr == "true"
		filter.IsActive = &active
	}
	
	page := 1
	if p, err := strconv.Atoi(pageStr); err == nil && p > 0 {
		page = p
	}
	pageSize := 10
	if ps, err := strconv.Atoi(pageSizeStr); err == nil && ps > 0 {
		pageSize = ps
	}
	filter.Page = page
	filter.PageSize = pageSize

	paginatedList, err := h.service.List(c.Request().Context(), filter)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, common.Response{Message: err.Error()})
	}

	response := PaginatedClientResponse{
		Items:      make([]ClientDTO, 0, len(paginatedList.Items)),
		TotalCount: paginatedList.TotalCount,
	}

	for _, item := range paginatedList.Items {
		response.Items = append(response.Items, mapDomainToDTO(item))
	}

	return c.JSON(http.StatusOK, response)
}

func mapDomainToDTO(c *domain.Client) ClientDTO {
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

	return ClientDTO{
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
