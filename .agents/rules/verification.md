---
trigger: always_on
---

# Verificação e Testes (Checklist do Agente)

Antes de criar, alterar ou concluir qualquer tarefa que modifique código:

1. **Testes**: Leia `docs/development/testing.md` e `docs/development/seeding-and-fake-data.md` antes de escrever ou refatorar testes.
2. **Geração Ent ORM**: Se alterou `ent/schema/`, execute `go generate ./ent`.
3. **Swagger e Postman**: Se alterou handlers, rotas, payloads ou responses, execute `make swagger` e atualize `docs/subclub.postman_collection.json`.
4. **Execução Obrigatória**: Rode `make test` (`go test -v ./...`) após a última edição e só conclua a tarefa quando todos os testes passarem (e execute `make lint` quando aplicável).
