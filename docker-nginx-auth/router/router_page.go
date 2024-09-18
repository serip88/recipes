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

	app.Get("/login", func(c *fiber.Ctx) error {
		data := fiber.Map{
			"error":     "Invalid credentials",
			"actionURL": "/api/auth/login", // Thay thế bằng URL và tham số mong muốn
		}
		return c.Render("embed/login", data)
	})

}
