package client

import (
	"context"
)

// Service defines the interface for client-related business logic
type Service interface {
	Create(ctx context.Context, client *Client) error
	GetByID(ctx context.Context, id string) (*Client, error)
	Update(ctx context.Context, client *Client) error
	Delete(ctx context.Context, id string) error
	List(ctx context.Context, filter Filter) (*PaginatedList, error)
}
