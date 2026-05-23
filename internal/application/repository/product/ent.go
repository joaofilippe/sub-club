package repository

import (
	"context"

	"github.com/google/uuid"
	"github.com/joaofilippe/subclub/ent"
	entproduct "github.com/joaofilippe/subclub/ent/product"
	"github.com/joaofilippe/subclub/internal/domain/product"
	"github.com/joaofilippe/subclub/internal/domain/product/model"
	"github.com/joaofilippe/subclub/internal/infra/tenantctx"
)

type ProductEntRepository struct {
	client *ent.Client
}

func NewProductEntRepository(client *ent.Client) product.Repository {
	return &ProductEntRepository{client: client}
}

func (r *ProductEntRepository) tenantClient(ctx context.Context) *ent.Client {
	if c := tenantctx.TenantClientFromContext(ctx); c != nil {
		return c
	}
	return r.client
}

func (r *ProductEntRepository) Create(ctx context.Context, p *model.Product) error {
	id, err := uuid.Parse(p.ID)
	if err != nil {
		id = uuid.New()
	}
	_, err = r.tenantClient(ctx).Product.Create().
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

func (r *ProductEntRepository) GetByID(ctx context.Context, id string) (*model.Product, error) {
	parsedID, err := uuid.Parse(id)
	if err != nil {
		return nil, err
	}
	p, err := r.tenantClient(ctx).Product.Get(ctx, parsedID)
	if err != nil {
		return nil, err
	}
	return mapEntProductToDomain(p), nil
}

func (r *ProductEntRepository) Update(ctx context.Context, p *model.Product) error {
	parsedID, err := uuid.Parse(p.ID)
	if err != nil {
		return err
	}
	_, err = r.tenantClient(ctx).Product.UpdateOneID(parsedID).
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
	_, err = r.tenantClient(ctx).Product.UpdateOneID(parsedID).SetActive(false).Save(ctx)
	return err
}

func (r *ProductEntRepository) List(ctx context.Context, filter model.Filter) (*model.PaginatedList, error) {
	q := r.tenantClient(ctx).Product.Query()

	if filter.IsActive != nil {
		q = q.Where(entproduct.ActiveEQ(*filter.IsActive))
	}
	if filter.Search != nil && *filter.Search != "" {
		q = q.Where(entproduct.NameContainsFold(*filter.Search))
	}
	if filter.Category != nil && *filter.Category != "" {
		q = q.Where(entproduct.CategoryEQ(*filter.Category))
	}

	totalCount, err := q.Count(ctx)
	if err != nil {
		return nil, err
	}

	if filter.Page > 0 && filter.PageSize > 0 {
		q = q.Offset((filter.Page - 1) * filter.PageSize).Limit(filter.PageSize)
	}

	entProducts, err := q.All(ctx)
	if err != nil {
		return nil, err
	}

	results := make([]*model.Product, 0, len(entProducts))
	for _, ep := range entProducts {
		results = append(results, mapEntProductToDomain(ep))
	}
	return &model.PaginatedList{Items: results, TotalCount: totalCount}, nil
}

func mapEntProductToDomain(ep *ent.Product) *model.Product {
	return &model.Product{
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
