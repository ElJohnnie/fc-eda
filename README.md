# Curso Full Cycle 3.0 - Módulo sobre Arquitetura Orientada a Eventos (EDA - Event Driven Architecture)

---

## Instruções

Para fazer o `build` de todos serviços no diretório raiz executo o comando `docker compose up --build` com isso as imagens serão baixadas e executadas.

Após todos containers estiverem em rodando, inicie as aplicações:

### Criação de tópicos no Kafka

Para criar o tópico no Kafka, execute o seguinte comando:

Acesse http://localhost:9092 e crie manualmente

### Serviço de Wallet

Execute o comando `docker compose exec wallet-app bash` após acessar o container execute o comando `go run cmd/walletcore/main.go` ele irá rodas as migrations e subir o sevidor na porta `8080`.

### Serviço de Balances

Execute o comando `docker compose exec balances-app bash` após acessar o container execute o comando `go run cmd/balances/main.go` ele irá rodas as migrations e subir o sevidor na porta `3003`.

Os dois serviços já tem seu arquivo de `api/client.http` já com os `IDs` corretos, mas nada impede te criar novos registros e usar os mesmos.
