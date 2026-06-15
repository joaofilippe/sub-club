package subscription

import (
	"errors"
	"net/http"
	"strconv"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/joaofilippe/subclub/internal/web/common"
	domain "github.com/joaofilippe/subclub/internal/domain/subscription"
	"github.com/joaofilippe/subclub/internal/domain/subscription/model"
)

type SubscriptionDTO struct {
	ID               string `json:"id"`
	CustomerID       string `json:"customerId"`
	CustomerName     string `json:"customerName"`
	PlanID           string `json:"planId"`
	PlanName         string `json:"planName"`
	Status           string `json:"status"`
	ShipmentStatus   string `json:"shipmentStatus"`
	StartDate        string `json:"startDate"`
	NextBillingDate  string `json:"nextBillingDate,omitempty"`
	NextShipmentDate string `json:"nextShipmentDate,omitempty"`
	DaysUntilRenewal int    `json:"daysUntilRenewal"`
	CreatedAt        string `json:"createdAt"`
}

type SubscriptionInputDTO struct {
	CustomerID       string `json:"customerId"`
	PlanID           string `json:"planId"`
	Status           string `json:"status"`
	ShipmentStatus   string `json:"shipmentStatus"`
	StartDate        string `json:"startDate"`
	NextBillingDate  string `json:"nextBillingDate,omitempty"`
	NextShipmentDate string `json:"nextShipmentDate,omitempty"`
}

type PaginatedSubscriptionResponse struct {
	Items      []SubscriptionDTO `json:"items"`
	TotalCount int               `json:"totalCount"`
}

type SubscriptionHandler struct {
	service domain.Service
}

func NewSubscriptionHandler(s domain.Service) *SubscriptionHandler {
	return &SubscriptionHandler{service: s}
}

// Create godoc
// @Summary      Create a subscription
// @Description  Creates a new subscription for a client and plan
// @Tags         subscriptions
// @Accept       json
// @Produce      json
// @Param        subscription  body      SubscriptionInputDTO  true  "Subscription data"
// @Success      201           {object}  common.Response{data=SubscriptionDTO}
// @Failure      400           {object}  common.Response
// @Failure      500           {object}  common.Response
// @Router       /api/v1/subscriptions [post]
func (h *SubscriptionHandler) Create(c echo.Context) error {
	var input SubscriptionInputDTO
	if err := c.Bind(&input); err != nil {
		return common.Error(c, http.StatusBadRequest, err.Error())
	}

	startDate, _ := time.Parse(time.RFC3339, input.StartDate)
	if startDate.IsZero() {
		startDate = time.Now()
	}

	ucInput := model.CreateSubscriptionInput{
		CustomerID:     input.CustomerID,
		PlanID:         input.PlanID,
		Status:         model.Status(input.Status),
		ShipmentStatus: model.ShipmentStatus(input.ShipmentStatus),
		StartDate:      startDate,
	}

	if input.NextBillingDate != "" {
		if nbd, err := time.Parse(time.RFC3339, input.NextBillingDate); err == nil {
			ucInput.NextBillingDate = &nbd
		}
	}
	if input.NextShipmentDate != "" {
		if nsd, err := time.Parse(time.RFC3339, input.NextShipmentDate); err == nil {
			ucInput.NextShipmentDate = &nsd
		}
	}

	sub, err := h.service.Create(c.Request().Context(), ucInput)
	if err != nil {
		return common.Error(c, http.StatusInternalServerError, err.Error())
	}
	return common.Success(c, http.StatusCreated, "Subscription created", mapDomainToDTO(sub))
}

// Get godoc
// @Summary      Get a subscription by ID
// @Description  Retrieves a subscription by its UUID
// @Tags         subscriptions
// @Produce      json
// @Param        id   path      string  true  "Subscription ID"
// @Success      200  {object}  common.Response{data=SubscriptionDTO}
// @Failure      404  {object}  common.Response
// @Router       /api/v1/subscriptions/{id} [get]
func (h *SubscriptionHandler) Get(c echo.Context) error {
	sub, err := h.service.GetByID(c.Request().Context(), c.Param("id"))
	if err != nil {
		return common.Error(c, http.StatusNotFound, "not found")
	}
	return common.Success(c, http.StatusOK, "OK", mapDomainToDTO(sub))
}

