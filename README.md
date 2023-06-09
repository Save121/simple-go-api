
# Simple go Api 🤓
Este es un proyecto de practica de una API hecha en Go, utilizando el inyector de dependencias de Uber [FX](https://github.com/uber-go/fx), [Echo](https://echo.labstack.com/) como framework web. Y siguiendo la guia del canal de youtube [Go Simplified](https://www.youtube.com/@GoSimplifiedChannel) el cual lo recomiendo encarecidamente si alguien desea aprender a profundidad Go. También intentando aplicar las mejores practicas aprendidas en [Codely](https://codely.com/), los cuales han sido de gran apollo para mejorar mi nivel de programación. Y finalmente intentando seguir la guía de estilos de uber para GO [👉 clic aquiiii 👈](https://github.com/friendsofgo/uber-go-guide-es)

## Donde inicializar 🥳
La base de datos es PostgreSQL, utilizando la versión `15.2`. El archivo `database/docker-compose.yml` contiene la configuración de la base de datos para Docker 🐋. Para iniciarla es necesario utilizar los siguientes comandos:
 ````bash
 cd database
 ````
  ````bash
 docker-compose up
 ````

 En la carpeta `settings/settings.yaml` se deben agregar las variables tanto de la base de datos como del puerto en que correra la aplicación. 

 Para correr los tests se utiliza el siguiente comando:

 ```` bash
go test ./... -v
 ````
 
 Y finalmente 😴, la aplicación se ejecuta con el siguiente comando:
 ```bash	
 go run main.go
 ```

