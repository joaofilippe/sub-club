package plan

import (
	"strconv"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/joaofilippe/subclub/internal/adapter/api/common"
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

type PaginatedPlanResponse struct {
	Items      []PlanDTO `json:"items"`
	TotalCount int       `json:"totalCount"`
}

type PlanHandler struct {
	service domain.Service
}

func NewPlanHandler(s domain.Service) *PlanHandler {
	return &PlanHandler{service: s}
}

// Create godoc
// @Summary      Create a plan
// @Description  Creates a new subscription plan
// @Tags         plans
// @Accept       json
// @Produce      json
// @Param        plan  body      PlanInputDTO  true  "Plan data"
// @Success      201   {object}  PlanDTO
// @Failure      400   {object}  common.Response
// @Failure      500   {object}  common.Response
// @Router       /plans [post]
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

// Get godoc
// @Summary      Get a plan by ID
// @Description  Retrieves a plan by its UUID
// @Tags         plans
// @Produce      json
// @Param        id   path      string  true  "Plan ID"
// @Success      200  {object}  PlanDTO
// @Failure      404  {object}  common.Response
// @Router       /plans/{id} [get]
func (h *PlanHandler) Get(c echo.Context) error {
	plan, err := h.service.GetByID(c.Request().Context(), c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusNotFound, common.Response{Message: "not found"})
	}
	return c.JSON(http.StatusOK, mapDomainToDTO(plan))
}

// Update godoc
// @Summary      Update a plan
// @Description  Updates an existing plan
// @Tags         plans
// @Accept       json
// @Produce      json
// @Param        id    path      string        true  "Plan ID"
// @Param        plan  body      PlanInputDTO  true  "Updated plan data"
// @Success      200   {object}  PlanDTO
// @Failure      400   {object}  common.Response
// @Failure      404   {object}  common.Response
// @Failure      500   {object}  common.Response
// @Router       /plans/{id} [put]
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

// Delete godoc
// @Summary      Delete a plan
// @Description  Soft deletes a plan by ID
// @Tags         plans
// @Produce      json
// @Param        id   path      string  true  "Plan ID"
// @Success      200  {string}  string  "OK"
// @Failure      500  {object}  common.Response
// @Router       /plans/{id} [delete]
func (h *PlanHandler) Delete(c echo.Context) error {
	if err := h.service.Delete(c.Request().Context(), c.Param("id")); err != nil {
		return c.JSON(http.StatusInternalServerError, common.Response{Message: err.Error()})
	}
	return c.NoContent(http.StatusOK)
}

// List godoc
// @Summary      List plans
// @Description  Get a paginated list of plans
// @Tags         plans
// @Produce      json
// @Param        search    query     string  false  "Search by name or code"
// @Param        active    query     bool    false  "Filter by active status"
// @Param        page      query     int     false  "Page number"
// @Param        pageSize  query     int     false  "Page size"
// @Success      200       {object}  PaginatedPlanResponse
// @Failure      500       {object}  common.Response
// @Router       /plans [get]
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

	response := PaginatedPlanResponse{
		Items:      make([]PlanDTO, 0, len(list.Items)),
		TotalCount: list.TotalCount,
	}
	for _, p := range list.Items {
		response.Items = append(response.Items, mapDomainToDTO(p))
	}
	if response.Items == nil {
		response.Items = []PlanDTO{}
	}
	return c.JSON(http.StatusOK, response)
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
