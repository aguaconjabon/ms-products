# ms-products

docker network create example

sudo docker build -t go-ms-product .

se debe generar un link simbolico para que pueda acceder a la bd de mongo que esta en otro contenedor
`docker run -d --rm -p 8080:8080 --name fiberproduct --link mongodb-local:serverdb go-ms-product`
