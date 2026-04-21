package services

import (
	"context"

	domain "github.com/joaofilippe/subclub/internal/domain/client"
)

type ClientService struct {
	repo domain.Repository
}

func NewClientService(repo domain.Repository) *ClientService {
	return &ClientService{repo: repo}
}

func (s *ClientService) Create(ctx context.Context, c *domain.Client) error {
	return s.repo.Create(ctx, c)
}

func (s *ClientService) GetByID(ctx context.Context, id string) (*domain.Client, error) {
	return s.repo.GetByID(ctx, id)
}

func (s *ClientService) Update(ctx context.Context, c *domain.Client) error {
	return s.repo.Update(ctx, c)
}

func (s *ClientService) Delete(ctx context.Context, id string) error {
	return s.repo.Delete(ctx, id)
}

func (s *ClientService) List(ctx context.Context, filter domain.Filter) (*domain.PaginatedList, error) {
	return s.repo.List(ctx, filter)
}
