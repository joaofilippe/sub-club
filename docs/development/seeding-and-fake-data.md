# Seeding and Fake Data

This document explains how the application handles initial data population and how to use the fake data generation utilities.

## Overview

To facilitate development and testing, the application includes a mechanism to automatically populate the database with realistic fake data. This is particularly useful for frontend development and manual testing of pagination, filters, and complex scenarios.

## Fake Data Utility

We use the [gofakeit](https://github.com/brianvoe/gofakeit) library to generate realistic data. A central utility is located at `internal/test/faker/faker.go`.

### Available Helpers

- `faker.FakeUser()`: Returns a `model.User` with a random name, email, and password.
- `faker.FakeClient()`: Returns a `model.Client` with a random name, CPF-like document, phone, and full address.
- `faker.FakeProduct()`: Returns a `model.Product` with a random coffee-related name, code, and price.
- `faker.FakePlan()`: Returns a `model.Plan` with calculated product value and discounts.

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

### Execution Rules

The seeder only runs if:

1. The database is empty (checked by counting rows in the `users` table).
2. The environment variable `APP_ENV` is set to `development`.

### Seeding Configuration

Currently, the seeder is configured to generate:

- **1 Admin**: `adm@adm.com` / `12345678`
- **10 Products**
- **4 Plans**
- **50 Customers**
- **25 Active Subscriptions**

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
