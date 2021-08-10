package api

import (
	connBD "ms-products/conections"

	"github.com/gofiber/fiber/v2"
)

const BD = "desafio_walmart"

func SearchProductHandler(c *fiber.Ctx) error {
	collection, err := connBD.GetMongoDbCollection(BD, "products")
	if err != nil {
		return fiber.NewError(500, "cannot bring products")
	}

	return getData(c, collection)

}

func SearchProductDiscountsHandler(c *fiber.Ctx) error {
	collection, err := connBD.GetMongoDbCollection(BD, "discounts")

	if err != nil {
		return fiber.NewError(500, "cannot bring products discounts")
	}

	return getData(c, collection)

}
