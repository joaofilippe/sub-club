package product

import (
	"strconv"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/joaofilippe/subclub/internal/adapter/api/common"
	"github.com/joaofilippe/subclub/internal/adapter/service"
	domain "github.com/joaofilippe/subclub/internal/domain/product"
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

type ProductHandler struct {
	service *services.ProductService
}

func NewProductHandler(s *services.ProductService) *ProductHandler {
	return &ProductHandler{service: s}
}

func (h *ProductHandler) Create(c echo.Context) error {
	var input ProductInputDTO
	if err := c.Bind(&input); err != nil {
		return c.JSON(http.StatusBadRequest, common.Response{Message: err.Error()})
	}

	img := ""
	if input.ImageURL != nil {
		img = *input.ImageURL
	}

	product := &domain.Product{
		Code:        input.Code,
		Name:        input.Name,
		Description: input.Description,
		CostPrice:   input.CostPrice,
		Category:    input.Category,
		ImageURL:    img,
		Active:      input.Active,
	}

	if err := h.service.Create(c.Request().Context(), product); err != nil {
		return c.JSON(http.StatusInternalServerError, common.Response{Message: err.Error()})
	}
	return c.JSON(http.StatusCreated, mapDomainToDTO(product))
}

func (h *ProductHandler) Get(c echo.Context) error {
	product, err := h.service.GetByID(c.Request().Context(), c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusNotFound, common.Response{Message: "not found"})
	}
	return c.JSON(http.StatusOK, mapDomainToDTO(product))
}

func (h *ProductHandler) Update(c echo.Context) error {
	existing, err := h.service.GetByID(c.Request().Context(), c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusNotFound, common.Response{Message: "not found"})
	}

	var input ProductInputDTO
	if err := c.Bind(&input); err != nil {
		return c.JSON(http.StatusBadRequest, common.Response{Message: err.Error()})
	}

	existing.Code = input.Code
	existing.Name = input.Name
	existing.Description = input.Description
	existing.CostPrice = input.CostPrice
	existing.Category = input.Category
	if input.ImageURL != nil {
		existing.ImageURL = *input.ImageURL
	}
	existing.Active = input.Active

	if err := h.service.Update(c.Request().Context(), existing); err != nil {
		return c.JSON(http.StatusInternalServerError, common.Response{Message: err.Error()})
	}
	return c.JSON(http.StatusOK, mapDomainToDTO(existing))
}

func (h *ProductHandler) Delete(c echo.Context) error {
	if err := h.service.Delete(c.Request().Context(), c.Param("id")); err != nil {
		return c.JSON(http.StatusInternalServerError, common.Response{Message: err.Error()})
	}
	return c.NoContent(http.StatusOK)
}

func (h *ProductHandler) List(c echo.Context) error {
	filter := domain.Filter{}
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

	list, err := h.service.List(c.Request().Context(), filter)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, common.Response{Message: err.Error()})
	}

	var res []ProductDTO
	for _, p := range list {
		res = append(res, mapDomainToDTO(p))
	}
	if res == nil {
		res = []ProductDTO{}
	}
	return c.JSON(http.StatusOK, res)
}

func mapDomainToDTO(p *domain.Product) ProductDTO {
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