// Update godoc
// @Summary      Update a subscription
// @Description  Updates an existing subscription
// @Tags         subscriptions
// @Accept       json
// @Produce      json
// @Param        id            path      string                true  "Subscription ID"
// @Param        subscription  body      SubscriptionInputDTO  true  "Updated subscription data"
// @Success      200           {object}  common.Response{data=SubscriptionDTO}
// @Failure      400           {object}  common.Response
// @Failure      404           {object}  common.Response
// @Failure      500           {object}  common.Response
// @Router       /api/v1/subscriptions/{id} [put]
func (h *SubscriptionHandler) Update(c echo.Context) error {
	var input SubscriptionInputDTO
	if err := c.Bind(&input); err != nil {
		return common.Error(c, http.StatusBadRequest, err.Error())
	}

	ucInput := model.UpdateSubscriptionInput{
		ID:             c.Param("id"),
		Status:         model.Status(input.Status),
		ShipmentStatus: model.ShipmentStatus(input.ShipmentStatus),
	}

	if input.NextBillingDate != "" {
		if nbd, err := time.Parse(time.RFC3339, input.NextBillingDate); err == nil {
			ucInput.NextBillingDate = &nbd
		}
	}
	if input.NextShipmentDate != "" {
		if nsd, err := time.Parse(time.RFC3339, input.NextShipmentDate); err == nil {
			ucInput.NextShipmentDate = &nsd
		}
	}

	sub, err := h.service.Update(c.Request().Context(), ucInput)
	if err != nil {
		if errors.Is(err, model.ErrNotFound) {
			return common.Error(c, http.StatusNotFound, "not found")
		}
		return common.Error(c, http.StatusInternalServerError, err.Error())
	}
	return common.Success(c, http.StatusOK, "Subscription updated", mapDomainToDTO(sub))
}

// Delete godoc
// @Summary      Delete a subscription
// @Description  Soft deletes a subscription by ID
// @Tags         subscriptions
// @Produce      json
// @Param        id   path      string  true  "Subscription ID"
// @Success      200  {object}  common.Response
// @Failure      500  {object}  common.Response
// @Router       /api/v1/subscriptions/{id} [delete]
func (h *SubscriptionHandler) Delete(c echo.Context) error {
	if err := h.service.Delete(c.Request().Context(), c.Param("id")); err != nil {
		return common.Error(c, http.StatusInternalServerError, err.Error())
	}
	return common.Success(c, http.StatusOK, "Subscription deleted", nil)
}

// List godoc
// @Summary      List subscriptions
// @Description  Get a paginated list of subscriptions
// @Tags         subscriptions
// @Produce      json
// @Param        customerId  query     string  false  "Filter by Customer ID"
// @Param        status    query     string  false  "Filter by Status"
// @Param        page      query     int     false  "Page number"
// @Param        pageSize  query     int     false  "Page size"
// @Success      200       {object}  common.Response{data=PaginatedSubscriptionResponse}
// @Failure      500       {object}  common.Response
// @Router       /api/v1/subscriptions [get]
func (h *SubscriptionHandler) List(c echo.Context) error {
	filter := model.Filter{}
	if s := c.QueryParam("search"); s != "" {
		filter.Search = &s
	}
	if st := c.QueryParam("status"); st != "" {
		status := model.Status(st)
		filter.Status = &status
	}
	if cid := c.QueryParam("customerId"); cid != "" {
		filter.CustomerID = &cid
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
		return common.Error(c, http.StatusInternalServerError, err.Error())
	}

	response := PaginatedSubscriptionResponse{
		Items:      make([]SubscriptionDTO, 0, len(paginatedList.Items)),
		TotalCount: paginatedList.TotalCount,
	}
	for _, item := range paginatedList.Items {
		response.Items = append(response.Items, mapDomainToDTO(item))
	}

	return common.Success(c, http.StatusOK, "OK", response)
}

func mapDomainToDTO(s *model.Subscription) SubscriptionDTO {
	nbd := ""
	if s.NextBillingDate != nil {
		nbd = s.NextBillingDate.Format(time.RFC3339)
	}
	nsd := ""
	if s.NextShipmentDate != nil {
		nsd = s.NextShipmentDate.Format(time.RFC3339)
	}

	return SubscriptionDTO{
		ID:               s.ID,
		CustomerID:       s.CustomerID,
		CustomerName:     s.CustomerName,
		PlanID:           s.PlanID,
		PlanName:         s.PlanName,
		Status:           string(s.Status),
		ShipmentStatus:   string(s.ShipmentStatus),
		StartDate:        s.StartDate.Format(time.RFC3339),
		NextBillingDate:  nbd,
		NextShipmentDate: nsd,
		DaysUntilRenewal: s.DaysUntilRenewal,
		CreatedAt:        s.CreatedAt.Format(time.RFC3339),
	}
}
