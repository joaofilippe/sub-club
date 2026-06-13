package accountdomain

import "context"

type Info struct {
	Slug string
	Name string
}

type Repository interface {
	GetByDomain(ctx context.Context, domain string) (*Info, error)
}
