# redis-as-broker-one
Redis como broker Grupo 1
# Instrucciones
```
docker compose up -d
```
Realizar peticion de tipo POST a la siguiente URL [Redis Producer](http://localhost:80/publish)
```
{
    "message" : "Prueba de REDIS con peticion desde POSTMAN"
}
```