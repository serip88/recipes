package main

import (
	"log" // Package for logging errors

	"yendoapi/database"
	"yendoapi/router"

	"github.com/gofiber/fiber/v2"                    // Importing the fiber package for handling HTTP requests
	"github.com/gofiber/fiber/v2/middleware/cors"    // Middleware for handling Cross-Origin Resource Sharing (CORS)
	"github.com/gofiber/fiber/v2/middleware/favicon" // Middleware for serving favicon
	"github.com/gofiber/fiber/v2/middleware/logger"  // Middleware for logging HTTP requests
	"github.com/gofiber/template/django/v3"
)

func main() {
	// Create a new engine
	engine := django.New("./views", ".html")

	// Or from an embedded system
	// See github.com/gofiber/embed for examples
	// engine := html.NewFileSystem(http.Dir("./views", ".django"))

	// Pass the engine to the Views
	app := fiber.New(fiber.Config{
		Views: engine,
	})

	// app := fiber.New() // Initialize a new Fiber instance
	// register middlewares
	app.Use(favicon.New()) // Use favicon middleware to serve favicon
	app.Use(cors.New())    // Use CORS middleware to allow cross-origin requests
	app.Use(logger.New())  // Use logger middleware to log HTTP requests

	database.ConnectDB()
	router.SetupRoutes(app)
	router.SetupPageRoutes(app)

	// Define a GET route for the path '/hello'
	app.Get("/hello", hello)
	// 404 Handler
	app.Use(notFound)

	log.Fatal(app.Listen(":3000")) // Start the server on port 5000 and log any errors
}

// Handler
func hello(c *fiber.Ctx) error {
	return c.SendString("I made a ☕ for you!")
}
func notFound(c *fiber.Ctx) error {
	return c.SendStatus(404) // => 404 "Not Found"
}
