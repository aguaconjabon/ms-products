package api

import (
	"github.com/gofiber/fiber/v2"
)

func SetupProductRoutes(app *fiber.App) {
	grp := app.Group("/product")
	grp.Get("/", SearchProductHandler)
	grp.Get("/discount", SearchProductDiscountsHandler)

}
