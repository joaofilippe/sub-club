package client

import (
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
