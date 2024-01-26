## API para buscar clima de ciadades a partir do cep


*Observações o projeto utiliza de uma API de terceiros para buscar o clima da cidade, essa API possui um token de acesso com expiração para 31/01/2024. Caso você tente executar o projeto apos essa data tenha erros. Vou tentar o possivel para manter o projeto em ar.*

``Test localmente - SEM DOCKER``

Para rodar o projeto execute

Lembre de trocar onde está escrito **YOUR_TOKEN** para seu token da https://www.weatherapi.com/

```sh
    WEATHER_TOKEN=YOUR_TOKEN go run cmd/server.go
```

Para testar tente rodar

```sh
    curl http://localhost:8080/80010000
```

``Test localmente - COM DOCKER``


Para rodar o projeto execute

Lembre de trocar no arquivo docker-compose.yaml a variável **YOUR_TOKEN** para seu token da https://www.weatherapi.com/

```sh
    docker compose up -d
```

Para testar tente rodar

```sh
    curl http://localhost:8080/80010000
```
    
``Test produção``

O projeto está sendo ospedado pelo Google Cloud Run é possivel acessar o ambiente via https://weather-api-y4majqddoq-uc.a.run.app/80010000

Ou testar com

```sh
    curl https://weather-api-y4majqddoq-uc.a.run.app/80010000
```