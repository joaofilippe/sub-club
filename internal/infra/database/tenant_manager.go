package database

import (
	"context"
	"database/sql"
	"fmt"
	"sync"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/stdlib"
	"github.com/joaofilippe/subclub/ent"
	entsql "entgo.io/ent/dialect/sql"
)

// TenantClientManager caches one *ent.Client per account slug.
// Each client is backed by a *sql.DB whose search_path is permanently set to
// "account_{slug}", public — so every query from that client goes to the
// correct tenant schema without needing per-request transactions.
type TenantClientManager struct {
	baseURL string
	mu      sync.RWMutex
	clients map[string]*ent.Client
	// globalDB is used for DDL operations (CREATE SCHEMA).
	globalDB *sql.DB
}

func NewTenantClientManager(baseURL string, globalDB *sql.DB) *TenantClientManager {
	return &TenantClientManager{
		baseURL:  baseURL,
		globalDB: globalDB,
		clients:  make(map[string]*ent.Client),
	}
}

// GetOrCreate returns the cached tenant client or creates a new one.
func (m *TenantClientManager) GetOrCreate(slug string) (*ent.Client, error) {
	m.mu.RLock()
	c, ok := m.clients[slug]
	m.mu.RUnlock()
	if ok {
		return c, nil
	}

	m.mu.Lock()
	defer m.mu.Unlock()

	// Double-check after acquiring write lock.
	if c, ok = m.clients[slug]; ok {
		return c, nil
	}

	c, err := m.newTenantClient(slug)
	if err != nil {
		return nil, err
	}
	m.clients[slug] = c
	return c, nil
}

// CreateTenantSchema creates the PostgreSQL schema for a new account and
// runs Ent auto-migration so that tenant tables are created inside it.
func (m *TenantClientManager) CreateTenantSchema(ctx context.Context, slug string) error {
	schemaName := tenantSchemaName(slug)

	// 1. Create the PostgreSQL schema.
	if _, err := m.globalDB.ExecContext(ctx,
		fmt.Sprintf(`CREATE SCHEMA IF NOT EXISTS %q`, schemaName),
	); err != nil {
		return fmt.Errorf("create schema %q: %w", schemaName, err)
	}

	// 2. Get (or create) the tenant-scoped Ent client.
	tenantClient, err := m.GetOrCreate(slug)
	if err != nil {
		return fmt.Errorf("get tenant client for %q: %w", slug, err)
	}

	// 3. Run Ent auto-migration.
	// With search_path = account_{slug}, public:
	//   - global tables (accounts, account_plans, users) are already present in public → skipped.
	//   - tenant tables (customers, plans, products, subscriptions) are absent → created in account_{slug}.
	if err := tenantClient.Schema.Create(ctx); err != nil {
		return fmt.Errorf("migrate tenant schema %q: %w", schemaName, err)
	}

	return nil
}

func (m *TenantClientManager) newTenantClient(slug string) (*ent.Client, error) {
	cfg, err := pgx.ParseConfig(m.baseURL)
	if err != nil {
		return nil, fmt.Errorf("parse database URL: %w", err)
	}

	// Set search_path so every connection in this pool targets the tenant schema first.
	cfg.RuntimeParams["search_path"] = fmt.Sprintf(`"%s", public`, tenantSchemaName(slug))

	connStr := stdlib.RegisterConnConfig(cfg)
	db, err := sql.Open("pgx", connStr)
	if err != nil {
		return nil, fmt.Errorf("open tenant db for %q: %w", slug, err)
	}
	db.SetMaxOpenConns(3)
	db.SetMaxIdleConns(1)

	drv := entsql.OpenDB("postgres", db)
	return ent.NewClient(ent.Driver(drv)), nil
}

func tenantSchemaName(slug string) string {
	return "account_" + slug
}
