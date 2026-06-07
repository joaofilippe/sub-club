# feat/account-owner-auto-creation

## O que mudou e por quê

Ao criar uma `Account`, o sistema agora provisiona automaticamente o primeiro `User` admin do tenant usando o mesmo email da `Account`. Antes, esse usuário precisava ser criado manualmente (ou só existia no seeder de dev). O campo `Account.email` passa a ter uso funcional real: é a credencial de login do owner do tenant.

## Endpoints alterados

### `POST /accounts`

**Request** — sem alteração.

**Response** — novo campo `ownerPassword` (presente apenas na resposta de criação, nunca em outros endpoints):

```json
{
  "id": "uuid",
  "name": "Café Origem",
  "email": "contato@cafeorigem.com.br",
  "document": "12.345.678/0001-90",
  "slug": "cafe-origem",
  "accountPlanId": "uuid",
  "subscriptionStatus": "trial",
  "active": true,
  "createdAt": "2026-06-07T00:00:00Z",
  "ownerPassword": "a3f8c2d1e4b07591"
}
```

`ownerPassword` é uma senha temporária gerada aleatoriamente (16 chars hex). O admin do SubClub deve exibi-la uma única vez e repassar ao cliente do tenant. Não é armazenada em texto claro — apenas o hash bcrypt fica no banco.

## Contratos / comportamento novo

- O tenant User owner é criado com `email = account.email`, `name = account.name`, `role = "admin"`.
- Em dev, o seeder usa `demo@subclub.com / 12345678` para o tenant demo (consistente com o email da Account demo).
- A senha temporária **não** é retornada em `GET /accounts`, `PUT /accounts` ou `GET /accounts/:id` — o campo é `omitempty`.

## Detalhes internos relevantes para o front

- O `ownerPassword` deve ser exibido em tela uma única vez após a criação da account, com aviso explícito para o admin anotá-lo.
- Não existe endpoint para recuperar essa senha — se perdida, será necessário um reset de senha (a implementar).
