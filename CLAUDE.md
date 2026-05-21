# SubClub — Claude Code Guide

## Common Commands

```bash
make build      # compile to bin/subclub
make test       # go test -v ./...
make run        # run locally with APP_ENV=development
make dev        # docker compose up (hot-reload via air + dlv debug)
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
| Domain | `internal/domain` | Interfaces, models, domain errors. No SQL, no JSON, no frameworks. |
| Application | `internal/application` | Services (use-case orchestration) and repository implementations (Ent ORM). |
| Web | `internal/web` | HTTP handlers, DTOs, `Handlers` struct that wires services to handlers. |
| Infrastructure | `internal/infra` | `Server` (Echo + middleware + routes), database connection, logger. |

### HTTP wiring

`Server` owns the full HTTP setup: it holds a `logger *slog.Logger` and an internal `router` struct as fields.
- `server.go` — Echo init, middleware, health route, calls `router.registerRoutes()`
- `router.go` — unexported `router` struct, registers all business routes
- `middleware/logger.go` — receives `*slog.Logger`, does not use `slog.Default()`

Initialization in `cmd/subclub/application.go`:
```go
h   := web.NewHandlers(app)
srv := server.NewServer(h)
srv.Start(cfg.Port)
```

## Key Conventions

### Client vs Customer
`Client` is the domain name used everywhere in Go code and API routes (`/clients`).  
`Customer` is the physical name in the database and Ent schema — required to avoid collision with `ent.Client` (the ORM connection type).  
Mapping happens in `internal/application/repository/client/`.

### Ent ORM
- Never write raw SQL. All queries go through Ent.
- Schemas are defined in `ent/schema/`. After any change, run `go generate ./ent`.
- The `sqlx.DB` connection in `internal/infra/database` exists solely for TCP connection pool management, not for queries.

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
- Use `faker` helpers (`faker.FakeUser()`, `faker.FakeClient()`, etc.) instead of hardcoded strings.
- Do not mock the database in integration tests — use a real connection.

### Fake Data Seeder
Runs automatically on startup when `APP_ENV` is `development`, `dev`, or `local`, and only if the database is empty. Seeds 1 admin, 10 products, 4 plans, 50 customers, 25 subscriptions.
