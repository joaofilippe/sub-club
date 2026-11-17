---
trigger: always_on
---

# Onde Consultar (Fonte da Verdade em `docs/`)

Tudo que diz respeito a **arquitetura, camadas, nomenclatura, convenções de código, banco de dados, HTTP e testes** tem sua fonte da verdade exclusiva em `docs/` (e `README.md` para setup geral).

Antes de propor ou alterar código, leia o documento correspondente ao escopo da tarefa:

| Assunto / Tarefa | Onde olhar |
|---|---|
| Arquitetura, camadas (`domain → application → web → infra`), idioma e nomenclatura | `docs/design/architecture.md` |
| Ent ORM, schemas, soft delete, `Customer` e passo a passo de nova entidade | `docs/design/orm.md` |
| Camada HTTP, handlers, DTOs, `common.Success`/`Error`, Swag e Postman | `docs/design/web-layer.md` |
| Multi-tenancy (`Schema-per-Tenant`), `TenantClientFromContext` e rotas | `docs/design/multi-tenant.md` |
| Usuários, Types e Roles | `docs/design/auth_levels.md` |
| Padrões e estrutura de testes em Go | `docs/development/testing.md` |
| Seeder e helpers `internal/test/faker` | `docs/development/seeding-and-fake-data.md` |
| Ambiente Docker, Air e Delve | `docs/development/docker-setup.md` |
| Histórico recente de mudanças e contratos de API | `docs/changes/` |
