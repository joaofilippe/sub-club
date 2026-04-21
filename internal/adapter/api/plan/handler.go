package plan

import (
	"strconv"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/joaofilippe/subclub/internal/adapter/api/common"
	"github.com/joaofilippe/subclub/internal/adapter/service"
	domain "github.com/joaofilippe/subclub/internal/domain/plan"
)

type PlanDTO struct {
	ID            string  `json:"id"`
	Code          string  `json:"code"`
	Name          string  `json:"name"`
	Description   string  `json:"description"`
	ProductValue  float64 `json:"productValue"`
	DiscountValue float64 `json:"discountValue"`
	Price         float64 `json:"price"`
	IntervalDays  int     `json:"intervalDays"`
	Active        bool    `json:"active"`
	ImageURL      *string `json:"imageUrl,omitempty"`
	CreatedAt     string  `json:"createdAt"`
}

type PlanInputDTO struct {
	Code          string  `json:"code"`
	Name          string  `json:"name"`
	Description   string  `json:"description"`
	ProductValue  float64 `json:"productValue"`
	DiscountValue float64 `json:"discountValue"`
	Price         float64 `json:"price"`
	IntervalDays  int     `json:"intervalDays"`
	Active        bool    `json:"active"`
	ImageURL      *string `json:"imageUrl,omitempty"`
}

type PlanHandler struct {
	service *services.PlanService
}

func NewPlanHandler(s *services.PlanService) *PlanHandler {
	return &PlanHandler{service: s}
}

func (h *PlanHandler) Create(c echo.Context) error {
	var input PlanInputDTO
	if err := c.Bind(&input); err != nil {
		return c.JSON(http.StatusBadRequest, common.Response{Message: err.Error()})
	}

	img := ""
	if input.ImageURL != nil {
		img = *input.ImageURL
	}

	plan := &domain.Plan{
		Code:          input.Code,
		Name:          input.Name,
		Description:   input.Description,
		ProductValue:  input.ProductValue,
		DiscountValue: input.DiscountValue,
		Price:         input.Price,
		IntervalDays:  input.IntervalDays,
		Active:        input.Active,
		ImageURL:      img,
	}

	if err := h.service.Create(c.Request().Context(), plan); err != nil {
		return c.JSON(http.StatusInternalServerError, common.Response{Message: err.Error()})
	}
	return c.JSON(http.StatusCreated, mapDomainToDTO(plan))
}

func (h *PlanHandler) Get(c echo.Context) error {
	plan, err := h.service.GetByID(c.Request().Context(), c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusNotFound, common.Response{Message: "not found"})
	}
	return c.JSON(http.StatusOK, mapDomainToDTO(plan))
}

func (h *PlanHandler) Update(c echo.Context) error {
	existing, err := h.service.GetByID(c.Request().Context(), c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusNotFound, common.Response{Message: "not found"})
	}

	var input PlanInputDTO
	if err := c.Bind(&input); err != nil {
		return c.JSON(http.StatusBadRequest, common.Response{Message: err.Error()})
	}

	existing.Code = input.Code
	existing.Name = input.Name
	existing.Description = input.Description
	existing.ProductValue = input.ProductValue
	existing.DiscountValue = input.DiscountValue
	existing.Price = input.Price
	existing.IntervalDays = input.IntervalDays
	existing.Active = input.Active
	if input.ImageURL != nil {
		existing.ImageURL = *input.ImageURL
	}

	if err := h.service.Update(c.Request().Context(), existing); err != nil {
		return c.JSON(http.StatusInternalServerError, common.Response{Message: err.Error()})
	}
	return c.JSON(http.StatusOK, mapDomainToDTO(existing))
}

func (h *PlanHandler) Delete(c echo.Context) error {
	if err := h.service.Delete(c.Request().Context(), c.Param("id")); err != nil {
		return c.JSON(http.StatusInternalServerError, common.Response{Message: err.Error()})
	}
	return c.NoContent(http.StatusOK)
}

func (h *PlanHandler) List(c echo.Context) error {
	filter := domain.Filter{}
	if s := c.QueryParam("search"); s != "" {
		filter.Search = &s
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

	var res []PlanDTO
	for _, p := range list {
		res = append(res, mapDomainToDTO(p))
	}
	if res == nil {
		res = []PlanDTO{}
	}
	return c.JSON(http.StatusOK, res)
}

func mapDomainToDTO(p *domain.Plan) PlanDTO {
	var img *string
	if p.ImageURL != "" {
		img = &p.ImageURL
	}
	return PlanDTO{
		ID:            p.ID,
		Code:          p.Code,
		Name:          p.Name,
		Description:   p.Description,
		ProductValue:  p.ProductValue,
		DiscountValue: p.DiscountValue,
		Price:         p.Price,
		IntervalDays:  p.IntervalDays,
		Active:        p.Active,
		ImageURL:      img,
		CreatedAt:     p.CreatedAt.Format(time.RFC3339),
	}
}
