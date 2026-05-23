package repository

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/joaofilippe/subclub/ent"
	entsub "github.com/joaofilippe/subclub/ent/subscription"
	"github.com/joaofilippe/subclub/internal/domain/subscription"
	"github.com/joaofilippe/subclub/internal/domain/subscription/model"
	"github.com/joaofilippe/subclub/internal/infra/tenantctx"
)

type SubscriptionEntRepository struct {
	client *ent.Client
}

func NewSubscriptionEntRepository(client *ent.Client) subscription.Repository {
	return &SubscriptionEntRepository{client: client}
}

func (r *SubscriptionEntRepository) tenantClient(ctx context.Context) *ent.Client {
	if c := tenantctx.TenantClientFromContext(ctx); c != nil {
		return c
	}
	return r.client
}

func (r *SubscriptionEntRepository) Create(ctx context.Context, s *model.Subscription) error {
	id, err := uuid.Parse(s.ID)
	if err != nil {
		id = uuid.New()
	}

	builder := r.tenantClient(ctx).Subscription.Create().
		SetID(id).
		SetStatus(string(s.Status)).
		SetShipmentStatus(string(s.ShipmentStatus)).
		SetStartDate(s.StartDate)

	if s.CustomerID != "" {
		customerID, _ := uuid.Parse(s.CustomerID)
		builder = builder.SetCustomerID(customerID)
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

func (r *SubscriptionEntRepository) GetByID(ctx context.Context, id string) (*model.Subscription, error) {
	parsedID, err := uuid.Parse(id)
	if err != nil {
		return nil, err
	}
	s, err := r.tenantClient(ctx).Subscription.Query().
		Where(entsub.IDEQ(parsedID)).
		WithCustomer().
		WithPlan().
		First(ctx)
	if err != nil {
		return nil, err
	}
	return mapEntSubscriptionToDomain(s), nil
}

func (r *SubscriptionEntRepository) Update(ctx context.Context, s *model.Subscription) error {
	parsedID, err := uuid.Parse(s.ID)
	if err != nil {
		return err
	}

	builder := r.tenantClient(ctx).Subscription.UpdateOneID(parsedID).
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
	_, err = r.tenantClient(ctx).Subscription.UpdateOneID(parsedID).SetStatus(string(model.StatusCancelled)).Save(ctx)
	return err
}

func (r *SubscriptionEntRepository) List(ctx context.Context, filter model.Filter) (*model.PaginatedList, error) {
	q := r.tenantClient(ctx).Subscription.Query().WithCustomer().WithPlan()

	if filter.Status != nil {
		q = q.Where(entsub.StatusEQ(string(*filter.Status)))
	}
	if filter.CustomerID != nil && *filter.CustomerID != "" {
		cID, err := uuid.Parse(*filter.CustomerID)
		if err == nil {
			q = q.Where(entsub.CustomerID(cID))
		}
	}

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

	results := make([]*model.Subscription, 0, len(entSubs))
	for _, es := range entSubs {
		results = append(results, mapEntSubscriptionToDomain(es))
	}
	return &model.PaginatedList{Items: results, TotalCount: totalCount}, nil
}

func mapEntSubscriptionToDomain(es *ent.Subscription) *model.Subscription {
	var cID, cName, pID, pName string

	if es.Edges.Customer != nil {
		cID = es.Edges.Customer.ID.String()
		cName = es.Edges.Customer.Name
	} else if es.CustomerID != uuid.Nil {
		cID = es.CustomerID.String()
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

	return &model.Subscription{
		ID:               es.ID.String(),
		CustomerID:       cID,
		CustomerName:     cName,
		PlanID:           pID,
		PlanName:         pName,
		Status:           model.Status(es.Status),
		ShipmentStatus:   model.ShipmentStatus(es.ShipmentStatus),
		StartDate:        es.StartDate,
		NextBillingDate:  nbd,
		NextShipmentDate: nsd,
		CreatedAt:        es.CreatedAt,
	}
}
