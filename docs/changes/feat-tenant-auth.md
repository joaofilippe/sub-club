# feat/tenant-auth

## O que mudou e por quê

Implementado o fluxo completo de autenticação para usuários de tenant. Antes só era possível logar como `system_user` (admin da plataforma). Agora o front consegue descobrir em quais tenants um email está cadastrado e fazer login em um tenant específico.

## Fluxo recomendado no front

```
1. POST /api/v1/auth/lookup   { email }
   → retorna [{ slug, name }]

   - Se retornar 1 item: avançar direto para o login com esse slug
   - Se retornar vários: mostrar seletor para o usuário escolher
   - Se retornar vazio: email não encontrado em nenhum tenant

2. POST /api/v1/auth/tenant-login   { email, password, account_slug }
   → retorna { token }
   → usar esse token no header Authorization: Bearer <token> para todas as rotas autenticadas
```

## Endpoints novos

### `POST /api/v1/auth/lookup`
Público. Nenhum token necessário.

**Request:**
```json
{ "email": "usuario@empresa.com" }
```

**Response 200:**
```json
{
  "data": [
    { "slug": "minha-cafeteria", "name": "Minha Cafeteria" }
  ]
}
```

Estratégia interna: tenta lookup por domínio do email (tabela `account_domains`); se o domínio for genérico (gmail, outlook, etc.), faz scan paralelo em todos os tenants.

---

### `POST /api/v1/auth/tenant-login`
Público. Nenhum token necessário.

**Request:**
```json
{
  "email": "usuario@empresa.com",
  "password": "senha",
  "account_slug": "minha-cafeteria"
}
```

**Response 200:**
```json
{
  "data": { "token": "<jwt>" }
}
```

O JWT gerado contém `account_slug` e `role` — é o mesmo token aceito por todas as rotas autenticadas existentes.

## Nova tabela: `account_domains`

Schema público. Mapeia domínios de email a accounts para lookup rápido.

```
account_domains
  id          UUID  PK
  domain      TEXT  UNIQUE  (ex: "minha-empresa.com")
  account_id  UUID  → accounts.id
```

O front não interage com essa tabela diretamente. Ela é populada via API quando um admin registra um domínio corporativo para o seu tenant (endpoint a ser criado futuramente).

## Sem breaking changes

O endpoint `POST /api/v1/auth/login` existente não foi alterado — continua funcionando para system users (admins da plataforma).
