---
trigger: always_on
---

# Diretrizes de Fluxo Git e Documentação (Branches, Commits e Changelog)

Sempre que uma alteração no repositório for solicitada, siga estritamente estas regras:

## 1. Verificação da Branch Atual

- Antes de realizar qualquer alteração em arquivos do projeto, verifique a branch atual (`git branch --show-current`).
- Se estiver na branch `main` (ou `master`), **nunca** faça alterações diretamente nela. Crie e mude imediatamente para uma nova branch (`git checkout -b <type>/<short-description>`).

## 2. Padrão de Nomeação de Branches (em Inglês)

Utilize `kebab-case` e nomes descritivos em **inglês**, seguindo os prefixos de Conventional Commits:

- `feat/<short-description>`: Novas funcionalidades.
- `fix/<short-description>`: Correções de bugs.
- `refactor/<short-description>`: Refatorações de código sem mudança de comportamento.
- `test/<short-description>`: Adição ou correção de testes.
- `docs/<short-description>`: Alterações apenas em documentação.
- `chore/<short-description>`: Configurações, dependências, regras de agentes ou tarefas de manutenção.

## 3. Padrão de Commits e Controle de Push

- Commitar ao fim de cada tarefa concluída.
- Todas as mensagens de commit devem ser escritas em **inglês** e seguir [Conventional Commits](https://www.conventionalcommits.org/):
  - Formato: `<type>(<optional-scope>): <description>`
  - Exemplos:
    - `feat(customer): add active status filter`
    - `fix(auth): validate tenant slug before lookup`
    - `chore(rules): adapt agent rules structure`
- **Não fazer push nem abrir PR** sem que o usuário peça explicitamente.

## 4. Pós-Commit e Sincronia com Documentação (SubClub)

1. **Changelog para o Front-end**: Após cada commit, criar `docs/changes/<YYMMDDHHMM>-<branch-name>.md` (com o timestamp `YYMMDDHHMM` do momento da criação para ordenação cronológica) em estilo PR, cobrindo:
   - O que mudou e por quê;
   - Endpoints novos ou alterados (método, rota, payload, response);
   - Contratos quebrados, campos renomeados ou detalhes relevantes para o front-end (ou indicar explicitamente quando não houver impacto de API).
2. **Verificar `docs/`**: Avaliar se algum documento em `docs/` (`design/`, `development/`, etc.) precisa ser atualizado para refletir a nova realidade. Se sim, incluir no mesmo commit ou em um commit imediatamente seguinte.
3. **Atualizar Swagger e Postman**: Após qualquer mudança que afete endpoints, executar `make swagger` para regenerar o OpenAPI (`docs/swagger.json`, `docs/swagger.yaml`, `docs/docs.go`) e atualizar `docs/subclub.postman_collection.json`, incluindo ambos no mesmo commit da mudança.
