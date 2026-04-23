package repository

import (
	"context"

	"github.com/google/uuid"
	"github.com/joaofilippe/subclub/ent"
	"github.com/joaofilippe/subclub/ent/plan"
	domain "github.com/joaofilippe/subclub/internal/domain/plan"
)

type PlanEntRepository struct {
	client *ent.Client
}

func NewPlanEntRepository(client *ent.Client) *PlanEntRepository {
	return &PlanEntRepository{client: client}
}

func (r *PlanEntRepository) Create(ctx context.Context, p *domain.Plan) error {
	id, err := uuid.Parse(p.ID)
	if err != nil {
		id = uuid.New()
	}
	_, err = r.client.Plan.Create().
		SetID(id).
		SetCode(p.Code).
		SetName(p.Name).
		SetDescription(p.Description).
		SetProductValue(p.ProductValue).
		SetDiscountValue(p.DiscountValue).
		SetPrice(p.Price).
		SetIntervalDays(p.IntervalDays).
		SetActive(p.Active).
		SetNillableImageURL(&p.ImageURL).
		Save(ctx)
	return err
}

func (r *PlanEntRepository) GetByID(ctx context.Context, id string) (*domain.Plan, error) {
	parsedID, err := uuid.Parse(id)
	if err != nil {
		return nil, err
	}
	p, err := r.client.Plan.Get(ctx, parsedID)
	if err != nil {
		return nil, err
	}
	return mapEntPlanToDomain(p), nil
}

func (r *PlanEntRepository) Update(ctx context.Context, p *domain.Plan) error {
	parsedID, err := uuid.Parse(p.ID)
	if err != nil {
		return err
	}
	_, err = r.client.Plan.UpdateOneID(parsedID).
		SetCode(p.Code).
		SetName(p.Name).
		SetDescription(p.Description).
		SetProductValue(p.ProductValue).
		SetDiscountValue(p.DiscountValue).
		SetPrice(p.Price).
		SetIntervalDays(p.IntervalDays).
		SetActive(p.Active).
		SetNillableImageURL(&p.ImageURL).
		Save(ctx)
	return err
}

func (r *PlanEntRepository) Delete(ctx context.Context, id string) error {
	parsedID, err := uuid.Parse(id)
	if err != nil {
		return err
	}
	// Soft delete for the Plan
	_, err = r.client.Plan.UpdateOneID(parsedID).SetActive(false).Save(ctx)
	return err
}

func (r *PlanEntRepository) List(ctx context.Context, filter domain.Filter) (*domain.PaginatedList, error) {
	q := r.client.Plan.Query()

	if filter.IsActive != nil {
		q = q.Where(plan.ActiveEQ(*filter.IsActive))
	}
	if filter.Search != nil && *filter.Search != "" {
		q = q.Where(plan.NameContainsFold(*filter.Search))
	}

	totalCount, err := q.Count(ctx)
	if err != nil {
		return nil, err
	}

	if filter.Page > 0 && filter.PageSize > 0 {
		offset := (filter.Page - 1) * filter.PageSize
		q = q.Offset(offset).Limit(filter.PageSize)
	}

	entPlans, err := q.All(ctx)
	if err != nil {
		return nil, err
	}

	var results []*domain.Plan
	for _, ep := range entPlans {
		results = append(results, mapEntPlanToDomain(ep))
	}
	if results == nil {
		results = []*domain.Plan{}
	}

	return &domain.PaginatedList{
		Items:      results,
		TotalCount: totalCount,
	}, nil
}

func mapEntPlanToDomain(ep *ent.Plan) *domain.Plan {
	return &domain.Plan{
		ID:            ep.ID.String(),
		Code:          ep.Code,
		Name:          ep.Name,
		Description:   ep.Description,
		ProductValue:  ep.ProductValue,
		DiscountValue: ep.DiscountValue,
		Price:         ep.Price,
		IntervalDays:  ep.IntervalDays,
		Active:        ep.Active,
		ImageURL:      ep.ImageURL,
		CreatedAt:     ep.CreatedAt,
	}
}
