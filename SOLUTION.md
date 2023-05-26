## Configuração

Certifique-se de ter os seguintes requisitos instalados:

- Node.js (versão 16)
- Docker

## Instalação

1. Clone este repositório:

   ```bash
   git clone https://github.com/nironwp/mono-challenge-intention.git
   ```
2. Navegue até o diretório:

   ```bash
   cd mono-challenge-intention
   ```
3. Rode o docker-compose up:

```bash
docker-compose up -d
```

## Uso

Certifique-se de ter realizado a configuração e instalação!

1. O serviço de produtos estará disponível no seguinte endpoint:

http://localhost:8080/graphql

2. O serviço de intenções estará disponivel no seguinte endpoint:

http://localhost:3001/query

## Detalhes sobre a solução

### Product service

- Linguagem: Nodejs
- Framework: Nestjs
- Arquitetura: DDD
- Padrões de desenvolvimento:TDD, DRY

- Extras: Foi incorporado rate-limiting (para evitar DDOS), csurf, fastify para
  respostas mais rapidas, e SWC para melhor desempenho da aplicação além de uma
  compilação mais rapida o que ajuda no ambiente de desenvolvimento

### Intent Service

- Linguagem: Golang
- Framework: Nenhum
- Arquitetura: Hexagonal
- Padrões de desenvolvimento:TDD, Solid, DRY

- Extras: Por ter sido feito usando arquitetura hexagonal, esse microserviço nos
  da a opção de "plugar" e "desplugar" componentes como por exemplo, podemos
  trocar nosso banco mongodb para algum banco RavenDb, Postgres ou até mesmo
  Redis, podemos passar a receber solicitações ao invés de graphQL com grpc
