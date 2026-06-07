# Arquitetura

O SubClub adere ao padrão de **Clean Architecture** (Arquitetura Limpa), dividindo o sistema em quatro camadas bem definidas que garantem o isolamento entre regras de negócio centralizadas e adaptadores externos (como bancos de dados e requisições HTTP):

```
domain → application → web → infra
```

As dependências sempre apontam para dentro — camadas externas nunca são importadas pelas internas.

---

## Camadas

### 1. Domain (`internal/domain`)

O coração da aplicação. Cada entidade de negócio (ex.: `customer`, `plan`, `account`) é organizada em:

```
internal/domain/<entidade>/
    model/          → structs do domínio, erros e inputs tipados
    repository.go   → interface do repositório (contrato de persistência)
    service.go      → interface do service (contrato de orquestração)
    usecase/        → implementações dos casos de uso (create, list, get_by_id, update, delete)
```

Essa camada não conhece SQL, JSON ou qualquer framework externo.

### 2. Application (`internal/application`)

Implementações concretas das interfaces do Domain, divididas em dois sub-pacotes:

- `service/` — orquestração dos fluxos de negócio (ex.: `service/customer`, `service/auth`).
- `repository/` — acesso ao banco de dados via Ent ORM (ex.: `repository/customer`, `repository/plan`).

### 3. Web (`internal/web`)

Camada de apresentação HTTP. Contém os _Handlers_ (controladores de rotas REST), os _DTOs_ (contratos de entrada e saída em JSON) e a struct `Handlers` em `internal/web/handlers.go`, que agrega e inicializa todos os handlers a partir dos services da camada Application.

Para detalhes de organização e fluxo HTTP, veja [web-layer.md](web-layer.md).

### 4. Infrastructure (`internal/infra`)

Ferramentas e serviços de suporte, organizados em sub-pacotes:

- `server/` — inicialização do Echo, middlewares e registro de rotas (`server.go`, `router.go`).
- `database/` — conexão com o banco, migrações, seeder e `TenantClientManager` (gerenciamento de schemas por tenant).
- `middleware/` — middlewares de autenticação e autorização (`auth.go`, `admin.go`, `logger.go`).
- `authctx/` — helpers para leitura/escrita do usuário autenticado no contexto da requisição.
- `tenantctx/` — helpers para leitura/escrita do tenant ativo no contexto da requisição.

---

## Documentos relacionados

- [ORM & Convenções de Banco](orm.md)
- [Camada Web](web-layer.md)
- [Multi-tenant](multi-tenant.md)
- [Auth / Roles](auth_levels.md)
