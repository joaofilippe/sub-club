# SubClub

SubClub is a subscription management application built with Go. It provides a robust API server using Echo and a versatile CLI using Cobra.

## Prerequisites

Before you begin, ensure you have the following installed:

- **Go**: 1.24+ 
- **Docker & Docker Compose**: For running the database and peripheral services.
- **Make**: To simplify common development tasks.

## Getting Started

1. **Clone the repository:**
   ```bash
   git clone https://github.com/joaofilippe/subclub.git
   cd subclub
   ```

2. **Install dependencies:**
   ```bash
   make deps
   ```

3. **Start the infrastructure:**
   Use Docker Compose to start the PostgreSQL database and PgAdmin:
   ```bash
   docker-compose up -d
   ```

4. **Run the application:**
   You can start the API server locally using:
   ```bash
   make run
   ```
   The API will be available at `http://localhost:8080`.

## Development

We use `air` for hot-reloading during development.

- **Hot Reload**: `make dev`
- **Run Tests**: `make test`
- **Linting**: `make lint`
- **Tidy Dependencies**: `make tidy`
- **Generate API Docs**: `make swagger`

### Infrastructure
- **API Server**: Runs on port `8080`.
- **Database**: PostgreSQL on port `5432`.
- **PgAdmin**: Available at `http://localhost:5050` (Login: `admin@admin.com` / `root`).

## API Documentation (Swagger)

The project includes automatically generated OpenAPI documentation using the **Swag** generator.

1. Generate the static definitions analyzing code annotations: `make swagger`
2. Start the application: `make run`
3. Access the interactive user interface in your browser: [http://localhost:8080/swagger/index.html](http://localhost:8080/swagger/index.html)

## Architecture & Database

This application utilizes **Clean Architecture** combined with the **Ent ORM** for standardizing interactions with the PostgreSQL database.

- **Clean Architecture**: 
  - Business rules lie securely within `/internal/domain`. 
  - Controllers, Services, and Handlers lie in `/internal/adapter`.
  - External adapters and connection establishments (`sqlx`, `echo`) reside in `/internal/infra`.

- **Ent ORM**:
  - The framework is heavily mapped internally. Schemas are defined at `/ent/schema/`. Run `go generate ./ent` to reflect mapping changes.
  - *Developer Note*: The internal API entity dealing with Clients is physically stored as **Customer** in the database to prevent naming collisions with the `ent.Client` keyword. The web routes continue to operate under `/clients` smoothly.

## Project Structure

Este projeto segue o layout padrão de projetos Go (Standard Go Project Layout). Abaixo está a descrição da finalidade de cada diretório:

- **`/cmd`**: Ponto de entrada das aplicações deste projeto. O código aqui deve conter apenas a função `main` e chamar o código localizado em `/internal`.
  - Exemplo: `/cmd/subclub/main.go` 
- **`/internal`**: Código privado da aplicação e bibliotecas. O código neste diretório não pode ser importado por outros projetos. É aqui que a lógica de negócio reside.
- **`/api`**: Especificações da API (OpenAPI/Swagger) e definições de esquema.
- **`/configs`**: Modelos de arquivos de configuração ou configs padrão.
- **`/scripts`**: Scripts para realizar operações de build, instalação, etc.
- **`/build`**: Configuração e scripts para empacotamento e CI.
- **`/deployments`**: Configurações de implantação (Docker Compose, Kubernetes).
- **`/test`**: Apps de teste externos adicionais e dados de teste.
- **`/docs`**: Documentação OpenAPI gerada pelo Swagger, guia arquiteturais de design e manuais do usuário.
- **`/pkg`**: (Opcional) Código de biblioteca que pode ser utilizado por aplicações externas.


