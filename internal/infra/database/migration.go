package database

import (
	"context"

	"github.com/joaofilippe/subclub/ent"
	entmigrate "github.com/joaofilippe/subclub/ent/migrate"
	entsqlschema "entgo.io/ent/dialect/sql/schema"
)

// PublicTables are the tables that live in the public schema (platform-level).
var PublicTables = []*entsqlschema.Table{
	entmigrate.AccountsTable,
	entmigrate.AccountPlansTable,
	entmigrate.AccountDomainsTable,
	entmigrate.SystemUsersTable,
	entmigrate.ModulesTable,
	entmigrate.AccountPlanModulesTable,
}

// TenantTables are the tables that live in each tenant schema (account_{slug}).
var TenantTables = []*entsqlschema.Table{
	entmigrate.UsersTable,
	entmigrate.CustomersTable,
	entmigrate.PlansTable,
	entmigrate.ProductsTable,
	entmigrate.SubscriptionsTable,
}

// MigratePublic creates or updates only the public-schema tables using the global client.
func MigratePublic(ctx context.Context, client *ent.Client) error {
	return entmigrate.Create(ctx, client.Schema, PublicTables)
}

// MigrateTenant creates or updates only the tenant-schema tables using a tenant client
// whose search_path is already set to "account_{slug}, public".
func MigrateTenant(ctx context.Context, client *ent.Client) error {
	return entmigrate.Create(ctx, client.Schema, TenantTables)
}
