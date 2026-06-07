# Architecture & ORM Design

## Visão Geral
O SubClub adere ao padrão de **Clean Architecture** (Arquitetura Limpa), dividindo o sistema em quatro camadas bem definidas que garantem o isolamento entre regras de negócio centralizadas e adaptadores externos (como bancos de dados e requisições HTTP):

1. **Domain (`internal/domain`)**: O coração da aplicação. Cada entidade de negócio (ex.: `customer`, `plan`, `account`) é organizada em:
   - `model/` — structs do domínio, erros e inputs tipados.
   - `repository.go` — interface do repositório (contrato de persistência).
   - `service.go` — interface do service (contrato de orquestração).
   - `usecase/` — implementações dos casos de uso (`create`, `list`, `get_by_id`, `update`, `delete`).
   Essa camada não conhece SQL, JSON ou qualquer framework externo.

2. **Application (`internal/application`)**: Implementações concretas das interfaces do Domain, divididas em dois sub-pacotes:
   - `service/` — orquestração dos fluxos de negócio (ex.: `service/customer`, `service/auth`).
   - `repository/` — acesso ao banco de dados via Ent ORM (ex.: `repository/customer`, `repository/plan`).

3. **Web (`internal/web`)**: Camada de apresentação HTTP. Contém os _Handlers_ (controladores de rotas REST), os _DTOs_ (contratos de entrada e saída em JSON) e a struct `Handlers` em `internal/web/handlers.go`, que agrega e inicializa todos os handlers a partir dos services da camada Application.

4. **Infrastructure (`internal/infra`)**: Ferramentas e serviços de suporte, organizados em sub-pacotes:
   - `server/` — inicialização do Echo, middlewares e registro de rotas (`server.go`, `router.go`).
   - `database/` — conexão com o banco, migrações, seeder e `TenantClientManager` (gerenciamento de schemas por tenant).
   - `middleware/` — middlewares de autenticação e autorização (`auth.go`, `admin.go`, `logger.go`).
   - `authctx/` — helpers para leitura/escrita do usuário autenticado no contexto da requisição.
   - `tenantctx/` — helpers para leitura/escrita do tenant ativo no contexto da requisição.

---

## Adoção do Ent ORM

Historicamente, o projeto realizava operações de banco de dados via *Raw SQL* formatadas na biblioteca `sqlx`. Para garantir a integridade dos dados, facilitar integrações, e garantir *auto-migrates* robustos, introduzimos a lib **Ent ORM** como motor principal de relacionamento relacional.

### Como funciona?
Diferente de geradores onde você anota "structs", com o `Ent` a arquitetura funciona com base na leitura declarativa:
- Desenvolvemos as características, atributos relacionais (*Edges*) e obrigatoriedades na pasta `ent/schema`.
- Rodamos o utilitário nativo via terminal: `go generate ./ent`.
- O framework varre as anotações gerando pastas robustas de Queries, Deleções, Predicates e Constraints (dentro de `ent/{nomeDaEntidade}`) de forma *typesafe* — nenhum SQL em texto puro vulnerável circula pelo projeto.

---

## Casos Especiais & Convenções Mapeadas

### Customer (antes "Client")
O nome `client` é palavra reservada pelo gerador Ent (usado em `ent.Client` para representar a conexão com o banco). Para evitar conflito de compilação, a entidade que representa o assinante final do tenant foi nomeada **Customer** em todas as camadas — domain (`internal/domain/customer`), ORM (`ent/schema/customer.go`) e web (`internal/web/customer`). As rotas HTTP expõem o recurso como `/customers`.

### Utilização do SQLX para TCP
O objeto `sqlx.DB` ainda existe residindo pacificamente em `internal/infra/database`. Seu papel não é lançar as queries diretas, mas sim cuidar ativamente da manutenção da _Connection Pool_, como *keep-alive*, número limite de portas atreladas ao Linux e gerenciamento TCP contra o container de Docker `postgres`.
Entregamos essa conexão TCP em forma limpa atestando: `drv := entsql.OpenDB("postgres", a.dbConnection.GetDB().DB)`.

### Soft Delete Dinâmico (Exclusão Lógica)
Alguns módulos, como o `User`, não removem linhas destrutivas nas chamadas da API do tipo `DELETE`, graças a injeção do conceito de *Soft Delete* no banco de dados.
Sempre que um model precisar de restro/histórico arquivável, adote o preceito do Ent usando o modificador dinâmico: `field.Time("deleted_at").Optional().Nillable()`. Nas buscas, adicione filtros seguros `.Where(user.DeletedAtIsNil())`.
