# feat/tenant-user-management

## O que mudou e por quê

Implementada gestão completa de usuários do tenant. Tenant admins agora podem criar, listar, editar e remover os outros usuários do seu schema (operadores, baristas, etc.). Usuários com role `operations` não têm acesso a esses endpoints — apenas `admin`.

**Breaking change**: o endpoint de criação de system users foi movido de `POST /users` para `POST /system-users` para evitar conflito de rota e deixar claro a distinção entre usuários da plataforma e usuários do tenant.

## Endpoints novos

Todos exigem `Authorization: Bearer <tenantToken>` com `role = "admin"`.

| Método | Rota | Descrição |
|--------|------|-----------|
| `POST` | `/users` | Cria um novo usuário no tenant |
| `GET` | `/users` | Lista todos os usuários ativos do tenant |
| `GET` | `/users/:id` | Retorna um usuário pelo ID |
| `PUT` | `/users/:id` | Atualiza nome e/ou role |
| `DELETE` | `/users/:id` | Soft delete |

### `POST /users` — body

```json
{
  "name": "Barista João",
  "email": "joao@cafeorigem.com.br",
  "password": "senha123",
  "role": "operations"
}
```

### `POST /users` e `GET /users/:id` — response

```json
{
  "id": "uuid",
  "name": "Barista João",
  "email": "joao@cafeorigem.com.br",
  "role": "operations",
  "createdAt": "2026-06-07T00:00:00Z",
  "updatedAt": "2026-06-07T00:00:00Z"
}
```

A senha **nunca** é retornada nas respostas.

## Endpoint alterado (breaking)

| Antes | Depois |
|-------|--------|
| `POST /users` (admin token) | `POST /system-users` (admin token) |

## Roles disponíveis

| Role | Acesso a `/users` |
|------|------------------|
| `admin` | Total (CRUD) |
| `operations` | Nenhum (403) |

## Detalhes internos

- Novo middleware `RequireTenantAdminMiddleware` — lido do `authctx.Claims.Role` já injetado pelo `AuthMiddleware`.
- Soft delete via campo `deleted_at`; todos os queries filtram com `DeletedAtIsNil()`.
- Hash bcrypt da senha feito no usecase `CreateTenantUserUseCase`, não no handler.
