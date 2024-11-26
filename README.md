# Curso Full Cycle 3.0 - Módulo sobre Arquitetura Orientada a Eventos (EDA - Event Driven Architecture)

---

## Resumo

O Curso Full Cycle é uma formação abrangente, projetada para capacitar desenvolvedores a trabalharem em projetos de grande porte, empregando as melhores práticas de desenvolvimento e criando aplicações robustas.

---

## Desafio

Implemente um microsserviço na linguagem de sua preferência que seja capaz de consumir, via Kafka, os eventos gerados pelo microsserviço "Wallet Core". Este serviço deverá persistir os saldos atualizados de cada conta em um banco de dados.

Além disso, crie um endpoint: `/balances/{account_id}`, que exiba o saldo atualizado para a conta especificada.

### Requisitos para entrega:
- [x] Toda a aplicação deve ser executável por meio de Docker / Docker Compose.
- [x] Um único comando `docker-compose up -d` deve ser capaz de iniciar todos os microsserviços, incluindo o "Wallet Core".
- [x] As migrações devem ser executadas automaticamente, e os bancos de dados (tanto do "Wallet Core" quanto do serviço de "Balances") devem ser populados com dados fictícios ao iniciar os serviços.
- [x] Crie um arquivo `.http` que permita realizar chamadas ao seu microsserviço, assim como foi feito no "Wallet Core".
- [x] O microsserviço de "Balances" deve estar acessível na porta: 3003.

---

## Instruções de Execução

Para compilar e iniciar todos os serviços, execute na raiz do projeto: 

```bash
docker-compose up --build

```

Esse comando irá baixar e configurar todas as imagens necessárias.

---

## Estrutura do Projeto

O projeto foi organizado em dois microsserviços principais:

1. **Wallet Core**: Responsável pela geração dos eventos relacionados às movimentações financeiras.
2. **Balances**: Responsável por consumir os eventos via Kafka, atualizar os saldos no banco de dados e fornecer um endpoint para consulta.

Cada serviço possui sua configuração independente, mas todos são integrados por meio do Docker Compose.

### Estrutura de Arquivos e Pastas

```plaintext
.
├── wallet-core/
│   ├── cmd/
│   │   └── walletcore/
│   │       └── main.go  # Arquivo principal para iniciar o serviço
│   ├── db/
│   │   ├── migrations/  # Arquivos de migração do banco de dados
│   │   └── seed.sql     # Dados fictícios para popular o banco
│   └── api/
│       └── client.http  # Arquivo para testes de API
├── balances/
│   ├── cmd/
│   │   └── balances/
│   │       └── main.go  # Arquivo principal para iniciar o serviço
│   ├── db/
│   │   ├── migrations/  # Arquivos de migração do banco de dados
│   │   └── seed.sql     # Dados fictícios para popular o banco
│   └── api/
│       └── client.http  # Arquivo para testes de API
├── docker-compose.yml  # Configuração do Docker Compose para orquestração
└── README.md           # Documentação do projeto
```

## Detalhes Técnicos

### Kafka
O sistema utiliza o Apache Kafka como broker de mensagens para comunicação assíncrona entre os microsserviços. A configuração inclui a criação de tópicos e a definição de consumidores e produtores.

- **Tópico utilizado:** `wallet-transactions`
- **Produtor:** O serviço "Wallet Core" publica eventos de transações no tópico.
- **Consumidor:** O serviço "Balances" consome os eventos publicados para processar e atualizar os saldos das contas.

#### Configuração do Kafka
O Kafka é configurado no `docker-compose.yml` para facilitar a inicialização junto aos serviços. Ele é exposto na porta `9092`.

### Banco de Dados
Ambos os microsserviços utilizam PostgreSQL como banco de dados, cada um com sua instância configurada no Docker Compose.

- **Wallet Core:**
  - Armazena as transações financeiras realizadas.
  - Possui uma tabela para registrar as operações realizadas por cada conta.
- **Balances:**
  - Armazena o saldo consolidado de cada conta, atualizado com base nos eventos consumidos do Kafka.



### Comunicação
Os serviços são expostos pelas seguintes portas:

- Wallet Core: 8080
- Balances: 3003

### Testes e Validação

#### Testando a API do Wallet Core
No diretório wallet-core/api, execute o arquivo client.http usando ferramentas como VS Code REST Client ou Postman.

#### Exemplo de requisição:
```http

POST http://localhost:8080/transactions
Content-Type: application/json

{
  "account_id": 1,
  "amount": 100.0
}
```

#### Testando a API do Balances
No diretório balances/api, execute o arquivo client.http para consultar os saldos.

Exemplo de requisição:
```http

GET http://localhost:3003/balances/1
Resposta esperada:
json
Copiar código
{
  "account_id": 1,
  "balance": 100.0
}
```

#### Automação com Docker Compose
O arquivo docker-compose.yml automatiza a configuração de toda a infraestrutura:

Configura os containers para os serviços Wallet Core, Balances, Kafka, e PostgreSQL.
Executa as migrações e popula os bancos de dados com dados fictícios.
Expõe as portas necessárias para acessar os serviços.
Para iniciar todo o sistema:

```bash

docker-compose up --build
```

Para encerrar:

```bash

docker-compose down

```

## Conceitos de Arquitetura Utilizados
### Arquitetura Orientada a Eventos (Event-Driven Architecture - EDA)
A EDA permite que os serviços sejam independentes e reagentes, processando eventos de forma assíncrona. No contexto deste projeto:

O "Wallet Core" gera eventos quando uma transação é registrada.

O serviço "Balances" consome esses eventos para atualizar os saldos das contas.
CQRS e Event Sourcing
CQRS (Command Query Responsibility Segregation): As operações de escrita e leitura são separadas, garantindo eficiência e escalabilidade.
Event Sourcing: Os eventos são armazenados como uma fonte de verdade, permitindo auditoria e reprodução de estados.
