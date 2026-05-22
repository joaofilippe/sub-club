package repository

import (
	"context"

	"github.com/google/uuid"
	"github.com/joaofilippe/subclub/ent"
	entaccount "github.com/joaofilippe/subclub/ent/account"
	"github.com/joaofilippe/subclub/internal/domain/account"
	"github.com/joaofilippe/subclub/internal/domain/account/model"
)

type AccountEntRepository struct {
	client *ent.Client
}

func NewAccountEntRepository(client *ent.Client) account.Repository {
	return &AccountEntRepository{client: client}
}

func (r *AccountEntRepository) Create(ctx context.Context, a *model.Account) error {
	id, err := uuid.Parse(a.ID)
	if err != nil {
		id = uuid.New()
	}

	builder := r.client.Account.Create().
		SetID(id).
		SetName(a.Name).
		SetEmail(a.Email).
		SetDocument(a.Document).
		SetSlug(a.Slug).
		SetSubscriptionStatus(entaccount.SubscriptionStatus(a.SubscriptionStatus)).
		SetActive(a.Active)

	if a.AccountPlanID != "" {
		planID, err := uuid.Parse(a.AccountPlanID)
		if err == nil {
			builder.SetAccountPlanID(planID)
		}
	}
	if a.SubscriptionExpiresAt != nil {
		builder.SetSubscriptionExpiresAt(*a.SubscriptionExpiresAt)
	}

	_, err = builder.Save(ctx)
	return err
}

func (r *AccountEntRepository) GetByID(ctx context.Context, id string) (*model.Account, error) {
	parsedID, err := uuid.Parse(id)
	if err != nil {
		return nil, err
	}
	a, err := r.client.Account.Get(ctx, parsedID)
	if err != nil {
		return nil, err
	}
	return mapEntAccountToDomain(a), nil
}

func (r *AccountEntRepository) GetBySlug(ctx context.Context, slug string) (*model.Account, error) {
	a, err := r.client.Account.Query().
		Where(entaccount.SlugEQ(slug)).
		First(ctx)
	if err != nil {
		return nil, err
	}
	return mapEntAccountToDomain(a), nil
}

func (r *AccountEntRepository) Update(ctx context.Context, a *model.Account) error {
	parsedID, err := uuid.Parse(a.ID)
	if err != nil {
		return err
	}

	builder := r.client.Account.UpdateOneID(parsedID).
		SetName(a.Name).
		SetEmail(a.Email).
		SetDocument(a.Document).
		SetSubscriptionStatus(entaccount.SubscriptionStatus(a.SubscriptionStatus)).
		SetActive(a.Active)

	if a.AccountPlanID != "" {
		planID, err := uuid.Parse(a.AccountPlanID)
		if err == nil {
			builder.SetAccountPlanID(planID)
		}
	} else {
		builder.ClearAccountPlanID()
	}

	if a.SubscriptionExpiresAt != nil {
		builder.SetSubscriptionExpiresAt(*a.SubscriptionExpiresAt)
	} else {
		builder.ClearSubscriptionExpiresAt()
	}

	_, err = builder.Save(ctx)
	return err
}

func (r *AccountEntRepository) Delete(ctx context.Context, id string) error {
	parsedID, err := uuid.Parse(id)
	if err != nil {
		return err
	}
	_, err = r.client.Account.UpdateOneID(parsedID).SetActive(false).Save(ctx)
	return err
}

func (r *AccountEntRepository) List(ctx context.Context) ([]*model.Account, error) {
	accounts, err := r.client.Account.Query().All(ctx)
	if err != nil {
		return nil, err
	}
	results := make([]*model.Account, 0, len(accounts))
	for _, a := range accounts {
		results = append(results, mapEntAccountToDomain(a))
	}
	return results, nil
}

func mapEntAccountToDomain(a *ent.Account) *model.Account {
	acc := &model.Account{
		ID:                 a.ID.String(),
		Name:               a.Name,
		Email:              a.Email,
		Document:           a.Document,
		Slug:               a.Slug,
		SubscriptionStatus: model.SubscriptionStatus(a.SubscriptionStatus),
		Active:             a.Active,
		CreatedAt:          a.CreatedAt,
	}
	if a.AccountPlanID != uuid.Nil {
		acc.AccountPlanID = a.AccountPlanID.String()
	}
	if a.SubscriptionExpiresAt != nil {
		acc.SubscriptionExpiresAt = a.SubscriptionExpiresAt
	}
	return acc
}
