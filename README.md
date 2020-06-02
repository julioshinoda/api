# api



>Antes de subir esse projeto é necessário ter instalado
>- docker 
>- docker-compose
>- make


## Stack utilizada
- Go 1.13
- Postgres 

## Utilizar a api
Para utilizar a API execute os seguintes comandos

Na primeira vez:

```docker-compose -f docker-compose.all.yml build```  -> esse comando ira buidar o container com o Go

``` docker-compose -f docker-compose.all.yml up ``` ->  esse comando ira subir o postres e o container com a API
``` make migration ``` -> roda a migration com a estrutura inicial do banco

Logo após os comandos acima, é possível fazer as requests:

1. Criação de uma conta
```curl -d '{"document_number" :"12343212312"}' -H 'Content-type: application/json' http://localhost:9876/accounts```  

2. Consulta de informações de uma conta
```curl http://localhost:9876/accounts/:accountId```

3. Criação de uma transação
```curl -d '{"account_id":15,"operation_type_id":4,"amount": 13.25}' -H 'Content-type: application/json' http://localhost:9876/transactions ```

