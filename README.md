## API para buscar clima de ciadades a partir do cep

``Test localmente``

*Observações o projeto utiliza de uma API de terceiros para buscar o clima da cidade, essa API possui um token de acesso com expiração para 31/01/2024. Caso você tente rodar o projeto depois dessa data entre em https://www.weatherapi.com/ gere seu token e subistitua no arquivo get_weather_by_location.go o valor key da url*

Para rodar o projeto execute

```sh
    go run cmd/server.go
```

Para testar tente rodar

```sh
    curl http://localhost:8000/01001000
```

    
``Teste deployado``

Acesse a seguinte url

https://weather-api-y4majqddoq-uc.a.run.app/30120060