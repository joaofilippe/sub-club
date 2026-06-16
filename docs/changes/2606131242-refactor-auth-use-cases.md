# refactor/auth-use-cases — atualizado

## O que mudou e por quê

Extração da lógica de negócio do `AuthService` para use cases dedicados. O service passou a ser um adapter puro entre a camada web e o domínio — cada método delega diretamente ao use case correspondente.

## Arquivos criados

| Arquivo | Responsabilidade |
|---|---|
| `internal/domain/auth/usecase.go` | Interfaces `LoginUseCase`, `TenantLoginUseCase`, `LookupUseCase` |
| `internal/application/usecase/auth/token.go` | Struct `Claims` (JWT) + helper `signToken` compartilhado |
| `internal/application/usecase/auth/login.go` | Implementação de `LoginUseCase` |
| `internal/application/usecase/auth/tenant_login.go` | Implementação de `TenantLoginUseCase` |
| `internal/application/usecase/auth/lookup.go` | Implementação de `LookupUseCase` |

## Arquivos modificados

| Arquivo | O que mudou |
|---|---|
| `internal/application/service/auth/service.go` | Virou adapter: 3 campos de interface, 3 métodos de 1 linha |
| `internal/application/application.go` | Wiring: instancia os 3 use cases e injeta no `NewAuthService` |
| `internal/infra/middleware/auth.go` | Import de `authsvc.Claims` trocado por `authusecase.Claims` |
| `internal/infra/middleware/admin.go` | Idem |
| `internal/infra/server/server_test.go` | `NewAuthService` atualizado para nova assinatura (3 args) |

## Endpoints

### Alterações de contrato (breaking)

| Antes | Depois |
|---|---|
| `POST /api/v1/auth/login` — body: `{ email, password }` | `POST /api/v1/auth/login` — body: `{ email, password, account_slug? }` |
| `POST /api/v1/auth/tenant-login` — body: `{ email, password, account_slug }` | **Removido** |
| `POST /api/v1/auth/lookup` | Inalterado |

### Novo comportamento de `/api/v1/auth/login`

- **Sem `account_slug`**: autentica system user (esquema público, `system_users`)
- **Com `account_slug`**: autentica usuário tenant (esquema `account_{slug}`, `users`)

Ambos retornam o mesmo `{ token: "eyJ..." }`.

## Notas para o front

- **Remover** as chamadas para `/auth/tenant-login`
- Passar `account_slug` no body do `/auth/login` quando for login de tenant
- O fluxo típico continua: Lookup → Login com o slug escolhido pelo usuário
