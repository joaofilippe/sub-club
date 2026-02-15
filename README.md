# SubClub

Este projeto segue o layout padrão de projetos Go (Standard Go Project Layout). Abaixo está a descrição da finalidade de cada diretório:

## Estrutura de Diretórios

- **`/cmd`**: Ponto de entrada das aplicações deste projeto. O código aqui deve conter apenas a função `main` e chamar o código localizado em `/internal` e `/pkg`.
  - Exemplo: `/cmd/subclub/main.go`
  - **Por que subpastas?**: Ter um diretório separado para cada aplicação (neste caso `subclub`) permite que o projeto evolua para ter múltiplos binários (ex: um worker, um CLI, uma versão webassembly) sem bagunça. Cada diretório gera um binário com o nome da pasta.

- **`/internal`**: Código privado da aplicação e bibliotecas. O código neste diretório não pode ser importado por outros projetos. É aqui que a lógica de negócio específica da aplicação deve residir.

- **`/pkg`**: Código de biblioteca que pode ser utilizado por aplicações externas. Outros projetos podem importar os pacotes localizados aqui. Coloque aqui apenas código que foi projetado para ser reutilizável.

- **`/api`**: Especificações da API (OpenAPI/Swagger), definições de esquema JSON, arquivos de definição de protocolo (como arquivos .proto).

- **`/configs`**: Modelos de arquivos de configuração ou configs padrão.

- **`/scripts`**: Scripts para realizar operações de build, instalação, análise, entre outros.

- **`/build`**: Configuração e scripts para empacotamento e Integração Contínua (CI).

- **`/deployments`**: Configurações e templates para implantação (IaaS, PaaS), orquestração de sistemas e containers (ex: docker-compose, charts do kubernetes).

- **`/test`**: Apps de teste externos adicionais e dados de teste.

- **`/docs`**: Documentação de design e do usuário (além da documentação gerada pelo godoc).

## Como Executar

Para rodar a aplicação principal:

```bash
go run cmd/subclub/main.go
```
