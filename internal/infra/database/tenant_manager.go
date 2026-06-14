package database

import (
	"context"
	"database/sql"
	"fmt"
	"sync"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/stdlib"
	"github.com/joaofilippe/subclub/ent"
	"github.com/joaofilippe/subclub/ent/account"
	entsql "entgo.io/ent/dialect/sql"
	"golang.org/x/crypto/bcrypt"
)

var (
	ErrAccountNotFound  = fmt.Errorf("account not found")
	ErrAccountSuspended = fmt.Errorf("account is suspended or cancelled")
)

// TenantClientManager caches one *ent.Client per account slug.
// Each client is backed by a *sql.DB whose search_path is permanently set to
// "account_{slug}", public — so every query from that client goes to the
// correct tenant schema without needing per-request transactions.
type TenantClientManager struct {
	baseURL      string
	isDev        bool
	mu           sync.RWMutex
	clients      map[string]*ent.Client
	globalDB     *sql.DB
	globalClient *ent.Client
}

func NewTenantClientManager(baseURL string, globalDB *sql.DB, globalClient *ent.Client, isDev bool) *TenantClientManager {
	return &TenantClientManager{
		baseURL:      baseURL,
		isDev:        isDev,
		globalDB:     globalDB,
		globalClient: globalClient,
		clients:      make(map[string]*ent.Client),
	}
}

// IsAccountAccessible returns nil if the account exists and its subscription
// status allows access (trial or active). Returns ErrAccountSuspended or
// ErrAccountNotFound otherwise.
func (m *TenantClientManager) IsAccountAccessible(ctx context.Context, slug string) error {
	acc, err := m.globalClient.Account.Query().
		Where(account.SlugEQ(slug)).
		Select(account.FieldSubscriptionStatus, account.FieldActive).
		First(ctx)
	if err != nil {
		return ErrAccountNotFound
	}
	if !acc.Active {
		return ErrAccountSuspended
	}
	switch acc.SubscriptionStatus {
	case "active", "trial":
		return nil
	default:
		return ErrAccountSuspended
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

	// 3. Run migration for tenant tables only.
	// The tenant client's search_path is "account_{slug}, public", so tables are
	// created inside account_{slug}. Public tables are intentionally excluded.
	if err := MigrateTenant(ctx, tenantClient); err != nil {
		return fmt.Errorf("migrate tenant schema %q: %w", schemaName, err)
	}

	// 4. Seed demo data on development environments.
	if m.isDev {
		SeedTenant(ctx, tenantClient, slug)
	}

	return nil
}

// CreateTenantOwner creates the first admin User inside the tenant schema.
// The password is hashed with bcrypt before storing.
func (m *TenantClientManager) CreateTenantOwner(ctx context.Context, slug, email, name, username, password string) error {
	tenantClient, err := m.GetOrCreate(slug)
	if err != nil {
		return fmt.Errorf("get tenant client for %q: %w", slug, err)
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return fmt.Errorf("hash password: %w", err)
	}

	_, err = tenantClient.User.Create().
		SetID(uuid.New()).
		SetName(name).
		SetUsername(username).
		SetEmail(email).
		SetPassword(string(hash)).
		SetRole("admin").
		Save(ctx)
	return err
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
