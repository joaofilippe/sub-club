# Arquitetura Multi-tenant

Este documento descreve como o SubClub isola os dados de cada tenant, como um novo tenant é provisionado e como o roteamento de contexto funciona em tempo de requisição.

---

## Modelo Conceitual

O SubClub é uma plataforma B2B: empresas que operam clubes de assinatura são os **tenants** (chamados de `Account`). Cada Account possui seus próprios clientes, planos e produtos, completamente isolados dos demais.

```
SubClub (plataforma)
├── Account "café-do-norte"   ← tenant A
│   ├── Customers
│   ├── Plans
│   ├── Products
│   └── Subscriptions
└── Account "vinhos-sul"      ← tenant B
    ├── Customers
    ├── Plans
    ├── Products
    └── Subscriptions
```

---

## Estratégia de Isolamento: Schema-per-Tenant

O isolamento é feito no banco de dados por meio de **schemas PostgreSQL dedicados**. Quando uma Account é criada, a infra provisiona automaticamente um schema `account_{slug}`:

| Schema | Tabelas |
|--------|---------|
| `public` | `accounts`, `account_plans`, `users` (globais, compartilhadas) |
| `account_café_do_norte` | `customers`, `plans`, `products`, `subscriptions` |
| `account_vinhos_sul` | `customers`, `plans`, `products`, `subscriptions` |

O roteamento para o schema correto acontece via `search_path` do PostgreSQL. Cada tenant recebe um `*sql.DB` exclusivo, configurado com:

```
search_path = "account_{slug}", public
```

Com esse `search_path`, as queries do Ent lêem as tabelas tenant do schema isolado e ainda enxergam as tabelas globais (como `account_plans`) a partir de `public`, sem necessitar de prefixação explícita.

---

## AccountPlan — Planos do SubClub para seus Tenants

`AccountPlan` é o plano de assinatura do SubClub contratado por cada Account (não confundir com os `Plan`s que cada tenant oferece aos seus próprios clientes). Ele define os limites operacionais da conta:

| Campo | Descrição |
|-------|-----------|
| `max_customers` | Limite de clientes que o tenant pode cadastrar |
| `max_plans` | Limite de planos que o tenant pode criar |
| `max_products` | Limite de produtos que o tenant pode cadastrar |
| `price` | Preço mensal cobrado pelo SubClub |

Uma `Account` referencia seu `AccountPlan` via `account_plan_id`. O campo é opcional para permitir contas em período de avaliação sem plano vinculado.

---

## Ciclo de Vida de um Tenant

### 1. Criação da Account

`POST /accounts` (rota admin) → `AccountService.Create`:

```
CreateAccountInput → usecase → Account salva em public.accounts
                             ↓
                    schemaCreator.CreateTenantSchema(ctx, slug)
                             ↓
               CREATE SCHEMA IF NOT EXISTS "account_{slug}"
                             ↓
               Ent auto-migration no novo schema
               (cria customers, plans, products, subscriptions)
```

A interface `account.SchemaCreator` (definida na camada de domínio em [internal/domain/account/schema.go](../../internal/domain/account/schema.go)) é implementada pelo `TenantClientManager` na camada de infra. Isso garante que o domínio não dependa de detalhes de banco de dados.

### 2. Status de Assinatura

Uma Account pode estar em um dos seguintes estados:

| Status | Acesso permitido? |
|--------|:-----------------:|
| `trial` | Sim |
| `active` | Sim |
| `suspended` | Não |
| `cancelled` | Não |

O campo `active` (boolean) complementa o status: uma conta com `active = false` é sempre bloqueada, independente do status.

---

## Fluxo de Requisição Autenticada

Toda requisição a rotas tenant passa pelo `AuthMiddleware`:

```
Request com JWT
      │
      ▼
AuthMiddleware
  1. Extrai e valida o JWT (HMAC)
  2. Lê claim account_slug do token
  3. IsAccountAccessible(slug) → verifica status no banco global
  4. TenantClientManager.GetOrCreate(slug) → retorna *ent.Client do tenant
  5. Injeta o client no contexto via tenantctx.WithTenantClient(...)
  6. Injeta as claims (UserID, AccountSlug, Role) via authctx.WithClaims(...)
      │
      ▼
Handler
  - tenantctx.TenantClientFromContext(ctx) → *ent.Client isolado
  - authctx.ClaimsFromContext(ctx) → dados do usuário autenticado
```

### TenantClientManager

Localizado em [internal/infra/database/tenant_manager.go](../../internal/infra/database/tenant_manager.go), mantém um cache `map[slug]*ent.Client` protegido por `sync.RWMutex` com double-checked locking na criação. O custo de conexão ao banco ocorre apenas na primeira requisição de cada tenant.

---

## Separação de Rotas

As rotas estão divididas em dois grupos com middlewares distintos:

| Grupo | Middleware | Rotas |
|-------|-----------|-------|
| Admin | `RequireAdminMiddleware` | `/accounts`, `/account-plans`, `/users` |
| Tenant | `AuthMiddleware` | `/customers`, `/plans`, `/products`, `/subscriptions` |

Rotas admin validam apenas o JWT e exigem role `admin`. Rotas tenant, além da autenticação, resolvem o tenant e injetam o `*ent.Client` isolado.

---

## Contextos de Injeção

O sistema usa dois pacotes de contexto distintos para não misturar responsabilidades:

| Pacote | Tipo injetado | Propósito |
|--------|--------------|-----------|
| `internal/infra/tenantctx` | `*ent.Client` | Client Ent apontando para o schema do tenant |
| `internal/infra/authctx` | `authctx.Claims` | Dados do usuário autenticado (UserID, AccountSlug, Role) |

Os handlers recuperam esses valores com:

```go
client := tenantctx.TenantClientFromContext(ctx)  // nunca nil em rotas tenant
claims, ok := authctx.ClaimsFromContext(ctx)
```

---

## Adicionando Tabelas a um Tenant

Para que uma nova tabela faça parte do schema isolado de cada tenant (e não do schema global `public`):

1. Crie o schema Ent em `ent/schema/`.
2. Rode `go generate ./ent`.
3. **Não adicione** a nova entidade a nenhuma edge que parte de `Account` ou `AccountPlan` (essas ficam em `public`).
4. O próximo `CreateTenantSchema` para qualquer conta nova criará a tabela automaticamente. Para contas existentes, será necessário rodar uma migração manual ou um script de provisionamento.
