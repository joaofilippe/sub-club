# feat/username-auth — Username-based authentication

## O que mudou e por quê

O fluxo de login foi redesenhado para eliminar a opção de `email + slug` e introduzir um login por `username + account_slug`. O novo comportamento:

1. `POST /auth/login` recebe apenas `email + password`. O sistema tenta login como system user (público); se não encontrar, faz lookup nos tenants pelo email.
   - 1 tenant encontrado → retorna token direto
   - 0 tenants → 401
   - 2+ tenants → 409 com mensagem orientando a usar o login por username
2. `POST /auth/login/username` recebe `username + account_slug + password` → retorna token do tenant.

O campo `username` foi adicionado ao schema de usuário tenant (`ent/schema/user.go`) como obrigatório e único dentro do tenant.

---

## Endpoints novos ou alterados

### `POST /api/v1/auth/login` — alterado

**Antes:**
```json
{ "email": "...", "password": "...", "account_slug": "..." }
```

**Depois** (campo `account_slug` removido):
```json
{ "email": "...", "password": "..." }
```

**Responses:**
- `200` — token retornado (sistema ou tenant único)
- `401` — credenciais inválidas
- `409` — múltiplos tenants encontrados; use `/auth/login/username`

---

### `POST /api/v1/auth/login/username` — novo

**Request:**
```json
{ "username": "joao", "account_slug": "minha-empresa", "password": "..." }
```

**Response `200`:**
```json
{ "data": { "token": "..." } }
```

**Responses:**
- `200` — token retornado
- `400` — campos obrigatórios ausentes
- `401` — credenciais inválidas

---

### `POST /api/v1/users` — campo `username` agora obrigatório

**Request atualizado:**
```json
{
  "name": "...",
  "username": "joao.barista",
  "email": "...",
  "password": "...",
  "role": "operations"
}
```

**Response atualizado** (campo `username` incluído):
```json
{
  "id": "...",
  "name": "...",
  "username": "joao.barista",
  "email": "...",
  "role": "operations",
  "createdAt": "...",
  "updatedAt": "..."
}
```

---

## Contratos quebrados / campos alterados

- **`POST /auth/login`**: campo `account_slug` **removido** do payload — qualquer client que enviava esse campo deve parar de fazê-lo.
- **`POST /api/v1/users`**: campo `username` **agora obrigatório** — requests sem ele retornarão 400.
- `CreateTenantOwner` (interno): assinatura alterada para incluir `username`. O seeder cria o owner do tenant demo com `username: "admin"`. Novos accounts criados via API recebem username derivado do prefixo do email do account.

---

## Seeder / dados de demo

- Login sistema: `adm@adm.com` / `12345678`
- Login tenant demo por email: `admin@demo.com` / `12345678`
- Login tenant demo por username: `username: admin`, `account_slug: demo`, `password: 12345678`
