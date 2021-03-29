# Operación Fuego Quasar

_Proyecto de evaluación para Mercado Libre_

El algoritmo que encuentra las coordenadas de la nave se creó a partir de la fórmula de Trilateración como se indica en esta fuente: https://math.stackexchange.com/questions/884807/find-x-location-using-3-known-x-y-location-using-trilateration

## Comenzando 🚀

_Estas instrucciones te permitirán obtener una copia del proyecto en funcionamiento en tu máquina local para propósitos de desarrollo y pruebas._

Mira **Deployment** para conocer como desplegar el proyecto.


## Pre-requisitos 📋

_Debes tener instalado Docker y Docker compose para poder levantar el proyecto en tu máquina_

## Instalación 🔧

_Clonamos el proyecto_
```
git clone git@github.com:lmartinezsch/operacion-fuego-quasar.git
```

_Una vez clonado y estando dentro del proyecto ejecutamos:_

```
docker-compose up -d --build
```

_Esto nos levantará 3 contenedores: ofq-app, ofq-mysql y ofq-phpmyadmin_

## REST API 🔧

### PING
Podemos corroborar si nuestra app está funcionando haciendo un GET:  

```
http://localhost:4000/api/v1.0/ping/
```

### Registrar usuario
_Para poder utilizar la api debemos crear un usuario haciendo un POST con el siguiente request:_  
`POST http://localhost:4000/api/v1.0/auth/register`  
`Body:`   

```
{
  "username": "admin",
  "password": "admin",
  "display_name": "OFQ"
}
```

### Loguear usuario
`POST http://localhost:4000/api/v1.0/auth/login`
`BODY`
```{ "username": "admin", "password": "admin" }```

### Agregar satelites
_Se deben agregar 3 satélites para el correcto funcionamiento:_  

`POST http://localhost:4000/api/v1.0/satellites`

```
{
  "name": "kenobi",
  "position": {
    "x": -500,
    "y": -200
  }
}
```  
`POST http://localhost:4000/api/v1.0/satellites`  
```
{
  "name": "Skywalker",
  "position": {
    "x": 100,
    "y": -100
  }
}
```  
`POST http://localhost:4000/api/v1.0/satellites`  
```
{
  "name": "Sato",
  "position": {
    "x": 500,
    "y": 100
  }
}
```

## Top Secret
_Se puede obtener las coordenadas de la nave y el mensaje secreto en el endpoint `topsecret`_  
`POST http://localhost:4000/api/v1.0/topsecret`  
Request Body:
```
{
    "satellites": [
        {
            "name": "kenobi",
            "distance": 485.41,
            "message": [
                "este",
                "",
                "",
                "mensaje",
                ""
            ]
        },
        {
            "name": "skywalker",
            "distance": 265.75,
            "message": [
                "",
                "es",
                "",
                "",
                "secreto"
            ]
        },
        {
            "name": "sato",
            "distance": 600.52,
            "message": [
                "este",
                "",
                "un",
                "",
                ""
            ]
        }
    ]
}
```
Response body:  
```
{
    "position": {
        "x": -100,
        "y": 75
    },
    "message": "este es un mensaje secreto"
}
```

### Top secret split CREATE
_Se pueden setear los contactos a los satelites separados_  

`POST http://localhost:4000/api/v1.0/topsecret_split/Kenobi`  
Request Body:  
```
{
    "name": "kenobi",
    "distance": 485.41,
    "message": [
        "este",
        "",
        "",
        "mensaje",
        ""
    ]
}
```

`POST http://localhost:4000/api/v1.0/topsecret_split/Skywalker`  
Request Body:  
```
{
    "name": "Skywalker",
    "distance": 265.75,
    "message": [
        "",
        "es",
        "",
        "",
        "secreto"
    ]
}
```

`POST http://localhost:4000/api/v1.0/topsecret_split/Sato`  
Request Body:  
```
{
    "name": "Sato",
    "distance": 600.52,
    "message": [
        "este",
        "",
        "un",
        "",
        ""
    ]
}
```

### Top secret split GET
_Se obtienen las coordenadas y el mensaje secreto de la nave_  
`GET http://localhost:4000/api/v1.0/topsecret_split/Sato`
Response body:  
```
{
    "position": {
        "x": -100,
        "y": 75
    },
    "message": "este es un mensaje secreto"
}
```


## Ejecutando las pruebas ⚙️

_Explica como ejecutar las pruebas automatizadas para este sistema_

### Analice las pruebas end-to-end 🔩

_Explica que verifican estas pruebas y por qué_

```
Da un ejemplo
```

### Y las pruebas de estilo de codificación ⌨️

_Explica que verifican estas pruebas y por qué_

```
Da un ejemplo
```
