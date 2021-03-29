# Operación Fuego Quasar

_Proyecto de evaluación para Mercado Libre_

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

## Despliegue 📦

_Agrega notas adicionales sobre como hacer deploy_

## Construido con 🛠️

_Menciona las herramientas que utilizaste para crear tu proyecto_

* [Dropwizard](http://www.dropwizard.io/1.0.2/docs/) - El framework web usado
* [Maven](https://maven.apache.org/) - Manejador de dependencias
* [ROME](https://rometools.github.io/rome/) - Usado para generar RSS

## Contribuyendo 🖇️

Por favor lee el [CONTRIBUTING.md](https://gist.github.com/villanuevand/xxxxxx) para detalles de nuestro código de conducta, y el proceso para enviarnos pull requests.

## Wiki 📖

Puedes encontrar mucho más de cómo utilizar este proyecto en nuestra [Wiki](https://github.com/tu/proyecto/wiki)

## Versionado 📌

Usamos [SemVer](http://semver.org/) para el versionado. Para todas las versiones disponibles, mira los [tags en este repositorio](https://github.com/tu/proyecto/tags).

## Autores ✒️

_Menciona a todos aquellos que ayudaron a levantar el proyecto desde sus inicios_

* **Andrés Villanueva** - *Trabajo Inicial* - [villanuevand](https://github.com/villanuevand)
* **Fulanito Detal** - *Documentación* - [fulanitodetal](#fulanito-de-tal)

También puedes mirar la lista de todos los [contribuyentes](https://github.com/your/project/contributors) quíenes han participado en este proyecto. 

## Licencia 📄

Este proyecto está bajo la Licencia (Tu Licencia) - mira el archivo [LICENSE.md](LICENSE.md) para detalles

## Expresiones de Gratitud 🎁

* Comenta a otros sobre este proyecto 📢
* Invita una cerveza 🍺 o un café ☕ a alguien del equipo. 
* Da las gracias públicamente 🤓.
* etc.



---
⌨️ con ❤️ por [Villanuevand](https://github.com/Villanuevand) 😊
