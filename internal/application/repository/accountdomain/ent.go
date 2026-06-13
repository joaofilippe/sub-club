package repository

import (
	"context"

	"github.com/joaofilippe/subclub/ent"
	entaccountdomain "github.com/joaofilippe/subclub/ent/accountdomain"
	"github.com/joaofilippe/subclub/internal/domain/accountdomain"
)

type AccountDomainEntRepository struct {
	client *ent.Client
}

func NewAccountDomainEntRepository(client *ent.Client) accountdomain.Repository {
	return &AccountDomainEntRepository{client: client}
}

func (r *AccountDomainEntRepository) GetByDomain(ctx context.Context, domain string) (*accountdomain.Info, error) {
	d, err := r.client.AccountDomain.Query().
		Where(entaccountdomain.DomainEQ(domain)).
		WithAccount().
		Only(ctx)
	if err != nil {
		return nil, err
	}
	return &accountdomain.Info{
		Slug: d.Edges.Account.Slug,
		Name: d.Edges.Account.Name,
	}, nil
}
