# Camada Web (HTTP)

Este documento descreve como a camada HTTP do SubClub é organizada e como suas peças se conectam.

---

## Visão Geral

A entrega HTTP da aplicação é dividida entre dois pacotes com responsabilidades distintas:

| Pacote | Responsabilidade |
|---|---|
| `internal/web` | Handlers HTTP e DTOs — lógica de apresentação da API |
| `internal/infra/server` | Servidor Echo, middleware e registro de rotas |

---

## `internal/web` — Handlers e DTOs

### Struct `Handlers`

`web.Handlers` agrega e inicializa todos os handlers da aplicação a partir dos services:

```go
type Handlers struct {
    Account     *account.AccountHandler
    AccountPlan *accountplan.AccountPlanHandler
    Auth        *auth.AuthHandler
    Customer    *customer.CustomerHandler
    Module      *module.ModuleHandler
    Plan        *plan.PlanHandler
    Product     *product.ProductHandler
    Subscription *subscription.SubscriptionHandler
    User        *user.UserHandler
}

func NewHandlers(app *application.Application) *Handlers
```

É o único ponto onde a camada `web` conhece a camada `application`. Cada handler recebe apenas o service de que precisa, sem acesso ao `Application` inteiro.

### Handlers por domínio

Cada subpacote de `internal/web` cuida de um domínio:

```
internal/web/
    handlers.go              → agrega todos os handlers
    account/
        handler.go           → CRUD de accounts (rotas admin)
        dto.go
    accountplan/
        handler.go           → CRUD de account plans (rotas admin)
        dto.go
    auth/
        handler.go           → login / geração de JWT
        dto.go
    customer/
        handler.go           → CRUD de clientes do tenant
        dto.go
    module/
        handler.go
        dto.go
    plan/
        handler.go           → CRUD de planos do tenant
    product/
        handler.go           → CRUD de produtos do tenant
    subscription/
        handler.go           → CRUD de assinaturas do tenant
    user/
        handler.go           → CRUD de usuários do tenant
    common/
        response.go          → helpers de resposta padronizada
```

---

## `internal/infra/server` — Servidor HTTP

### Struct `Server`

`Server` é o ponto de entrada único para toda a infraestrutura HTTP. Ele compõe internamente um `router` e um `logger`:

```go
type Server struct {
    echo   *echo.Echo
    logger *slog.Logger
    router *router
}

func NewServer(handlers *web.Handlers) *Server
```

Ao ser construído, o `Server`:
1. Inicializa o Echo
2. Configura os middlewares (logger estruturado e recovery)
3. Registra a rota `/health`
4. Delega o registro das rotas de negócio ao `router`

### Struct `router` (interna)

`router` é uma struct não exportada dentro do pacote `server`, responsável exclusivamente pelo registro das rotas de negócio:

```go
type router struct {
    echo     *echo.Echo
    handlers *web.Handlers
}
```

Essa separação garante que mudanças em rotas não afetam a configuração do servidor, e vice-versa.

---

## Fluxo de inicialização

```
cmd/subclub/application.go
    │
    ├── application.New(db)          → instancia services e repositories
    │
    ├── web.NewHandlers(app)         → inicializa todos os handlers
    │
    └── server.NewServer(handlers)   → configura middleware, registra rotas, sobe o servidor
```

---

## Fluxo de uma requisição

```
HTTP Request
    → Echo (infra/server)
    → Middleware (logger, recover)
    → Handler (internal/web/<domínio>)
    → Service (internal/application/service/<domínio>)
    → Repository (internal/application/repository/<domínio>)
    → Ent ORM → PostgreSQL
```

## Padrão de Response

Toda resposta da API segue o mesmo envelope JSON, definido em `internal/web/common/response.go`:

```go
type ErrorDetail struct {
    Message string `json:"message"`
}

type Response struct {
    Message string       `json:"message"`
    Data    any          `json:"data,omitempty"`
    Error   *ErrorDetail `json:"error,omitempty"`
}
```

Os campos `data` e `error` usam `omitempty` — apenas um deles estará presente por resposta.

### Resposta de sucesso

```json
{
  "message": "Customer created",
  "data": {
    "id": "a1b2c3d4-...",
    "name": "João Silva",
    "email": "joao@email.com"
  }
}
```

### Resposta de erro

```json
{
  "message": "invalid request body",
  "error": {
    "message": "invalid request body"
  }
}
```

### Helpers obrigatórios

Todos os handlers devem usar exclusivamente os helpers abaixo. Chamadas diretas a `c.JSON` são proibidas fora do pacote `common`.

```go
// Sucesso — popula data, omite error
common.Success(c, http.StatusOK, "mensagem", dto)

// Erro — popula error, omite data
common.Error(c, http.StatusBadRequest, "mensagem de erro")
```

### Convenções de mensagem

| Operação | Mensagem de sucesso |
|---|---|
| `POST` (create) | `"<Recurso> created"` — ex: `"Customer created"` |
| `GET` | `"OK"` |
| `PUT` (update) | `"<Recurso> updated"` — ex: `"Plan updated"` |
| `DELETE` | `"<Recurso> deleted"` — ex: `"Account deleted"` |
| `GET` (list) | `"OK"` |