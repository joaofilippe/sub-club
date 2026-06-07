package account

import "context"

// SchemaCreator is implemented by the infra layer and called by the service
// to provision a dedicated PostgreSQL schema whenever a new account is created.
type SchemaCreator interface {
	CreateTenantSchema(ctx context.Context, slug string) error
	CreateTenantOwner(ctx context.Context, slug, email, name, password string) error
}
