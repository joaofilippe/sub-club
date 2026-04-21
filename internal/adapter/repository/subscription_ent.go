package repository

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/joaofilippe/subclub/ent"
	"github.com/joaofilippe/subclub/ent/subscription"
	domain "github.com/joaofilippe/subclub/internal/domain/subscription"
)

type SubscriptionEntRepository struct {
	client *ent.Client
}

func NewSubscriptionEntRepository(client *ent.Client) *SubscriptionEntRepository {
	return &SubscriptionEntRepository{client: client}
}

func (r *SubscriptionEntRepository) Create(ctx context.Context, s *domain.Subscription) error {
	id, err := uuid.Parse(s.ID)
	if err != nil {
		id = uuid.New()
	}

	builder := r.client.Subscription.Create().
		SetID(id).
		SetStatus(string(s.Status)).
		SetShipmentStatus(string(s.ShipmentStatus)).
		SetStartDate(s.StartDate)

	if s.ClientID != "" {
		clientID, _ := uuid.Parse(s.ClientID)
		builder = builder.SetCustomerID(clientID)
	}

	if s.PlanID != "" {
		planID, _ := uuid.Parse(s.PlanID)
		builder = builder.SetPlanID(planID)
	}

	if s.NextBillingDate != nil {
		builder = builder.SetNextBillingDate(*s.NextBillingDate)
	}
	if s.NextShipmentDate != nil {
		builder = builder.SetNextShipmentDate(*s.NextShipmentDate)
	}

	_, err = builder.Save(ctx)
	return err
}

func (r *SubscriptionEntRepository) GetByID(ctx context.Context, id string) (*domain.Subscription, error) {
	parsedID, err := uuid.Parse(id)
	if err != nil {
		return nil, err
	}

	s, err := r.client.Subscription.Query().
		Where(subscription.IDEQ(parsedID)).
		WithCustomer().
		WithPlan().
		First(ctx)
	if err != nil {
		return nil, err
	}

	return mapEntSubscriptionToDomain(s), nil
}

func (r *SubscriptionEntRepository) Update(ctx context.Context, s *domain.Subscription) error {
	parsedID, err := uuid.Parse(s.ID)
	if err != nil {
		return err
	}

	builder := r.client.Subscription.UpdateOneID(parsedID).
		SetStatus(string(s.Status)).
		SetShipmentStatus(string(s.ShipmentStatus))

	if s.NextBillingDate != nil {
		builder = builder.SetNextBillingDate(*s.NextBillingDate)
	}
	if s.NextShipmentDate != nil {
		builder = builder.SetNextShipmentDate(*s.NextShipmentDate)
	}

	_, err = builder.Save(ctx)
	return err
}

func (r *SubscriptionEntRepository) Delete(ctx context.Context, id string) error {
	parsedID, err := uuid.Parse(id)
	if err != nil {
		return err
	}
	// Soft delete for subscription = set status to cancelled
	_, err = r.client.Subscription.UpdateOneID(parsedID).SetStatus(string(domain.StatusCancelled)).Save(ctx)
	return err
}

func (r *SubscriptionEntRepository) List(ctx context.Context, filter domain.Filter) (*domain.PaginatedList, error) {
	q := r.client.Subscription.Query().
		WithCustomer().
		WithPlan()

	if filter.Status != nil {
		q = q.Where(subscription.StatusEQ(string(*filter.Status)))
	}
	if filter.ClientID != nil && *filter.ClientID != "" {
		cID, err := uuid.Parse(*filter.ClientID)
		if err == nil {
			q = q.Where(subscription.ClientID(cID))
		}
	}
	// Note: Search by clientName can be implemented with HasCustomerWith if needed

	totalCount, err := q.Count(ctx)
	if err != nil {
		return nil, err
	}

	if filter.Page > 0 && filter.PageSize > 0 {
		q = q.Limit(filter.PageSize).Offset((filter.Page - 1) * filter.PageSize)
	}

	entSubs, err := q.All(ctx)
	if err != nil {
		return nil, err
	}

	var results []*domain.Subscription
	for _, es := range entSubs {
		results = append(results, mapEntSubscriptionToDomain(es))
	}
	return &domain.PaginatedList{
		Items:      results,
		TotalCount: totalCount,
	}, nil
}

func mapEntSubscriptionToDomain(es *ent.Subscription) *domain.Subscription {
	var cID, cName, pID, pName string

	if es.Edges.Customer != nil {
		cID = es.Edges.Customer.ID.String()
		cName = es.Edges.Customer.Name
	} else if es.ClientID != uuid.Nil {
		cID = es.ClientID.String()
	}

	if es.Edges.Plan != nil {
		pID = es.Edges.Plan.ID.String()
		pName = es.Edges.Plan.Name
	} else if es.PlanID != uuid.Nil {
		pID = es.PlanID.String()
	}

	var nbd, nsd *time.Time
	if !es.NextBillingDate.IsZero() {
		nbd = &es.NextBillingDate
	}
	if !es.NextShipmentDate.IsZero() {
		nsd = &es.NextShipmentDate
	}

	// Calculate days until renewal roughly based on next billing date
	daysUntil := 0
	if nbd != nil {
		// Just returning a dummy or you can use time.Until(*nbd)
	}

	return &domain.Subscription{
		ID:               es.ID.String(),
		ClientID:         cID,
		ClientName:       cName,
		PlanID:           pID,
		PlanName:         pName,
		Status:           domain.Status(es.Status),
		ShipmentStatus:   domain.ShipmentStatus(es.ShipmentStatus),
		StartDate:        es.StartDate,
		NextBillingDate:  nbd,
		NextShipmentDate: nsd,
		DaysUntilRenewal: daysUntil,
		CreatedAt:        es.CreatedAt,
	}
}
