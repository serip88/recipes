package router

import (
	"github.com/gofiber/fiber/v2"
)

// SetupRoutes setup router api
func SetupPageRoutes(app *fiber.App) {
	app.Get("/embed", func(c *fiber.Ctx) error {
		// Render index within layouts/main
		return c.Render("embed", fiber.Map{
			"Title": "Hello, World!!",
		}, "layouts/main2")
	})

	app.Get("/", func(c *fiber.Ctx) error {
		// Render index within layouts/main
		return c.Render("index", fiber.Map{
			"Title": "Hello, World!!",
		}, "layouts/index")
	})
}
