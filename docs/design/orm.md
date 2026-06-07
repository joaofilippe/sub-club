# ORM & Convenções de Banco

---

## Ent ORM

O projeto usa o **Ent ORM** como motor principal de acesso ao banco de dados. Nenhuma query SQL raw deve ser escrita — tudo passa pelo Ent.

### Como funciona

Diferente de ORMs baseados em anotações de structs, o Ent usa um modelo declarativo de schemas:

1. Defina a entidade em `ent/schema/<entidade>.go` (campos, edges, índices).
2. Rode `go generate ./ent`.
3. O gerador produz código typesafe em `ent/<entidade>/` com queries, predicates e constraints — nenhum SQL em texto puro circula pelo projeto.

---

## Convenções

### Customer (antes "Client")

O nome `client` é palavra reservada pelo gerador Ent — é usado em `ent.Client` para representar a conexão com o banco. Para evitar conflito de compilação, a entidade que representa o assinante final do tenant foi nomeada **Customer** em todas as camadas:

- Domain: `internal/domain/customer`
- ORM: `ent/schema/customer.go`
- Web: `internal/web/customer`
- Rotas HTTP: `/customers`

### Pool TCP via `*sql.DB`

O objeto `*sql.DB` em `internal/infra/database` não executa queries — existe exclusivamente para gerenciar o pool de conexões TCP (keep-alive, limite de portas, reconexão contra o container Postgres). O driver Ent é alimentado a partir dele:

```go
drv := entsql.OpenDB("postgres", db)
client := ent.NewClient(ent.Driver(drv))
```

### Soft Delete (Exclusão Lógica)

Entidades que precisam de histórico arquivável não são deletadas fisicamente. Use o campo `deleted_at`:

```go
field.Time("deleted_at").Optional().Nillable()
```

Toda query deve filtrar registros deletados:

```go
.Where(<entidade>.DeletedAtIsNil())
```

---

## Adicionando uma nova entidade

1. Crie o schema em `ent/schema/<entidade>.go`.
2. Rode `go generate ./ent`.
3. Implemente o repositório em `internal/application/repository/<entidade>/ent.go`.
4. Defina as interfaces em `internal/domain/<entidade>/repository.go` e `service.go`.
5. Implemente os casos de uso em `internal/domain/<entidade>/usecase/`.
6. Implemente o service em `internal/application/service/<entidade>/service.go`.
7. Crie o handler e DTOs em `internal/web/<entidade>/`.
8. Registre o handler em `internal/web/handlers.go` e a rota em `internal/infra/server/router.go`.

> Se a entidade pertencer ao schema de tenant (não ao `public`), **não** adicione edges a partir de `Account` ou `AccountPlan`. Veja [multi-tenant.md](multi-tenant.md) para detalhes.
