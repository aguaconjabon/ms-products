package main

import (
	"fmt"
	"ms-products/controller/api"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func main() {
	app := fiber.New(fiber.Config{
		ErrorHandler: func(ctx *fiber.Ctx, err error) error {
			// Status code defaults to 500
			code := fiber.StatusInternalServerError
			var msg string
			// Retrieve the custom status code if it's an fiber.*Error
			if e, ok := err.(*fiber.Error); ok {
				code = e.Code
				msg = e.Message
			}

			if msg == "" {
				msg = "cannot process the http call"
			}

			// Send custom error page
			err = ctx.Status(code).JSON(internalError{
				Message: msg,
			})

			return err
		},
	})
	app.Use(
		recover.New(),
		cors.New(cors.Config{AllowOrigins: "*"}))

	api.SetupProductRoutes(app)

	err := app.Listen(":8080")

	if err != nil {
		fmt.Printf("ERROR LISTEN PORT: %v   ERROR: %v \n", ":8080", err)
	}

}

type internalError struct {
	Message string `json:"message"`
}
