# chore/adapt-agents-rules

## O que mudou e por quê

Reorganização das instruções do agente de IA em `.agents/` para o padrão modular em `.agents/rules/` (adaptado do repositório `intiv2`), preservando integralmente a arquitetura e as convenções do SubClub.

- Substituído `.agents/AGENTS.md` pelos arquivos modulares em `.agents/rules/` (`trigger: always_on`):
  - `persona.md`: persona sênior pragmática, Clean Code, previsibilidade antes de grandes alterações e abordagem spec-driven consultando `docs/design/` e `docs/development/`.
  - `architecture.md`: guardrails críticos da arquitetura (`domain → application → web → infra`, Ent ORM sem SQL raw, `tenantctx.TenantClientFromContext`, soft delete, `Customer`, `common.Success`/`Error`), comandos `make` e ponteiros para `docs/design/` e `docs/development/` (sem duplicar os documentos de design).
  - `git-workflow.md`: verificação de branch antes de editar, criação de branch `kebab-case`, Conventional Commits em inglês, controle de push/PR e criação de `docs/changes/`.
  - `go-test.md`: padrão Table-Driven Tests, AAA, `t.Helper()`, helpers de `internal/test/faker`, mocks para testes unitários e banco real/`enttest` para testes de repositório e integração.
- Não foi incluído `rpa-rules.md` por não se aplicar ao escopo de backend REST do SubClub.

## Impacto para o front-end

Nenhum. Alteração restrita às regras internas de trabalho do agente de desenvolvimento; sem mudanças em código de produção, rotas, payloads ou contratos da API.
