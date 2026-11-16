---
trigger: always_on
---

# Diretrizes de Arquitetura e Convenções (SubClub)

## 1. Clean Architecture (4 Camadas) e Estrutura de Pastas

As dependências sempre apontam **para dentro**:

```
domain → application → web → infra
```

- **`internal/domain/<entidade>/`**: Núcleo puro da aplicação. **Nunca** deve importar `application`, `web`, `infra` ou `ent`.
  - `model/`: Entidades, Value Objects, Enums, erros de domínio e inputs tipados.
  - `repository.go`: Interface que define o contrato de persistência (Porta de Saída).
  - `service.go`: Interface que define o contrato de orquestração (Porta).
  - `usecase/`: Casos de uso e regras de negócio (`create`, `list`, `get_by_id`, `update`, `delete`, etc.), dependendo apenas de `model` e das interfaces do domínio.
- **`internal/application/`**: Implementações concretas das interfaces do domínio.
  - `repository/<entidade>/`: Repositórios concretos usando **Ent ORM**.
  - `service/<entidade>/`: Services concretos que orquestram fluxos de negócio e atuam como facade dos `usecase`s para a camada Web.
- **`internal/web/`**: Camada de apresentação HTTP.
  - `handlers.go`: Struct `Handlers` (`NewHandlers`) que agrega e injeta os services em cada handler.
  - `<entidade>/`: Handlers REST (Echo), DTOs de requisição/resposta JSON e anotações Swagger (`// @Summary`, `// @Router`, etc.).
  - `common/`: Helpers e envelope padronizado de resposta (`response.go`).
- **`internal/infra/`**: Adaptadores de infraestrutura e entrega.
  - `server/`: Inicialização do Echo, loggers, middlewares e registro de rotas (`server.go`, `router.go`).
  - `database/`: Pool TCP (`*sql.DB`), migrações, seeder e `TenantClientManager` (schemas isolados por tenant).
  - `middleware/`: Middlewares de autenticação e autorização (`auth.go`, `admin.go`, `logger.go`).
  - `authctx/` e `tenantctx/`: Helpers para leitura/escrita do usuário autenticado e do tenant ativo no contexto da requisição.
- **`ent/schema/`**: Schemas declarativos do Ent ORM (código typesafe gerado em `ent/` via `go generate ./ent`).

## 2. Ent ORM, Multi-tenant e Convenções de Dados

- **Ent ORM (Sem SQL Raw)**: Nunca escrever SQL raw. Toda query passa pelo Ent. Após alterar `ent/schema/`, executar `go generate ./ent`. O `*sql.DB` em `internal/infra/database` existe apenas para gerenciar o pool TCP e alimentar o driver Ent.
- **Multi-tenant (Schema-per-Tenant)**:
  - Schema público (`public`): `system_users`, `accounts`, `account_plans`.
  - Schema por tenant (`account_{slug}`): `users`, `customers`, `plans`, `products`, `subscriptions`.
  - Handlers de rotas tenant devem obter o `*ent.Client` isolado via `tenantctx.TenantClientFromContext(ctx)`. **Nunca** usar o client global em rotas tenant.
  - Entidades de tenant não devem possuir edges partindo de `Account` ou `AccountPlan` (`public`).
- **Soft Delete**: Entidades arquiváveis usam `field.Time("deleted_at").Optional().Nillable()`, e toda query deve filtrar com `.Where(<entity>.DeletedAtIsNil())`.
- **Nomenclatura `Customer`**: O assinante final do tenant chama-se `Customer` em todas as camadas (evitando conflito com `ent.Client`).

## 3. Modelagem e Idioma

- **Idioma do Código**: Use **inglês** para todo o código (nomes de pacotes, structs, interfaces, métodos, funções e variáveis).
- **IDs e Construtores**: Prefira IDs UUID e tipagem explícita de inputs/outputs no `domain/<entidade>/model`.

## 4. Swagger, Postman e Comandos Comuns

- Todos os handlers HTTP devem manter anotações Swag atualizadas.
- Sempre que houver mudança em endpoints (handler, rota, payload ou response), executar `make swagger` e atualizar `docs/subclub.postman_collection.json`.

```bash
make build         # compila para bin/subclub
make test          # go test -v ./...
make run           # executa localmente com APP_ENV=development
make dev           # docker compose up -d db pgadmin && docker compose up app
make swagger       # regenera docs/swagger.* e docs/docs.go a partir dos handlers
make lint          # golangci-lint run
make tidy          # go mod tidy
go generate ./ent  # regenera código Ent ORM após mudanças em ent/schema/
```
