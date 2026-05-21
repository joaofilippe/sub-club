# Architecture & ORM Design

## Visão Geral
O SubClub adere ao padrão de **Clean Architecture** (Arquitetura Limpa), dividindo o sistema em quatro camadas bem definidas que garantem o isolamento entre regras de negócio centralizadas e adaptadores externos (como bancos de dados e requisições HTTP):

1. **Domain (`internal/domain`)**: O coração da aplicação. Define os modelos brutos (`structs`), as interfaces dos Repositórios e regras comuns. Essa camada não conhece SQL, JSON, ou qualquer banco de dados de terceiros.
2. **Application (`internal/application`)**: Contém os _Services_ (orquestradores de fluxo de entrada/saída) e os _Repositories_ (implementações concretas que traduzem o Domain em queries no BD via Ent ORM).
3. **Web (`internal/web`)**: Camada de apresentação HTTP. Contém os _Handlers_ (controladores de rotas REST), os _DTOs_ (contratos de entrada e saída em JSON) e a struct `Handlers`, que agrega e inicializa todos os handlers a partir dos services da camada Application.
4. **Infrastructure (`internal/infra`)**: Ferramentas e serviços de suporte. Contém o `Server` (inicialização do Echo, middlewares e registro de rotas), a conexão com o banco de dados e o logger.

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

### Cliente vs Customer
O nome estrutural "Client" (`client`) é uma palavra reservada pela própria arquitetura interna do gerador `ent` (que a utiliza na raiz para instanciar as conexões e ponteiros, i.e., `ent.Client`). Se nomeássemos o schema do usuário como `Client`, o escopo interno do gerador causaria conflitos infinitos na compilação em Go.
Para solucionar, adotamos a premissa de *Boundary Context* (Contextos Mapeáveis):

* **No Domain e na Web**: O modelo se chama **Client** (`internal/domain/client`). Todos os Web Endpoints (`GET`, `POST`, etc.) rodam perfeitamente na rota de subdomínio `/clients`.
* **No Banco de Dados e ORM**: A estrutura física chama-se **Customer** (`ent/schema/customer.go`). Nossos conversores mapeadores em `internal/adapter/repository/client_ent.go` absorvem do banco como *Customer*, mas devolvem via API e regras de negócio como *Client*.

### Utilização do SQLX para TCP
O objeto `sqlx.DB` ainda existe residindo pacificamente em `internal/infra/database`. Seu papel não é lançar as queries diretas, mas sim cuidar ativamente da manutenção da _Connection Pool_, como *keep-alive*, número limite de portas atreladas ao Linux e gerenciamento TCP contra o container de Docker `postgres`.
Entregamos essa conexão TCP em forma limpa atestando: `drv := entsql.OpenDB("postgres", a.dbConnection.GetDB().DB)`.

### Soft Delete Dinâmico (Exclusão Lógica)
Alguns módulos, como o `User`, não removem linhas destrutivas nas chamadas da API do tipo `DELETE`, graças a injeção do conceito de *Soft Delete* no banco de dados.
Sempre que um model precisar de restro/histórico arquivável, adote o preceito do Ent usando o modificador dinâmico: `field.Time("deleted_at").Optional().Nillable()`. Nas buscas, adicione filtros seguros `.Where(user.DeletedAtIsNil())`.
