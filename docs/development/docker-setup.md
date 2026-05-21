# Ambiente de Desenvolvimento Local com Docker

Este documento detalha a configuração do ambiente de desenvolvimento local utilizando Docker, Docker Compose, Air (para hot-reload) e Delve (para debug remoto).

## 🚀 Visão Geral

O objetivo desta configuração é fornecer um ambiente de desenvolvimento robusto, isolado e reproduzível, que espelhe as dependências de produção (como PostgreSQL) enquanto oferece uma experiência de desenvolvimento ágil.

### Componentes

| Componente | Função | Porta Host | Detalhes |
| ------------ | -------- | ------------ | ---------- |
| **App** | Aplicação Go | `:8080` | Roda via `air` para hot-reload e `dlv` para debug. |
| **PostgreSQL** | Banco de Dados | `:5432` | Versão 15-alpine. Dados persistidos no volume `postgres_data`. |
| **PgAdmin** | Gestão de Banco | `:5050` | Interface web para gerenciar o PostgreSQL. |
| **Swagger** | Documentação API | `:8080/swagger` | Interface interativa para testar os endpoints. |
| **Health Check** | Status da API | `:8080/health` | Endpoint para verificar se a API está online. |
| **Delve** | Debugger | `:2345` | Servidor de debug remoto. |

## 🛠️ Modificações Realizadas

### 1. Dockerfile

Criado um `Dockerfile` multi-stage baseado em `golang:1.24-alpine`.

- **Instalações**: `git` e `curl` (para ferramentas de build) e `dlv` (debugger).
- **Air**: Instalado via script binário para evitar dependências de build no runtime.
- **Exposição**: Portas `8080` (HTTP) e `2345` (Debug).

### 2. docker-compose.yml

Define a orquestração dos serviços:

- **`app`**: Monta o diretório atual em `/app` para permitir que o `air` detecte mudanças no código. Configurado com `security_opt: ["seccomp:unconfined"]` e `cap_add: ["SYS_PTRACE"]` para permitir o funcionamento do debugger.
- **`db`**: Instância padrão do Postgres.
- **`pgadmin`**: Interface administrativa para o banco.

### 4. Variáveis de Ambiente e Seeding

O ambiente está configurado para habilitar a população automática de dados fakes (Seeder):

- **`APP_ENV=development`**: Ativa a lógica de seeding.
- **Seeding Automático**: Se o banco estiver vazio, o sistema cria 50 clientes, 10 produtos, 4 planos e um administrador (`adm@adm.com` / `12345678`).

### 3. Configuração do Air (.air.toml)

Ajustado para suportar o ambiente Linux do container e o debugger:

- **Binário**: Caminho do binário alterado para `./tmp/main` (sem `.exe`).
- **Build**: Adicionadas flags `-gcflags="all=-N -l"` para desabilitar otimizações e permitir debug preciso.
- **Execução**: O app é iniciado através do `dlv exec` com a flag `--continue`, garantindo que o servidor inicie imediatamente sem bloquear aguardando conexão do debugger.

### 4. VS Code Launch Configuration (.vscode/launch.json)

Adicionada configuração "Attach to Docker" para debug remoto.

- **Debug Adapter**: Configurado como `"legacy"` para compatibilidade com a API v2 do Delve rodando no container.

## 🏃 Como Rodar

1.**Iniciar o ambiente**:
    ```bash
    docker-compose up --build
    ```
2. **Acessar a API**:
    - URL: `http://localhost:8080`
    - Health Check: `http://localhost:8080/health`

3.**Acessar PgAdmin**:
    - URL: `http://localhost:5050`
    - **Email**: `admin@admin.com`
    - **Senha**: `root`
    - **Conectar ao Banco**:
        - Host: `db`
        - Username: `user`
        - Password: `password`
        - Maintenance DB: `subclub`
        - SSL Mode: `Disable`

4.**Desenvolvimento**:
    - Basta editar os arquivos `.go` localmente. O `air` detectará a mudança, recompilará o binário e reiniciará o servidor automaticamente dentro do container.

## 🐛 Como Debugar

O debugger está sempre ativo na porta `2345`.

1.No VS Code, vá para a aba **Run and Debug** (Ctrl+Shift+D).
2.Selecione **"Attach to Docker"**.
3.Clique no botão de Play (▶).
4.Adicione breakpoints no seu código local e faça requisições para testar.

> **Nota**: O ambiente usa `dlv --continue`, então a aplicação não para na inicialização. Ela só parará quando atingir um breakpoint que você definir.
