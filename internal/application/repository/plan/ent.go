package repository

import (
	"context"

	"github.com/google/uuid"
	"github.com/joaofilippe/subclub/ent"
	entplan "github.com/joaofilippe/subclub/ent/plan"
	"github.com/joaofilippe/subclub/internal/domain/plan"
	"github.com/joaofilippe/subclub/internal/domain/plan/model"
)

type PlanEntRepository struct {
	client *ent.Client
}

func NewPlanEntRepository(client *ent.Client) plan.Repository {
	return &PlanEntRepository{client: client}
}

func (r *PlanEntRepository) Create(ctx context.Context, p *model.Plan) error {
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

func (r *PlanEntRepository) GetByID(ctx context.Context, id string) (*model.Plan, error) {
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

func (r *PlanEntRepository) Update(ctx context.Context, p *model.Plan) error {
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
	_, err = r.client.Plan.UpdateOneID(parsedID).SetActive(false).Save(ctx)
	return err
}

func (r *PlanEntRepository) List(ctx context.Context, filter model.Filter) (*model.PaginatedList, error) {
	q := r.client.Plan.Query()

	if filter.IsActive != nil {
		q = q.Where(entplan.ActiveEQ(*filter.IsActive))
	}
	if filter.Search != nil && *filter.Search != "" {
		q = q.Where(entplan.NameContainsFold(*filter.Search))
	}

	totalCount, err := q.Count(ctx)
	if err != nil {
		return nil, err
	}

	if filter.Page > 0 && filter.PageSize > 0 {
		q = q.Offset((filter.Page - 1) * filter.PageSize).Limit(filter.PageSize)
	}

	entPlans, err := q.All(ctx)
	if err != nil {
		return nil, err
	}

	results := make([]*model.Plan, 0, len(entPlans))
	for _, ep := range entPlans {
		results = append(results, mapEntPlanToDomain(ep))
	}
	return &model.PaginatedList{Items: results, TotalCount: totalCount}, nil
}

func mapEntPlanToDomain(ep *ent.Plan) *model.Plan {
	return &model.Plan{
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
