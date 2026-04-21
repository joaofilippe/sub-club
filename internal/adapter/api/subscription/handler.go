package subscription

import (
	"strconv"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/joaofilippe/subclub/internal/adapter/api/common"
	"github.com/joaofilippe/subclub/internal/adapter/service"
	domain "github.com/joaofilippe/subclub/internal/domain/subscription"
)

type SubscriptionDTO struct {
	ID               string `json:"id"`
	ClientID         string `json:"clientId"`
	ClientName       string `json:"clientName"`
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
	ClientID         string `json:"clientId"`
	PlanID           string `json:"planId"`
	Status           string `json:"status"`
	ShipmentStatus   string `json:"shipmentStatus"`
	StartDate        string `json:"startDate"` // e.g. RFC3339
	NextBillingDate  string `json:"nextBillingDate,omitempty"`
	NextShipmentDate string `json:"nextShipmentDate,omitempty"`
}

type PaginatedSubscriptionResponse struct {
	Items      []SubscriptionDTO `json:"items"`
	TotalCount int               `json:"totalCount"`
}

type SubscriptionHandler struct {
	service *services.SubscriptionService
}

func NewSubscriptionHandler(s *services.SubscriptionService) *SubscriptionHandler {
	return &SubscriptionHandler{service: s}
}

func (h *SubscriptionHandler) Create(c echo.Context) error {
	var input SubscriptionInputDTO
	if err := c.Bind(&input); err != nil {
		return c.JSON(http.StatusBadRequest, common.Response{Message: err.Error()})
	}

	startDate, _ := time.Parse(time.RFC3339, input.StartDate)
	if startDate.IsZero() {
		startDate = time.Now()
	}

	sub := &domain.Subscription{
		ClientID:       input.ClientID,
		PlanID:         input.PlanID,
		Status:         domain.Status(input.Status),
		ShipmentStatus: domain.ShipmentStatus(input.ShipmentStatus),
		StartDate:      startDate,
	}

	if input.NextBillingDate != "" {
		nbd, err := time.Parse(time.RFC3339, input.NextBillingDate)
		if err == nil {
			sub.NextBillingDate = &nbd
		}
	}

	if input.NextShipmentDate != "" {
		nsd, err := time.Parse(time.RFC3339, input.NextShipmentDate)
		if err == nil {
			sub.NextShipmentDate = &nsd
		}
	}

	if err := h.service.Create(c.Request().Context(), sub); err != nil {
		return c.JSON(http.StatusInternalServerError, common.Response{Message: err.Error()})
	}
	return c.JSON(http.StatusCreated, mapDomainToDTO(sub))
}

func (h *SubscriptionHandler) Get(c echo.Context) error {
	sub, err := h.service.GetByID(c.Request().Context(), c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusNotFound, common.Response{Message: "not found"})
	}
	return c.JSON(http.StatusOK, mapDomainToDTO(sub))
}

func (h *SubscriptionHandler) Update(c echo.Context) error {
	existing, err := h.service.GetByID(c.Request().Context(), c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusNotFound, common.Response{Message: "not found"})
	}

	var input SubscriptionInputDTO
	if err := c.Bind(&input); err != nil {
		return c.JSON(http.StatusBadRequest, common.Response{Message: err.Error()})
	}

	existing.Status = domain.Status(input.Status)
	existing.ShipmentStatus = domain.ShipmentStatus(input.ShipmentStatus)
	
	if input.NextBillingDate != "" {
		nbd, err := time.Parse(time.RFC3339, input.NextBillingDate)
		if err == nil {
			existing.NextBillingDate = &nbd
		}
	}
	if input.NextShipmentDate != "" {
		nsd, err := time.Parse(time.RFC3339, input.NextShipmentDate)
		if err == nil {
			existing.NextShipmentDate = &nsd
		}
	}

	if err := h.service.Update(c.Request().Context(), existing); err != nil {
		return c.JSON(http.StatusInternalServerError, common.Response{Message: err.Error()})
	}
	return c.JSON(http.StatusOK, mapDomainToDTO(existing))
}

func (h *SubscriptionHandler) Delete(c echo.Context) error {
	if err := h.service.Delete(c.Request().Context(), c.Param("id")); err != nil {
		return c.JSON(http.StatusInternalServerError, common.Response{Message: err.Error()})
	}
	return c.NoContent(http.StatusOK)
}

func (h *SubscriptionHandler) List(c echo.Context) error {
	filter := domain.Filter{}
	if s := c.QueryParam("search"); s != "" {
		filter.Search = &s
	}
	if st := c.QueryParam("status"); st != "" {
		status := domain.Status(st)
		filter.Status = &status
	}
	if cid := c.QueryParam("clientId"); cid != "" {
		filter.ClientID = &cid
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

	response := PaginatedSubscriptionResponse{
		Items:      make([]SubscriptionDTO, 0, len(paginatedList.Items)),
		TotalCount: paginatedList.TotalCount,
	}

	for _, item := range paginatedList.Items {
		response.Items = append(response.Items, mapDomainToDTO(item))
	}

	return c.JSON(http.StatusOK, response)
}

func mapDomainToDTO(s *domain.Subscription) SubscriptionDTO {
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
		ClientID:         s.ClientID,
		ClientName:       s.ClientName,
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
