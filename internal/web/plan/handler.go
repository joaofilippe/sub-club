package plan

import (
	"errors"
	"net/http"
	"strconv"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/joaofilippe/subclub/internal/web/common"
	domain "github.com/joaofilippe/subclub/internal/domain/plan"
	"github.com/joaofilippe/subclub/internal/domain/plan/model"
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
// @Success      201   {object}  common.Response{data=PlanDTO}
// @Failure      400   {object}  common.Response
// @Failure      500   {object}  common.Response
// @Router       /api/v1/plans [post]
func (h *PlanHandler) Create(c echo.Context) error {
	var input PlanInputDTO
	if err := c.Bind(&input); err != nil {
		return common.Error(c, http.StatusBadRequest, err.Error())
	}

	img := ""
	if input.ImageURL != nil {
		img = *input.ImageURL
	}

	p, err := h.service.Create(c.Request().Context(), model.CreatePlanInput{
		Code:          input.Code,
		Name:          input.Name,
		Description:   input.Description,
		ProductValue:  input.ProductValue,
		DiscountValue: input.DiscountValue,
		Price:         input.Price,
		IntervalDays:  input.IntervalDays,
		Active:        input.Active,
		ImageURL:      img,
	})
	if err != nil {
		return common.Error(c, http.StatusInternalServerError, err.Error())
	}
	return common.Success(c, http.StatusCreated, "Plan created", mapDomainToDTO(p))
}

// Get godoc
// @Summary      Get a plan by ID
// @Description  Retrieves a plan by its UUID
// @Tags         plans
// @Produce      json
// @Param        id   path      string  true  "Plan ID"
// @Success      200  {object}  common.Response{data=PlanDTO}
// @Failure      404  {object}  common.Response
// @Router       /api/v1/plans/{id} [get]
func (h *PlanHandler) Get(c echo.Context) error {
	p, err := h.service.GetByID(c.Request().Context(), c.Param("id"))
	if err != nil {
		return common.Error(c, http.StatusNotFound, "not found")
	}
	return common.Success(c, http.StatusOK, "OK", mapDomainToDTO(p))
}

// Update godoc
// @Summary      Update a plan
// @Description  Updates an existing plan
// @Tags         plans
// @Accept       json
// @Produce      json
// @Param        id    path      string        true  "Plan ID"
// @Param        plan  body      PlanInputDTO  true  "Updated plan data"
// @Success      200   {object}  common.Response{data=PlanDTO}
// @Failure      400   {object}  common.Response
// @Failure      404   {object}  common.Response
// @Failure      500   {object}  common.Response
// @Router       /api/v1/plans/{id} [put]
func (h *PlanHandler) Update(c echo.Context) error {
	var input PlanInputDTO
	if err := c.Bind(&input); err != nil {
		return common.Error(c, http.StatusBadRequest, err.Error())
	}

	img := ""
	if input.ImageURL != nil {
		img = *input.ImageURL
	}

	p, err := h.service.Update(c.Request().Context(), model.UpdatePlanInput{
		ID:            c.Param("id"),
		Code:          input.Code,
		Name:          input.Name,
		Description:   input.Description,
		ProductValue:  input.ProductValue,
		DiscountValue: input.DiscountValue,
		Price:         input.Price,
		IntervalDays:  input.IntervalDays,
		Active:        input.Active,
		ImageURL:      img,
	})
	if err != nil {
		if errors.Is(err, model.ErrNotFound) {
			return common.Error(c, http.StatusNotFound, "not found")
		}
		return common.Error(c, http.StatusInternalServerError, err.Error())
	}
	return common.Success(c, http.StatusOK, "Plan updated", mapDomainToDTO(p))
}

// Delete godoc
// @Summary      Delete a plan
// @Description  Soft deletes a plan by ID
// @Tags         plans
// @Produce      json
// @Param        id   path      string  true  "Plan ID"
// @Success      200  {object}  common.Response
// @Failure      500  {object}  common.Response
// @Router       /api/v1/plans/{id} [delete]
func (h *PlanHandler) Delete(c echo.Context) error {
	if err := h.service.Delete(c.Request().Context(), c.Param("id")); err != nil {
		return common.Error(c, http.StatusInternalServerError, err.Error())
	}
	return common.Success(c, http.StatusOK, "Plan deleted", nil)
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
// @Success      200       {object}  common.Response{data=PaginatedPlanResponse}
// @Failure      500       {object}  common.Response
// @Router       /api/v1/plans [get]
func (h *PlanHandler) List(c echo.Context) error {
	filter := model.Filter{}
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
		return common.Error(c, http.StatusInternalServerError, err.Error())
	}

	response := PaginatedPlanResponse{
		Items:      make([]PlanDTO, 0, len(list.Items)),
		TotalCount: list.TotalCount,
	}
	for _, p := range list.Items {
		response.Items = append(response.Items, mapDomainToDTO(p))
	}
	return common.Success(c, http.StatusOK, "OK", response)
}

func mapDomainToDTO(p *model.Plan) PlanDTO {
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
