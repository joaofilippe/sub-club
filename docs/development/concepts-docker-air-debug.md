# Mergulho Profundo: Docker, Air e Debug Remoto em Go

Este documento explica os conceitos fundamentais por trás do nosso ambiente de desenvolvimento local. O objetivo é desmistificar como as peças se encaixam: desde a exposição de portas no Docker até a mágica do hot-reload e como o VS Code consegue "entrar" no container para debugar.

## 1. Docker e Redes: "Quem ouve o quê?"

Quando rodamos uma aplicação dentro do Docker, ela está em uma rede isolada. Se o servidor Go diz "estou ouvindo na porta 8080", ele está gritando em uma sala fechada (o container). Ninguém de fora (seu computador, o "Host") consegue ouvir.

### Exposição de Portas (`ports`)
No `docker-compose.yml`, usamos a diretiva `ports` para criar um "buraco na parede":

```yaml
ports:
  - "8080:8080"  # Porta do Host : Porta do Container
```

- **Esquerda (Host)**: A porta que você acessa no seu navegador (`localhost:8080`).
- **Direita (Container)**: A porta onde sua aplicação Go está rodando (`:8080`).

Se mudássemos para `"3000:8080"`, você acessaria `localhost:3000`, mas o Docker redirecionaria o tráfego para a porta 8080 dentro do container.

### Permissões Especiais (Para o Debugger)
Para o debugger funcionar, precisamos de permissões de baixo nível do Linux:
- `cap_add: ["SYS_PTRACE"]`: Permite que um processo (o debugger) controle outro processo (sua app).
- `security_opt: ["seccomp:unconfined"]`: Desabilita alguns filtros de segurança padrão do Docker que bloqueiam chamadas de sistema necessárias para o debug.

---

## 2. Air: O Motor de Hot-Reload

O **Air** é uma ferramenta que melhora a experiência de desenvolvimento (DX). Em Go, como a linguagem é compilada, para ver uma mudança você precisa: Parar o server -> Recompilar -> Iniciar o server. O Air automatiza isso.

### Como funciona?
1.  **Watcher**: O Air fica vigiando os arquivos `.go` no diretório.
2.  **Detecção**: Quando você salva um arquivo, o Air detecta a mudança.
3.  **Build**: Ele roda o comando de build definido no `.air.toml` (`go build ...`).
4.  **Restart**: Ele mata o processo antigo e inicia o novo binário.

### Configuração Crítica para Debug
No nosso caso, não queremos apenas rodar o binário (`./tmp/main`). Queremos rodá-lo *através* do debugger. Por isso a configuração no `.air.toml` é complexa:

```toml
full_bin = "dlv exec --headless --continue --listen=:2345 ... -- ./tmp/main"
```

Isso diz ao Air: "Quando você reiniciar a aplicação, não rode ela direto. Rode o **Delve**, e peça para o Delve rodar a aplicação".

---

## 3. Delve e VS Code: A Conexão Remota

Aqui é onde a mágica do "Attach" acontece.

### O Servidor de Debug (Delve)
O **Delve** (`dlv`) é o debugger padrão do Go. No nosso container, ele está rodando em modo **Headless** (sem interface gráfica), agindo como um servidor.
- `--listen=:2345`: Ele fica ouvindo na porta 2345 por comandos.
- `--continue`: Ele inicia a aplicação imediatamente, sem esperar que um cliente (VS Code) se conecte. Sem isso, o app ficaria travado no início esperando o "ok" do debugger.

### O Cliente (VS Code)
O VS Code não está rodando o código Go. Ele está apenas agindo como uma interface gráfica (UI) para o Delve que está lá no container.

Qunado você clica em "Attach to Docker" no `launch.json`:
1.  O VS Code se conecta via TCP em `localhost:2345`.
2.  Graças ao mapeamento de portas do Docker, essa conexão chega no container.
3.  O VS Code começa a falar com a API do Delve: "Coloque um breakpoint na linha 50 do main.go".
4.  O Delve (no container) insere essa parada na memória do processo rodando.

### "Attach" vs "Launch"
- **Launch**: O VS Code inicia o processo. Comum em desenvolvimento local sem Docker.
- **Attach**: O processo já existe (iniciado pelo Air/Docker). O VS Code apenas se "gruda" nele para inspecionar.

---

## Resumo do Fluxo

1.  Você salva um arquivo.
2.  **Air** detecta -> Recompila -> Reinicia o processo usando **Delve**.
3.  **Delve** sobe a aplicação na porta 8080 e fica ouvindo comandos de debug na 2345.
4.  Você no **VS Code** se conecta na 2345.
5.  Você acessa `localhost:8080`, sua requisição entra no Docker, bate na aplicação, o Delve pausa a execução no breakpoint, e o VS Code mostra o estado das variáveis.
