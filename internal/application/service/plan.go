package services

import (
	"context"

	domain "github.com/joaofilippe/subclub/internal/domain/plan"
)

type PlanService struct {
	repo domain.Repository
}

func NewPlanService(repo domain.Repository) *PlanService {
	return &PlanService{repo: repo}
}

func (s *PlanService) Create(ctx context.Context, p *domain.Plan) error {
	return s.repo.Create(ctx, p)
}

func (s *PlanService) GetByID(ctx context.Context, id string) (*domain.Plan, error) {
	return s.repo.GetByID(ctx, id)
}

func (s *PlanService) Update(ctx context.Context, p *domain.Plan) error {
	return s.repo.Update(ctx, p)
}

func (s *PlanService) Delete(ctx context.Context, id string) error {
	return s.repo.Delete(ctx, id)
}

func (s *PlanService) List(ctx context.Context, filter domain.Filter) (*domain.PaginatedList, error) {
	return s.repo.List(ctx, filter)
}
