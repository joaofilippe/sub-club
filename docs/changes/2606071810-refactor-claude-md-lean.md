# refactor/claude-md-lean

## O que mudou e por quê

Refatoração interna do `CLAUDE.md` — arquivo de instruções para o agente de IA (Claude Code). Nenhuma alteração de código de produção, endpoints ou contratos.

O arquivo foi reduzido de ~120 para ~65 linhas, substituindo seções detalhadas de arquitetura por referências aos documentos já existentes em `docs/`. O conteúdo inline que permaneceu é apenas o que é relevante em praticamente toda tarefa: comandos, convenções e workflow do agente.

## Impacto para o front-end

Nenhum. Esta mudança é exclusivamente de documentação interna do agente de backend.
