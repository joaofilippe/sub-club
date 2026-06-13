# refactor/api-v1-prefix

## O que mudou e por quê

Todas as rotas da API foram prefixadas com `/api/v1`. O objetivo é seguir a convenção REST de versionamento de API e deixar espaço para uma `/api/v2` no futuro sem breaking changes.

## Contrato quebrado — migração obrigatória no front

**Todas** as chamadas HTTP precisam ser atualizadas. O padrão é:

```
ANTES:  /auth/login
DEPOIS: /api/v1/auth/login
```

## Rotas afetadas

| Método | Antes | Depois |
|--------|-------|--------|
| POST | `/auth/login` | `/api/v1/auth/login` |
| POST/GET | `/account-plans` | `/api/v1/account-plans` |
| GET/PUT/DELETE | `/account-plans/:id` | `/api/v1/account-plans/:id` |
| POST/GET | `/accounts` | `/api/v1/accounts` |
| GET/PUT/DELETE | `/accounts/:id` | `/api/v1/accounts/:id` |
| POST/GET | `/modules` | `/api/v1/modules` |
| GET/PUT/DELETE | `/modules/:id` | `/api/v1/modules/:id` |
| POST | `/system-users` | `/api/v1/system-users` |
| POST/GET | `/users` | `/api/v1/users` |
| GET/PUT/DELETE | `/users/:id` | `/api/v1/users/:id` |
| POST/GET | `/customers` | `/api/v1/customers` |
| GET/PUT/DELETE | `/customers/:id` | `/api/v1/customers/:id` |
| POST/GET | `/plans` | `/api/v1/plans` |
| GET/PUT/DELETE | `/plans/:id` | `/api/v1/plans/:id` |
| POST/GET | `/subscriptions` | `/api/v1/subscriptions` |
| GET/PUT/DELETE | `/subscriptions/:id` | `/api/v1/subscriptions/:id` |
| POST/GET | `/products` | `/api/v1/products` |
| GET/PUT/DELETE | `/products/:id` | `/api/v1/products/:id` |

## Rotas sem alteração

- `GET /health` — sem prefixo (endpoint de infraestrutura)
- `GET /swagger/*` — sem prefixo
