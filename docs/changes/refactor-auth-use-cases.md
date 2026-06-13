# refactor/auth-use-cases

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

Nenhum endpoint foi alterado. Rotas, payloads e responses permanecem idênticos:

- `POST /api/v1/auth/login`
- `POST /api/v1/auth/lookup`
- `POST /api/v1/auth/tenant-login`

## Contratos quebrados

Nenhum — refactor interno. O front-end não precisa de nenhuma alteração.

## Notas para o front

Nenhuma mudança observável. Esta PR é puramente estrutural.
