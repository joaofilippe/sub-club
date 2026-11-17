# Padrões de Testes em Go

Este documento define as convenções de escrita, organização e execução de testes no SubClub.

---

## Convenções de Código de Teste

1. **Table-Driven Tests**: Utilize o padrão de Table-Driven Tests (slice de structs anônimas com `t.Run(tt.name, func(t *testing.T) { ... })`) para cobrir múltiplos cenários de entrada e saída na mesma função de teste.
2. **Nomenclatura Clara**: Use o formato `Test[NomeDaFuncaoOuStruct]_[Cenario]` (ex.: `TestUserHandler_Create`, `TestCustomerUseCase_Update`). Os nomes dos cenários dentro da tabela (`name`) devem ser descritivos.
3. **Padrão AAA (Arrange, Act, Assert)**: Mantenha o corpo do laço `for _, tt := range tests` limpo e sequencial — prepare os dados, execute a função e verifique o resultado/erro esperado.
4. **Cobertura e Tratamento de Erros**: Valide tanto o caminho feliz quanto casos de borda e erro (ex.: `tt.wantErr`, status HTTP, conteúdo da mensagem de erro quando aplicável).
5. **Test Helpers**: Qualquer função auxiliar extraída para limpar o setup do teste deve obrigatoriamente chamar `t.Helper()` na primeira linha.

---

## Dados Fake (`internal/test/faker`)

- Utilize os helpers de `internal/test/faker` (`faker.FakeUser()`, `faker.FakeCustomer()`, etc.) para gerar dados de domínio — evite strings hardcoded repetitivas para fixtures.
- Veja a lista completa de helpers disponíveis em [seeding-and-fake-data.md](seeding-and-fake-data.md).

---

## Mocks vs Banco Real

- **Testes unitários** (handlers, usecases, services): devem depender das interfaces do domínio (`Repository`, `Service`) via Dependency Injection, usando mocks locais no próprio arquivo `*_test.go` ou pacote de teste.
- **Testes de integração e repositórios Ent**: devem usar conexão real com o banco de dados (ou `enttest` com SQLite em memória, conforme o padrão da suíte em `internal/application/repository/user/ent_test.go`) — sem mocks de banco de dados.

---

## Execução

```bash
make test          # executa go test -v ./...
```
