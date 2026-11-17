# chore/adapt-agents-rules

## O que mudou e por quê

Reorganização das instruções do agente em `.agents/rules/` (padrão modular adaptado do `intiv2`) com separação estrita de responsabilidades em relação a `docs/`:

- **`.agents/rules/`** (apenas comportamento do agente, onde consultar, fluxo Git, atualização de docs e verificação):
  - `persona.md`: postura sênior pragmática, qualidade de código, previsibilidade e abordagem spec-driven.
  - `references.md`: mapa de para onde o agente deve olhar em `docs/design/`, `docs/development/` e `docs/changes/`.
  - `git-workflow.md`: verificação e criação de branch `kebab-case`, Conventional Commits em inglês, controle de push/PR, criação de `docs/changes/`, revisão de `docs/` e sincronia de Swagger/Postman.
  - `verification.md`: checklist de validação antes de concluir tarefas (`go generate ./ent`, `make swagger`, `make test`, `make lint`).
- **`docs/`** (fonte única da verdade sobre código, arquitetura, nomenclatura e convenções):
  - `docs/design/architecture.md`: explícita a restrição de `internal/domain` não importar `ent` e adicionada seção de idioma/nomenclatura em inglês.
  - `docs/design/web-layer.md`: adicionada seção de documentação OpenAPI (Swag) e Postman.
  - `docs/development/testing.md`: criado documento consolidando todas as convenções de testes em Go (Table-Driven Tests, `Test[Nome]_[Cenario]`, AAA, `t.Helper()`, `faker`, mocks vs banco real/`enttest` e `make test`).
- Removido `.agents/AGENTS.md` (e regras duplicadas de arquitetura/testes em `.agents/`).

## Impacto para o front-end

Nenhum. Alteração restrita às regras do agente de desenvolvimento e à documentação interna de engenharia; sem mudanças em código de produção, rotas, payloads ou contratos da API.
