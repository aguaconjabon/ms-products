# ms-products

# levantar ms-products con docker

Para poder correr este proyecto es necesario tener levantado en uan imagen docker la base dedatos proporcioanda por walmart

`https://github.com/walmartdigital/brand-discounts-db`

- Una vez levantada la imagen de base de dato se debe generar una imagen de este proyecto usando el siguiente comando.

`sudo docker build -t go-ms-product . `

- para levantar el contenedor se debe generar tambien un link simbolico para que el microservicio pueda conectarse al la BD mongo

`docker run -d --rm -p 8080:8080 --name fiberproduct --link mongodb-local:serverdb go-ms-product`

- una vez realizado esto se puede acceder al microservicio desplegado en el puerto :8080

`http://127.0.0.1:8080/product/`

# levantar proyecto mediante terminal

ejecutar
`go run .`

en path del proyecto
