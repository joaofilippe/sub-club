# SubClub — Claude Code Guide

## Claude Behavior

Atue como um **Dev Sênior Fullstack**: opine sobre design, aponte trade-offs, sugira melhorias, mas implemente apenas o que for pedido — sem abstrações prematuras nem escopo extra.

### Workflow obrigatório para qualquer mudança de código

1. **Criar branch** antes de qualquer alteração:
   ```bash
   git checkout -b <tipo>/<descricao-curta>
   # ex: feat/customer-filter, fix/auth-middleware, refactor/tenant-ctx
   ```
2. **Commitar** ao fim de cada tarefa concluída, com mensagem em inglês seguindo Conventional Commits (`feat:`, `fix:`, `refactor:`, `docs:`, `chore:`).
3. **Não fazer push nem abrir PR** sem o usuário pedir explicitamente.

---

## Common Commands

```bash
make build      # compile to bin/subclub
make test       # go test -v ./...
make run        # run locally with APP_ENV=development
make dev        # docker compose up -d db pgadmin && docker compose up app (hot-reload via air + dlv debug)
make swagger    # regenerate OpenAPI docs from handler annotations
make lint       # golangci-lint run
make tidy       # go mod tidy
go generate ./ent  # regenerate Ent ORM code after schema changes
```

## Architecture

Four layers — dependencies always point inward, outer layers never imported by inner ones:

```
domain → application → web → infra
```

| Layer | Package | Responsibility |
|---|---|---|
| Domain | `internal/domain` | Models, repository/service interfaces, use cases (`usecase/`). No SQL, no JSON, no frameworks. |
| Application | `internal/application` | `service/` (use-case orchestration) and `repository/` (Ent ORM implementations). |
| Web | `internal/web` | HTTP handlers, DTOs, `Handlers` struct that wires services to handlers. |
| Infrastructure | `internal/infra` | `server/` (Echo + middleware + routes), `database/` (connection, migrations, seeder, TenantClientManager), `middleware/`, `authctx/`, `tenantctx/`. |

### Domain structure per entity

Each business entity under `internal/domain/<entity>/` follows this layout:

```
model/          → structs, typed inputs, domain errors
repository.go   → persistence interface
service.go      → service interface
usecase/        → use-case implementations (create, list, get_by_id, update, delete)
```

### HTTP wiring

`Server` owns the full HTTP setup: it holds a `logger *slog.Logger` and an internal `router` struct as fields.
- `server.go` — Echo init, middleware, health route, calls `router.registerRoutes()`
- `router.go` — unexported `router` struct, registers all business routes under two groups (see Multi-tenant below)
- `middleware/logger.go` — receives `*slog.Logger`, does not use `slog.Default()`

Initialization in `cmd/subclub/application.go`:
```go
h   := web.NewHandlers(app)
srv := server.NewServer(h)
srv.Start(cfg.Port)
```

## Multi-tenant Architecture

SubClub is a B2B platform: each tenant is an `Account`. Data is isolated per tenant using **PostgreSQL schema-per-tenant** (`account_{slug}`).

| Schema | Tables |
|--------|--------|
| `public` | `system_users`, `accounts`, `account_plans` |
| `account_{slug}` | `users`, `customers`, `plans`, `products`, `subscriptions` |

### Users: SystemUser vs User

| Type | Ent entity | Schema | Purpose |
|---|---|---|---|
| `SystemUser` | `ent.SystemUser` | `public` | SubClub platform admins |
| `User` | `ent.User` | `account_{slug}` | Tenant-level admins/operators |

### Route groups and middleware

| Group | Middleware | Routes |
|-------|-----------|--------|
| Admin | `RequireAdminMiddleware` — validates JWT, requires role `admin` | `/accounts`, `/account-plans`, `/users` |
| Tenant | `AuthMiddleware` — validates JWT, resolves tenant schema, injects context | `/customers`, `/plans`, `/products`, `/subscriptions`, `/modules` |

### Request context

Handlers in the tenant group read two injected values from the Echo context:

```go
client := tenantctx.TenantClientFromContext(ctx)   // *ent.Client scoped to tenant schema
claims, ok := authctx.ClaimsFromContext(ctx)        // UserID, AccountSlug, Role
```

### TenantClientManager

Located in `internal/infra/database/tenant_manager.go`. Caches one `*ent.Client` per slug (backed by a `*sql.DB` with `search_path = "account_{slug}", public`). On first request per tenant it opens the connection; subsequent requests hit the cache.

### AccountPlan

`AccountPlan` is SubClub's own subscription plan for its tenants — not to be confused with the `Plan` entities tenants sell to their customers. It defines `max_customers`, `max_plans`, `max_products`, and `price`.

## Key Conventions

### Ent ORM
- Never write raw SQL. All queries go through Ent.
- Schemas are defined in `ent/schema/`. After any change, run `go generate ./ent`.
- The `*sql.DB` connection in `internal/infra/database` exists solely for TCP connection pool management and to seed the Ent driver.

### Soft Delete
Entities that need logical deletion use a nullable `deleted_at` field:
```go
field.Time("deleted_at").Optional().Nillable()
```
Queries must always filter with `.Where(<entity>.DeletedAtIsNil())`.

### Swagger Annotations
All handler methods must have Swag annotations (`// @Summary`, `// @Router`, etc.).
Regenerate docs with `make swagger` after any change.

### Tests
- Use `faker` helpers (`faker.FakeUser()`, `faker.FakeCustomer()`, etc.) instead of hardcoded strings.
- Do not mock the database in integration tests — use a real connection.

### Fake Data Seeder
Runs automatically on startup when `APP_ENV=development` and the public schema is empty. Seeds two phases:

**Public schema** (`SeedAll`): 1 SystemUser admin (`adm@adm.com` / `12345678`), 1 Demo AccountPlan, 1 Demo Account (slug `demo`).

**Tenant schema** (`SeedTenant`, triggered by account creation): 1 tenant User (`admin@demo.com` / `12345678`), 10 fixed coffee products, 3 fixed plans (Básico / Intermediário / Avançado), 50 fake customers, 25 active subscriptions.
