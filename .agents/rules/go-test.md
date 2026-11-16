---
trigger: always_on
---

# Diretrizes para Testes em Golang

Ao criar ou refatorar testes no SubClub, siga estritamente estas regras:

1. **Table-Driven Tests**: Sempre utilize o padrão de Table-Driven Tests (slice de structs anônimas com `t.Run(tt.name, func(t *testing.T) { ... })`) para cobrir múltiplos cenários de entrada e saída na mesma função de teste.
2. **Nomenclatura Clara**: Use o formato `Test[NomeDaFuncaoOuStruct]_[Cenario]` (ex.: `TestUserHandler_Create`, `TestCustomerUseCase_Update`). Os nomes dos cenários dentro da tabela (`name`) devem ser descritivos.
3. **Padrão AAA (Arrange, Act, Assert)**: Mantenha o corpo do laço `for _, tt := range tests` limpo e sequencial: prepare os dados, execute a função e verifique resultado/erro.
4. **Cobertura e Tratamento de Erros**: Valide tanto o caminho feliz quanto casos de borda e erro (ex.: `tt.wantErr`, status HTTP, conteúdo da mensagem de erro quando aplicável).
5. **Test Helpers e Faker**:
   - Qualquer função auxiliar extraída para limpar o setup do teste deve obrigatoriamente chamar `t.Helper()` na sua primeira linha.
   - Use os helpers de `internal/test/faker` (`faker.FakeUser()`, `faker.FakeCustomer()`, etc.) para dados de domínio — evite strings hardcoded repetitivas para fixtures.
6. **Mocks (Unitários) vs Banco Real (Integração/Repositórios)**:
   - **Testes unitários** (handlers, usecases, services): devem depender das interfaces do domínio (`Repository`, `Service`) via Dependency Injection, usando mocks locais no mesmo arquivo `*_test.go` ou pacote de teste.
   - **Testes de integração e repositórios Ent**: devem usar conexão real com o banco de dados (ou `enttest`, conforme o padrão já usado na suíte) — sem mocks de banco de dados.
7. **Verificação Obrigatória**: Antes de concluir qualquer tarefa com alteração de código, garanta que os testes passem executando `make test` (`go test -v ./...`).
