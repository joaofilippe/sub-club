package repository

import (
	"context"

	"github.com/google/uuid"
	"github.com/joaofilippe/subclub/ent"
	"github.com/joaofilippe/subclub/ent/product"
	domain "github.com/joaofilippe/subclub/internal/domain/product"
)

type ProductEntRepository struct {
	client *ent.Client
}

func NewProductEntRepository(client *ent.Client) *ProductEntRepository {
	return &ProductEntRepository{client: client}
}

func (r *ProductEntRepository) Create(ctx context.Context, p *domain.Product) error {
	id, err := uuid.Parse(p.ID)
	if err != nil {
		id = uuid.New()
	}

	_, err = r.client.Product.Create().
		SetID(id).
		SetCode(p.Code).
		SetName(p.Name).
		SetDescription(p.Description).
		SetCostPrice(p.CostPrice).
		SetCategory(p.Category).
		SetNillableImageURL(&p.ImageURL).
		SetActive(p.Active).
		Save(ctx)
	return err
}

func (r *ProductEntRepository) GetByID(ctx context.Context, id string) (*domain.Product, error) {
	parsedID, err := uuid.Parse(id)
	if err != nil {
		return nil, err
	}

	p, err := r.client.Product.Get(ctx, parsedID)
	if err != nil {
		return nil, err
	}
	return mapEntProductToDomain(p), nil
}

func (r *ProductEntRepository) Update(ctx context.Context, p *domain.Product) error {
	parsedID, err := uuid.Parse(p.ID)
	if err != nil {
		return err
	}

	_, err = r.client.Product.UpdateOneID(parsedID).
		SetCode(p.Code).
		SetName(p.Name).
		SetDescription(p.Description).
		SetCostPrice(p.CostPrice).
		SetCategory(p.Category).
		SetNillableImageURL(&p.ImageURL).
		SetActive(p.Active).
		Save(ctx)
	return err
}

func (r *ProductEntRepository) Delete(ctx context.Context, id string) error {
	parsedID, err := uuid.Parse(id)
	if err != nil {
		return err
	}
	_, err = r.client.Product.UpdateOneID(parsedID).SetActive(false).Save(ctx)
	return err
}

func (r *ProductEntRepository) List(ctx context.Context, filter domain.Filter) (*domain.PaginatedList, error) {
	q := r.client.Product.Query()

	if filter.IsActive != nil {
		q = q.Where(product.ActiveEQ(*filter.IsActive))
	}
	if filter.Search != nil && *filter.Search != "" {
		q = q.Where(product.NameContainsFold(*filter.Search))
	}
	if filter.Category != nil && *filter.Category != "" {
		q = q.Where(product.CategoryEQ(*filter.Category))
	}

	totalCount, err := q.Count(ctx)
	if err != nil {
		return nil, err
	}

	if filter.Page > 0 && filter.PageSize > 0 {
		offset := (filter.Page - 1) * filter.PageSize
		q = q.Offset(offset).Limit(filter.PageSize)
	}

	entProducts, err := q.All(ctx)
	if err != nil {
		return nil, err
	}

	var results []*domain.Product
	for _, ep := range entProducts {
		results = append(results, mapEntProductToDomain(ep))
	}
	if results == nil {
		results = []*domain.Product{}
	}

	return &domain.PaginatedList{
		Items:      results,
		TotalCount: totalCount,
	}, nil
}

func mapEntProductToDomain(ep *ent.Product) *domain.Product {
	return &domain.Product{
		ID:          ep.ID.String(),
		Code:        ep.Code,
		Name:        ep.Name,
		Description: ep.Description,
		CostPrice:   ep.CostPrice,
		Category:    ep.Category,
		ImageURL:    ep.ImageURL,
		Active:      ep.Active,
		CreatedAt:   ep.CreatedAt,
	}
}
