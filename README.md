## API para buscar clima de cidades a partir do CEP

*Observações: o projeto utiliza uma API de terceiros para buscar o clima da cidade. Essa API possui um token de acesso com expiração em 31/01/2024. Caso você tente executar o projeto após essa data, poderá encontrar erros. Farei o possível para manter o projeto atualizado.*

No projeto você vai encontrar um arquivo em *api/get_weather.http* ele pode te auxiliar nos testes caso você use a extenção [Rest Client](https://github.com/Huachao/vscode-restclient/tree/master) no VScode. 

``Teste localmente - SEM DOCKER``

Para rodar o projeto, execute:

Lembre-se de substituir onde está escrito **YOUR_TOKEN** pelo seu token da https://www.weatherapi.com/

```sh
    WEATHER_TOKEN=YOUR_TOKEN go run cmd/server.go
```

Para testar, tente rodar

```sh
    curl http://localhost:8080/80010000
```

``Test localmente - COM DOCKER``


Para rodar o projeto, execute:

Lembre-se de substituir no arquivo docker-compose.yaml a variável YOUR_TOKEN pelo seu token da https://www.weatherapi.com/

```sh
    docker compose up -d
```

Para testar, tente rodar

```sh
    curl http://localhost:8080/80010000
```
    
``Test produção``

O projeto está sendo hospedado pelo Google Cloud Run. É possível acessar o ambiente via https://weather-api-y4majqddoq-uc.a.run.app/80010000

Ou testar com

```sh
    curl https://weather-api-y4majqddoq-uc.a.run.app/80010000
```