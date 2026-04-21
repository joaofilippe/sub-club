package services

import (
	"context"

	domain "github.com/joaofilippe/subclub/internal/domain/product"
)

type ProductService struct {
	repo domain.Repository
}

func NewProductService(repo domain.Repository) *ProductService {
	return &ProductService{repo: repo}
}

func (s *ProductService) Create(ctx context.Context, p *domain.Product) error {
	return s.repo.Create(ctx, p)
}

func (s *ProductService) GetByID(ctx context.Context, id string) (*domain.Product, error) {
	return s.repo.GetByID(ctx, id)
}

func (s *ProductService) Update(ctx context.Context, p *domain.Product) error {
	return s.repo.Update(ctx, p)
}

func (s *ProductService) Delete(ctx context.Context, id string) error {
	return s.repo.Delete(ctx, id)
}

func (s *ProductService) List(ctx context.Context, filter domain.Filter) ([]*domain.Product, error) {
	return s.repo.List(ctx, filter)
}
