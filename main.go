package main

import (
	a "ms-products/controller/product"

	"github.com/gofiber/fiber"
)

func main() {
	app := fiber.New()
	app.Get("/product/", a.GetProduct)

	app.Listen(8080)
}
