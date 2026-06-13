package product

import (
	"errors"
	"net/http"
	"strconv"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/joaofilippe/subclub/internal/web/common"
	domain "github.com/joaofilippe/subclub/internal/domain/product"
	"github.com/joaofilippe/subclub/internal/domain/product/model"
)

type ProductDTO struct {
	ID          string  `json:"id"`
	Code        string  `json:"code"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	CostPrice   float64 `json:"costPrice"`
	Category    string  `json:"category"`
	ImageURL    *string `json:"imageUrl,omitempty"`
	Active      bool    `json:"active"`
	CreatedAt   string  `json:"createdAt"`
}

type ProductInputDTO struct {
	Code        string  `json:"code"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	CostPrice   float64 `json:"costPrice"`
	Category    string  `json:"category"`
	ImageURL    *string `json:"imageUrl,omitempty"`
	Active      bool    `json:"active"`
}

type PaginatedProductResponse struct {
	Items      []ProductDTO `json:"items"`
	TotalCount int          `json:"totalCount"`
}

type ProductHandler struct {
	service domain.Service
}

func NewProductHandler(s domain.Service) *ProductHandler {
	return &ProductHandler{service: s}
}

// Create godoc
// @Summary      Create a product
// @Description  Creates a new product
// @Tags         products
// @Accept       json
// @Produce      json
// @Param        product  body      ProductInputDTO  true  "Product data"
// @Success      201      {object}  ProductDTO
// @Failure      400      {object}  common.Response
// @Failure      500      {object}  common.Response
// @Router       /api/v1/products [post]
func (h *ProductHandler) Create(c echo.Context) error {
	var input ProductInputDTO
	if err := c.Bind(&input); err != nil {
		return c.JSON(http.StatusBadRequest, common.Response{Message: err.Error()})
	}

	img := ""
	if input.ImageURL != nil {
		img = *input.ImageURL
	}

	product, err := h.service.Create(c.Request().Context(), model.CreateProductInput{
		Code:        input.Code,
		Name:        input.Name,
		Description: input.Description,
		CostPrice:   input.CostPrice,
		Category:    input.Category,
		ImageURL:    img,
		Active:      input.Active,
	})
	if err != nil {
		return c.JSON(http.StatusInternalServerError, common.Response{Message: err.Error()})
	}
	return c.JSON(http.StatusCreated, mapDomainToDTO(product))
}

// Get godoc
// @Summary      Get a product by ID
// @Description  Retrieves a product by its UUID
// @Tags         products
// @Produce      json
// @Param        id   path      string  true  "Product ID"
// @Success      200  {object}  ProductDTO
// @Failure      404  {object}  common.Response
// @Router       /api/v1/products/{id} [get]
func (h *ProductHandler) Get(c echo.Context) error {
	product, err := h.service.GetByID(c.Request().Context(), c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusNotFound, common.Response{Message: "not found"})
	}
	return c.JSON(http.StatusOK, mapDomainToDTO(product))
}

// Update godoc
// @Summary      Update a product
// @Description  Updates an existing product
// @Tags         products
// @Accept       json
// @Produce      json
// @Param        id       path      string           true  "Product ID"
// @Param        product  body      ProductInputDTO  true  "Updated product data"
// @Success      200      {object}  ProductDTO
// @Failure      400      {object}  common.Response
// @Failure      404      {object}  common.Response
// @Failure      500      {object}  common.Response
// @Router       /api/v1/products/{id} [put]
func (h *ProductHandler) Update(c echo.Context) error {
	var input ProductInputDTO
	if err := c.Bind(&input); err != nil {
		return c.JSON(http.StatusBadRequest, common.Response{Message: err.Error()})
	}

	img := ""
	if input.ImageURL != nil {
		img = *input.ImageURL
	}

	product, err := h.service.Update(c.Request().Context(), model.UpdateProductInput{
		ID:          c.Param("id"),
		Code:        input.Code,
		Name:        input.Name,
		Description: input.Description,
		CostPrice:   input.CostPrice,
		Category:    input.Category,
		ImageURL:    img,
		Active:      input.Active,
	})
	if err != nil {
		if errors.Is(err, model.ErrNotFound) {
			return c.JSON(http.StatusNotFound, common.Response{Message: "not found"})
		}
		return c.JSON(http.StatusInternalServerError, common.Response{Message: err.Error()})
	}
	return c.JSON(http.StatusOK, mapDomainToDTO(product))
}

// Delete godoc
// @Summary      Delete a product
// @Description  Soft deletes a product by ID
// @Tags         products
// @Produce      json
// @Param        id   path      string  true  "Product ID"
// @Success      200  {string}  string  "OK"
// @Failure      500  {object}  common.Response
// @Router       /api/v1/products/{id} [delete]
func (h *ProductHandler) Delete(c echo.Context) error {
	if err := h.service.Delete(c.Request().Context(), c.Param("id")); err != nil {
		return c.JSON(http.StatusInternalServerError, common.Response{Message: err.Error()})
	}
	return c.NoContent(http.StatusOK)
}

// List godoc
// @Summary      List products
// @Description  Get a paginated list of products
// @Tags         products
// @Produce      json
// @Param        search    query     string  false  "Search by name or code"
// @Param        category  query     string  false  "Filter by category"
// @Param        active    query     bool    false  "Filter by active status"
// @Param        page      query     int     false  "Page number"
// @Param        pageSize  query     int     false  "Page size"
// @Success      200       {object}  PaginatedProductResponse
// @Failure      500       {object}  common.Response
// @Router       /api/v1/products [get]
func (h *ProductHandler) List(c echo.Context) error {
	filter := model.Filter{}
	if s := c.QueryParam("search"); s != "" {
		filter.Search = &s
	}
	if cat := c.QueryParam("category"); cat != "" {
		filter.Category = &cat
	}
	if act := c.QueryParam("active"); act != "" {
		b, err := strconv.ParseBool(act)
		if err == nil {
			filter.IsActive = &b
		}
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

	list, err := h.service.List(c.Request().Context(), filter)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, common.Response{Message: err.Error()})
	}

	response := PaginatedProductResponse{
		Items:      make([]ProductDTO, 0, len(list.Items)),
		TotalCount: list.TotalCount,
	}
	for _, p := range list.Items {
		response.Items = append(response.Items, mapDomainToDTO(p))
	}
	return c.JSON(http.StatusOK, response)
}

func mapDomainToDTO(p *model.Product) ProductDTO {
	var img *string
	if p.ImageURL != "" {
		img = &p.ImageURL
	}
	return ProductDTO{
		ID:          p.ID,
		Code:        p.Code,
		Name:        p.Name,
		Description: p.Description,
		CostPrice:   p.CostPrice,
		Category:    p.Category,
		ImageURL:    img,
		Active:      p.Active,
		CreatedAt:   p.CreatedAt.Format(time.RFC3339),
	}
}
