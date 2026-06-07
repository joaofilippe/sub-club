# refactor/split-architecture-orm-docs

## O que mudou e por quê

Reorganização interna de documentação. O arquivo `docs/design/architecture_orm.md` foi dividido em dois documentos com responsabilidades distintas:

- `docs/design/architecture.md` — descreve as quatro camadas da Clean Architecture (domain, application, web, infra) e a estrutura de pacotes.
- `docs/design/orm.md` — descreve o Ent ORM, convenções de banco (Customer, pool TCP, Soft Delete) e o passo a passo para adicionar uma nova entidade.

## Impacto para o front-end

Nenhum. Mudança exclusivamente de documentação interna de backend.
