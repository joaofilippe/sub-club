package model

import "time"

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
