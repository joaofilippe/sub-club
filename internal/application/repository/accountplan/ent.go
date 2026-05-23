package repository

import (
	"context"

	"github.com/google/uuid"
	"github.com/joaofilippe/subclub/ent"
	entaccountplan "github.com/joaofilippe/subclub/ent/accountplan"
	"github.com/joaofilippe/subclub/internal/domain/accountplan"
	"github.com/joaofilippe/subclub/internal/domain/accountplan/model"
)

type AccountPlanEntRepository struct {
	client *ent.Client
}

func NewAccountPlanEntRepository(client *ent.Client) accountplan.Repository {
	return &AccountPlanEntRepository{client: client}
}

func (r *AccountPlanEntRepository) Create(ctx context.Context, p *model.AccountPlan) error {
	id, err := uuid.Parse(p.ID)
	if err != nil {
		id = uuid.New()
	}
	_, err = r.client.AccountPlan.Create().
		SetID(id).
		SetName(p.Name).
		SetDescription(p.Description).
		SetPrice(p.Price).
		SetMaxCustomers(p.MaxCustomers).
		SetMaxPlans(p.MaxPlans).
		SetMaxProducts(p.MaxProducts).
		SetActive(p.Active).
		Save(ctx)
	return err
}

func (r *AccountPlanEntRepository) GetByID(ctx context.Context, id string) (*model.AccountPlan, error) {
	parsedID, err := uuid.Parse(id)
	if err != nil {
		return nil, err
	}
	p, err := r.client.AccountPlan.Get(ctx, parsedID)
	if err != nil {
		return nil, err
	}
	return mapEntAccountPlanToDomain(p), nil
}

func (r *AccountPlanEntRepository) Update(ctx context.Context, p *model.AccountPlan) error {
	parsedID, err := uuid.Parse(p.ID)
	if err != nil {
		return err
	}
	_, err = r.client.AccountPlan.UpdateOneID(parsedID).
		SetName(p.Name).
		SetDescription(p.Description).
		SetPrice(p.Price).
		SetMaxCustomers(p.MaxCustomers).
		SetMaxPlans(p.MaxPlans).
		SetMaxProducts(p.MaxProducts).
		SetActive(p.Active).
		Save(ctx)
	return err
}

func (r *AccountPlanEntRepository) Delete(ctx context.Context, id string) error {
	parsedID, err := uuid.Parse(id)
	if err != nil {
		return err
	}
	_, err = r.client.AccountPlan.UpdateOneID(parsedID).SetActive(false).Save(ctx)
	return err
}

func (r *AccountPlanEntRepository) List(ctx context.Context) ([]*model.AccountPlan, error) {
	plans, err := r.client.AccountPlan.Query().
		Where(entaccountplan.ActiveEQ(true)).
		All(ctx)
	if err != nil {
		return nil, err
	}
	results := make([]*model.AccountPlan, 0, len(plans))
	for _, p := range plans {
		results = append(results, mapEntAccountPlanToDomain(p))
	}
	return results, nil
}

func mapEntAccountPlanToDomain(p *ent.AccountPlan) *model.AccountPlan {
	return &model.AccountPlan{
		ID:           p.ID.String(),
		Name:         p.Name,
		Description:  p.Description,
		Price:        p.Price,
		MaxCustomers: p.MaxCustomers,
		MaxPlans:     p.MaxPlans,
		MaxProducts:  p.MaxProducts,
		Active:       p.Active,
		CreatedAt:    p.CreatedAt,
	}
}
