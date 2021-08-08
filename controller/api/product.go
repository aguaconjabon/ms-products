package api

import (
	"context"

	connBD "ms-products/conections"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
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

func getData(c *fiber.Ctx, collection *mongo.Collection) error {

	var filter bson.M = bson.M{}

	if c.Params("id") != "" {
		id := c.Params("id")
		objID, _ := primitive.ObjectIDFromHex(id)
		filter = bson.M{"_id": objID}
	}

	var results []bson.M
	cur, err := collection.Find(context.Background(), filter)
	defer cur.Close(context.Background())

	if err != nil {
		return fiber.NewError(500, "cannot bring products")
	}

	cur.All(context.Background(), &results)

	if results == nil {
		c.SendStatus(404)
		return fiber.NewError(500, "cannot bring products")
	}

	return c.JSON(results)

}
