# SubClub — Claude Code Guide

## Claude Behavior

Atue como um **Dev Sênior Fullstack**: opine sobre design, aponte trade-offs, sugira melhorias, mas implemente apenas o que for pedido — sem abstrações prematuras nem escopo extra.

### Workflow obrigatório para qualquer mudança de código

1. **Criar branch** antes de qualquer alteração:
   ```bash
   git checkout -b <tipo>/<descricao-curta>
   # ex: feat/customer-filter, fix/auth-middleware, refactor/tenant-ctx
   ```
2. **Commitar** ao fim de cada tarefa concluída, com mensagem em inglês seguindo Conventional Commits (`feat:`, `fix:`, `refactor:`, `docs:`, `chore:`).
3. **Não fazer push nem abrir PR** sem o usuário pedir explicitamente.
4. **Após cada commit**, criar `docs/changes/<branch-name>.md` com descrição no estilo PR para o agente de front-end, cobrindo: o que mudou e por quê, endpoints novos ou alterados (método, rota, payload, response), contratos quebrados ou campos renomeados, e qualquer detalhe que o front precise saber.
5. **Verificar `/docs`**: após qualquer mudança de código, avaliar se algum documento em `docs/` precisa ser atualizado para refletir a nova realidade — arquitetura, contratos, convenções, seeder, etc. Se sim, atualizar e incluir no mesmo commit ou num commit imediatamente seguinte.
6. **Atualizar Swagger e Postman**: após qualquer mudança que afete endpoints (novo handler, rota alterada, payload ou response modificado), executar `make swagger` para regenerar o OpenAPI e atualizar o `postman_collection.json` correspondente. Incluir ambos no mesmo commit da mudança.

---

## Common Commands

```bash
make build      # compile to bin/subclub
make test       # go test -v ./...
make run        # run locally with APP_ENV=development
make dev        # docker compose up -d db pgadmin && docker compose up app (hot-reload via air + dlv debug)
make swagger    # regenerate OpenAPI docs from handler annotations
make lint       # golangci-lint run
make tidy       # go mod tidy
go generate ./ent  # regenerate Ent ORM code after schema changes
```

---

## Architecture

Quatro camadas — dependências sempre apontam para dentro:

```
domain → application → web → infra
```

Para detalhes completos de cada camada e pacotes, consulte:
- [Arquitetura & ORM](docs/design/architecture_orm.md)
- [Camada Web](docs/design/web-layer.md)
- [Multi-tenant](docs/design/multi-tenant.md)
- [Auth / Roles](docs/design/auth_levels.md)
- [Seeder & Fake Data](docs/development/seeding-and-fake-data.md)
- [Docker / Air / Debug](docs/development/docker-setup.md)

---

## Key Conventions

### Ent ORM
- Nunca escrever SQL raw. Todas as queries passam pelo Ent.
- Schemas em `ent/schema/`. Após qualquer mudança: `go generate ./ent`.
- O `*sql.DB` em `internal/infra/database` existe apenas para gerenciar o pool TCP e alimentar o driver Ent.

### Multi-tenant
- Schema público (`public`): `system_users`, `accounts`, `account_plans`.
- Schema por tenant (`account_{slug}`): `users`, `customers`, `plans`, `products`, `subscriptions`.
- Handlers de rotas tenant obtêm o `*ent.Client` isolado via `tenantctx.TenantClientFromContext(ctx)`.
- Nunca usar o client global em handlers de rotas tenant.

### Soft Delete
```go
field.Time("deleted_at").Optional().Nillable()
// Sempre filtrar com:
.Where(<entity>.DeletedAtIsNil())
```

### Swagger & Postman
- Todos os handlers precisam de anotações Swag (`// @Summary`, `// @Router`, etc.).
- Após qualquer mudança de endpoint: executar `make swagger` para regenerar o OpenAPI e atualizar o `postman_collection.json`.
- Swagger e Postman devem sempre estar em sincronia com o código — nunca commitar um handler novo sem atualizar ambos.

### Tests
- Usar helpers `faker.FakeUser()`, `faker.FakeCustomer()`, etc. — nunca strings hardcoded.
- Testes de integração usam conexão real com o banco — sem mocks de banco.
