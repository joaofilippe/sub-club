# PR: Add `demot` account to database seeder

## O que mudou e por quê?
Foi adicionada a criação automática da account `Demot` (slug: `demot`, email: `demot@demot.com`) no seeder da aplicação (`SeedAll`).
Isso facilita o ambiente de desenvolvimento, permitindo que ao iniciar a aplicação localmente (`make dev` ou `docker compose up`), dois tenants já estejam disponíveis para testes: `demo` e `demot`.

A documentação em `docs/development/seeding-and-fake-data.md` também foi atualizada para refletir a nova account gerada por padrão.

## Endpoints novos ou alterados
Nenhum endpoint foi criado ou alterado.

## Contratos quebrados ou campos renomeados
Nenhum.

## Detalhes para o Front-end
- Agora há uma segunda account pronta para uso no ambiente local (development) que usa as mesmas credenciais de admin (`admin@demot.com` / `12345678`), útil para testar fluxos de troca de tenant ou isolamento de dados sem precisar criar manualmente pelo painel admin.
