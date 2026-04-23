package repository

import (
	"context"

	"github.com/google/uuid"
	"github.com/joaofilippe/subclub/ent"
	"github.com/joaofilippe/subclub/ent/customer"
	"github.com/joaofilippe/subclub/ent/schema"
	"github.com/joaofilippe/subclub/internal/domain/client"
	"github.com/joaofilippe/subclub/internal/domain/client/model"
)

type ClientEntRepository struct {
	entClient *ent.Client
}

func NewClientEntRepository(entClient *ent.Client) client.Repository {
	return &ClientEntRepository{entClient: entClient}
}

func (r *ClientEntRepository) Create(ctx context.Context, c *model.Client) error {
	id, err := uuid.Parse(c.ID)
	if err != nil {
		id = uuid.New()
	}

	var address *schema.Address
	if c.Address != nil {
		address = &schema.Address{
			ZipCode:      c.Address.ZipCode,
			Street:       c.Address.Street,
			Number:       c.Address.Number,
			Complement:   c.Address.Complement,
			Neighborhood: c.Address.Neighborhood,
			City:         c.Address.City,
			State:        c.Address.State,
		}
	}

	builder := r.entClient.Customer.Create().
		SetID(id).
		SetName(c.Name).
		SetEmail(c.Email).
		SetNillablePhone(&c.Phone).
		SetNillableDocument(&c.Document).
		SetActive(c.Active)

	if address != nil {
		builder.SetAddress(address)
	}
	_, err = builder.Save(ctx)
	return err
}

func (r *ClientEntRepository) GetByID(ctx context.Context, id string) (*model.Client, error) {
	parsedID, err := uuid.Parse(id)
	if err != nil {
		return nil, err
	}
	c, err := r.entClient.Customer.Get(ctx, parsedID)
	if err != nil {
		return nil, err
	}
	return mapEntCustomerToDomain(c), nil
}

func (r *ClientEntRepository) Update(ctx context.Context, c *model.Client) error {
	parsedID, err := uuid.Parse(c.ID)
	if err != nil {
		return err
	}

	var address *schema.Address
	if c.Address != nil {
		address = &schema.Address{
			ZipCode:      c.Address.ZipCode,
			Street:       c.Address.Street,
			Number:       c.Address.Number,
			Complement:   c.Address.Complement,
			Neighborhood: c.Address.Neighborhood,
			City:         c.Address.City,
			State:        c.Address.State,
		}
	}

	builder := r.entClient.Customer.UpdateOneID(parsedID).
		SetName(c.Name).
		SetEmail(c.Email).
		SetNillablePhone(&c.Phone).
		SetNillableDocument(&c.Document).
		SetActive(c.Active)

	if address != nil {
		builder.SetAddress(address)
	} else {
		builder.ClearAddress()
	}

	_, err = builder.Save(ctx)
	return err
}

func (r *ClientEntRepository) Delete(ctx context.Context, id string) error {
	parsedID, err := uuid.Parse(id)
	if err != nil {
		return err
	}
	_, err = r.entClient.Customer.UpdateOneID(parsedID).SetActive(false).Save(ctx)
	return err
}

func (r *ClientEntRepository) List(ctx context.Context, filter model.Filter) (*model.PaginatedList, error) {
	q := r.entClient.Customer.Query()

	if filter.IsActive != nil {
		q = q.Where(customer.ActiveEQ(*filter.IsActive))
	}
	if filter.Search != nil && *filter.Search != "" {
		q = q.Where(
			customer.Or(
				customer.NameContainsFold(*filter.Search),
				customer.EmailContainsFold(*filter.Search),
			),
		)
	}

	totalCount, err := q.Count(ctx)
	if err != nil {
		return nil, err
	}

	if filter.Page > 0 && filter.PageSize > 0 {
		q = q.Limit(filter.PageSize).Offset((filter.Page - 1) * filter.PageSize)
	}

	entCustomers, err := q.All(ctx)
	if err != nil {
		return nil, err
	}

	var results []*model.Client
	for _, ec := range entCustomers {
		results = append(results, mapEntCustomerToDomain(ec))
	}
	return &model.PaginatedList{
		Items:      results,
		TotalCount: totalCount,
	}, nil
}

func mapEntCustomerToDomain(ec *ent.Customer) *model.Client {
	var address *model.Address
	if ec.Address != nil {
		address = &model.Address{
			ZipCode:      ec.Address.ZipCode,
			Street:       ec.Address.Street,
			Number:       ec.Address.Number,
			Complement:   ec.Address.Complement,
			Neighborhood: ec.Address.Neighborhood,
			City:         ec.Address.City,
			State:        ec.Address.State,
		}
	}

	return &model.Client{
		ID:        ec.ID.String(),
		Name:      ec.Name,
		Email:     ec.Email,
		Phone:     ec.Phone,
		Document:  ec.Document,
		Active:    ec.Active,
		Address:   address,
		CreatedAt: ec.CreatedAt,
	}
}
