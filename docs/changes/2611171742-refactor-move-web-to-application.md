# refactor/move-web-to-application

## O que mudou e por quê

Movida a camada de apresentação HTTP de `internal/web` para `internal/application/web`, alinhando a estrutura de pacotes do SubClub ao modelo `domain → application (service / repository / web) → infra`:

- Movido todo o pacote `internal/web/` (handlers, DTOs, `common/response.go` e `handlers.go`) para `internal/application/web/`.
- Atualizados os imports em `cmd/subclub/application.go`, `internal/application/web/...`, `internal/infra/middleware/...` e `internal/infra/server/...`.
- Atualizada a documentação em `docs/design/architecture.md`, `docs/design/web-layer.md`, `docs/design/orm.md`, `.agents/rules/references.md`, `CLAUDE.md` e `README.md`.

## Impacto para o front-end

Nenhum. Refatoração exclusivamente interna de organização de pacotes em Go; nenhuma rota, método HTTP, payload, response ou contrato de API foi alterado.
