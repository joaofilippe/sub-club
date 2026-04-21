package client

import (
	"context"
	"time"
)

type Address struct {
	ZipCode      string
	Street       string
	Number       string
	Complement   string
	Neighborhood string
	City         string
	State        string
}

type Client struct {
	ID        string
	Name      string
	Email     string
	Phone     string
	Document  string
	Active    bool
	Address   *Address
	CreatedAt time.Time
}

type Filter struct {
	Search   *string
	IsActive *bool
	Page     int
	PageSize int
}

type PaginatedList struct {
	Items      []*Client
	TotalCount int
}

// Repository defines the interface for interacting with client storage
type Repository interface {
	Create(ctx context.Context, client *Client) error
	GetByID(ctx context.Context, id string) (*Client, error)
	Update(ctx context.Context, client *Client) error
	Delete(ctx context.Context, id string) error
	List(ctx context.Context, filter Filter) (*PaginatedList, error)
}
