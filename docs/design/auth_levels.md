# Arquitetura de Usuários (Types vs Roles)

Este documento descreve as definições e a lógica por trás da estrutura de usuários no SubClub.

## Racional da Separação

Para garantir escalabilidade e flexibilidade (especialmente pensando em um modelo B2B e B2C), decidimos separar a **natureza** do usuário de suas **permissões**.

1.  **UserType (O que o usuário É)**: Define a entidade jurídica/natureza da conta. Essencial para regras de negócio como faturamento, impostos e fluxos de cadastro.
2.  **UserRole (O que o usuário FAZ)**: Define as permissões de acesso dentro da plataforma. Um "Customer" Pessoal e um "Customer" Corporativo podem ter o mesmo Role, mas são Types diferentes.

---

## Tipos de Usuário (User Types)

O `UserType` define a **natureza** da conta, o que impacta em como os dados são tratados (ex: CPF vs CNPJ).

1. **Individual (`individual`)**: Pessoa física, o cliente padrão do clube.
2. **Corporativo (`corporate`)**: Empresas ou escritórios que assinam café para o time.
3. **Sistema (`system`)**: Contas para automações ou integrações externas.

## Papéis de Usuário (User Roles)

O `UserRole` define o que o usuário **pode fazer** (permissões).

### 1. Administrador (`admin`)
- **Acesso:** Total ao sistema.
- **Responsabilidades:** Configurações globais, gestão de usuários e relatórios financeiros.

### 2. Operações (`operations`)
- **Acesso:** Gestão logística e estoque.
- **Responsabilidades:** Processamento de pedidos, envios e controle de café disponível.

### 3. Cliente (`customer`)
- **Acesso:** Área do cliente e e-commerce.
- **Responsabilidades:** Gestão da própria assinatura, compras avulsas e histórico pessoal.

## Futuras Expansões (Ideias)
- **Curador/Barista:** Para gestão técnica de lotes e notas sensoriais.
- **Suporte:** Para atendimento ao cliente sem acesso a configurações críticas.
- **VIP/Assinante:** Benefícios exclusivos dentro da categoria de cliente.
