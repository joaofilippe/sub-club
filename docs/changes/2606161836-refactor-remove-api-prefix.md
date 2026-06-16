# refactor/remove-api-prefix

## O que mudou e por quê

Removido o segmento `/api` do prefixo de todas as rotas da API. O prefixo anterior `/api/v1` foi substituído por `/v1`.

A mudança simplifica a URL base e alinha com a convenção de usar apenas o número de versão como prefixo.

## Endpoints alterados

**Todas as rotas** tiveram o prefixo alterado de `/api/v1/` para `/v1/`. Nenhuma rota foi adicionada ou removida.

| Recurso | Antes | Depois |
|---|---|---|
| Auth | `/api/v1/auth/*` | `/v1/auth/*` |
| Accounts | `/api/v1/accounts/*` | `/v1/accounts/*` |
| Account Plans | `/api/v1/account-plans/*` | `/v1/account-plans/*` |
| Modules | `/api/v1/modules/*` | `/v1/modules/*` |
| System Users | `/api/v1/system-users` | `/v1/system-users` |
| Users (tenant) | `/api/v1/users/*` | `/v1/users/*` |
| Customers | `/api/v1/customers/*` | `/v1/customers/*` |
| Plans | `/api/v1/plans/*` | `/v1/plans/*` |
| Products | `/api/v1/products/*` | `/v1/products/*` |
| Subscriptions | `/api/v1/subscriptions/*` | `/v1/subscriptions/*` |

## Contratos quebrados

**Breaking change**: a URL base muda de `http://localhost:8080/api/v1` para `http://localhost:8080/v1`. O front-end deve atualizar a `baseURL` do cliente HTTP.

## O que o front precisa fazer

- Atualizar a `baseURL` / `baseUrl` configurada no cliente HTTP (axios, fetch, etc.) removendo o segmento `/api`.
- Nenhuma mudança em payloads, headers, autenticação ou responses.
