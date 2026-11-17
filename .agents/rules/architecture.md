---
trigger: always_on
---

# Guardrails de Arquitetura e Convenções (SubClub)

A especificação detalhada da arquitetura vive em `docs/design/` e `docs/development/`. Esta regra mantém apenas as restrições invioláveis e os comandos do dia a dia.

## 1. Direção de Dependências (Clean Architecture)

```
domain → application → web → infra
```

- **`internal/domain/`** é o núcleo puro: **nunca** deve importar `application`, `web`, `infra` ou `ent`.
- Siga a organização e os contratos detalhados em `docs/design/architecture.md` e `docs/design/web-layer.md`.

## 2. Guardrails Críticos

- **Ent ORM (Sem SQL Raw)**: Toda query passa pelo Ent (`ent/schema/`). Após alterar schemas, execute `go generate ./ent`. O `*sql.DB` em `internal/infra/database` existe apenas para pool TCP e driver Ent.
- **Multi-tenant (`Schema-per-Tenant`)**:
  - Handlers de rotas tenant **devem** obter o `*ent.Client` isolado via `tenantctx.TenantClientFromContext(ctx)` — nunca use o client global em rotas tenant.
  - Não crie edges partindo de entidades do schema `public` (`Account`, `AccountPlan`) para entidades de tenant (`account_{slug}`).
- **Soft Delete**: Entidades arquiváveis usam `field.Time("deleted_at").Optional().Nillable()` e toda query deve filtrar `.Where(<entity>.DeletedAtIsNil())`.
- **Nomenclatura `Customer`**: O assinante final do tenant chama-se `Customer` em todas as camadas (evita conflito com `ent.Client`).
- **Camada Web**: Handlers HTTP devem usar `common.Success` / `common.Error` (`internal/web/common/response.go`) em vez de `c.JSON` direto, e manter anotações Swag atualizadas.
- **Idioma**: Use **inglês** para todo código (pacotes, structs, interfaces, métodos, funções e variáveis).

## 3. Documentação de Referência (Fonte da Verdade)

- Arquitetura e camadas: `docs/design/architecture.md`
- Ent ORM e nova entidade: `docs/design/orm.md`
- Camada HTTP, Handlers e DTOs: `docs/design/web-layer.md`
- Multi-tenancy e isolamento: `docs/design/multi-tenant.md`
- Usuários, Types e Roles: `docs/design/auth_levels.md`
- Seeder / Fake data e Docker: `docs/development/seeding-and-fake-data.md`, `docs/development/docker-setup.md`

## 4. Comandos Comuns

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
