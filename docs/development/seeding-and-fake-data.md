# Seeding and Fake Data

This document explains how the application handles initial data population and how to use the fake data generation utilities.

## Overview

To facilitate development and testing, the application includes a mechanism to automatically populate the database with realistic fake data. This is particularly useful for frontend development and manual testing of pagination, filters, and complex scenarios.

The seeder follows the multi-tenant architecture: it first seeds the **public schema** (system-level entities), then provisions and seeds a **demo tenant schema**.

## Fake Data Utility

We use the [gofakeit](https://github.com/brianvoe/gofakeit) library to generate realistic data. A central utility is located at `internal/test/faker/faker.go`.

### Available Helpers

| Helper | Returns | Description |
|---|---|---|
| `faker.FakeUser()` | `*usermodel.User` | Random name, email, and password for a tenant user. |
| `faker.FakeCustomer()` | `*customermodel.Customer` | Random name, document, phone, email, and full address. |
| `faker.FakeProduct()` | `*productmodel.Product` | Random coffee-related name, code, and price. |
| `faker.FakePlan()` | `*planmodel.Plan` | Random plan with calculated product value and discounts. |
| `faker.FakeAccount()` | `*accountmodel.Account` | Random company name, email, document, and slug. |
| `faker.FakeAccountPlan()` | `*accountplanmodel.AccountPlan` | Random SubClub subscription plan with limits. |
| `faker.FakeModule()` | `*modulemodel.Module` | Random module name. |
| `faker.CoffeeProducts()` | `[]*productmodel.Product` | 10 fixed coffee products used for seeding. |
| `faker.FixedPlans()` | `[]*planmodel.Plan` | 3 fixed subscription tiers (Básico, Intermediário, Avançado). |

### How to use in Tests

You should use these helpers instead of hardcoding strings in your tests. This makes tests more robust and less prone to collisions.

```go
func TestSomething(t *testing.T) {
    u := faker.FakeUser()
    // Use u in your test...
}
```

## Database Seeder

The seeder logic resides in `internal/infra/database/seeder.go`. It is called during application startup (see `internal/application/application.go`).

It exposes two functions:

- **`SeedAll`** — seeds the public schema; called on startup.
- **`SeedTenant`** — seeds a newly provisioned tenant schema; called automatically when a new Account is created in development.

### Execution Rules

`SeedAll` only runs if:

1. The environment variable `APP_ENV` is set to `development`.
2. The public schema is empty (checked by counting rows in the `system_users` table).

### What Gets Seeded

#### Public Schema (`SeedAll`)

| Entity | Credentials / Details |
|---|---|
| **SystemUser** (admin) | `adm@adm.com` / `12345678` |
| **AccountPlan** | Name: `Demo` — free, up to 100 customers, 5 plans, 20 products |
| **Account** | Name: `Demo`, slug: `demo`, email: `demo@demo.com` |

After the Account is created, `CreateTenantSchema` is called, which provisions the `demo` schema and invokes `SeedTenant`.

#### Tenant Schema (`SeedTenant`, slug `demo`)

| Entity | Details |
|---|---|
| **User** (tenant admin) | `admin@demo.com` / `12345678` |
| **Products** | 10 fixed coffee products (e.g. Espresso, Cappuccino, Cold Brew…) |
| **Plans** | 3 tiers: Básico (R$ 44,90), Intermediário (R$ 79,90), Avançado (R$ 129,90) |
| **Customers** | 50 fake customers |
| **Subscriptions** | 25 active subscriptions (first 25 customers, round-robin across plans) |

To change these quantities, modify the loops in `internal/infra/database/seeder.go`.

## Environment Setup

Ensure your local environment sets the `APP_ENV` variable.

### Docker Compose

The `docker-compose.yml` is already configured with:

```yaml
environment:
  - APP_ENV=development
```

### Makefile

The `make run` command is also updated:

```makefile
run:
    APP_ENV=development $(GOCMD) run $(MAIN_PATH)/*.go server
```
