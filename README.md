# redis-as-broker-one
Redis como broker Grupo 1
# Instrucciones
```
1.- docker compose up -d
```
2.- Realizar peticion de tipo POST a la siguiente URL [Redis Producer](http://localhost:80/publish)
```
{
    "message" : "Prueba de REDIS con peticion desde POSTMAN"
}
```
# Fin