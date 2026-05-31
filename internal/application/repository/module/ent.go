package repository

import (
	"context"

	"github.com/google/uuid"
	"github.com/joaofilippe/subclub/ent"
	entmodule "github.com/joaofilippe/subclub/ent/module"
	"github.com/joaofilippe/subclub/internal/domain/module"
	"github.com/joaofilippe/subclub/internal/domain/module/model"
)

type ModuleEntRepository struct {
	client *ent.Client
}

func NewModuleEntRepository(client *ent.Client) module.Repository {
	return &ModuleEntRepository{client: client}
}

func (r *ModuleEntRepository) Create(ctx context.Context, m *model.Module) error {
	id, err := uuid.Parse(m.ID)
	if err != nil {
		id = uuid.New()
	}
	_, err = r.client.Module.Create().
		SetID(id).
		SetName(m.Name).
		SetActive(m.Active).
		Save(ctx)
	return err
}

func (r *ModuleEntRepository) GetByID(ctx context.Context, id string) (*model.Module, error) {
	parsedID, err := uuid.Parse(id)
	if err != nil {
		return nil, err
	}
	m, err := r.client.Module.Get(ctx, parsedID)
	if err != nil {
		return nil, err
	}
	return mapEntModuleToDomain(m), nil
}

func (r *ModuleEntRepository) Update(ctx context.Context, m *model.Module) error {
	parsedID, err := uuid.Parse(m.ID)
	if err != nil {
		return err
	}
	_, err = r.client.Module.UpdateOneID(parsedID).
		SetName(m.Name).
		SetActive(m.Active).
		Save(ctx)
	return err
}

func (r *ModuleEntRepository) Delete(ctx context.Context, id string) error {
	parsedID, err := uuid.Parse(id)
	if err != nil {
		return err
	}
	_, err = r.client.Module.UpdateOneID(parsedID).SetActive(false).Save(ctx)
	return err
}

func (r *ModuleEntRepository) List(ctx context.Context) ([]*model.Module, error) {
	modules, err := r.client.Module.Query().
		Where(entmodule.ActiveEQ(true)).
		All(ctx)
	if err != nil {
		return nil, err
	}
	results := make([]*model.Module, 0, len(modules))
	for _, m := range modules {
		results = append(results, mapEntModuleToDomain(m))
	}
	return results, nil
}

func mapEntModuleToDomain(m *ent.Module) *model.Module {
	return &model.Module{
		ID:        m.ID.String(),
		Name:      m.Name,
		Active:    m.Active,
		CreatedAt: m.CreatedAt,
	}
}
