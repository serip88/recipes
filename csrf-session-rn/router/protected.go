package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
)

func ProtectedRoutes(app *fiber.App, csrfMiddleware func(*fiber.Ctx) error, store *session.Store) {

	// Route for the protected content
	app.Get("/protected", csrfMiddleware, func(c *fiber.Ctx) error {
		//B Check if the user is logged in
		session, err := store.Get(c)
		if err != nil {
			return c.SendStatus(fiber.StatusInternalServerError)
		}
		loggedIn, _ := session.Get("loggedIn").(bool)
		if !loggedIn {
			// User is not authenticated, redirect to the login page
			return c.Redirect("/login")
		}
		//E Check if the user is logged in

		csrfToken, ok := c.Locals("csrf").(string)
		if !ok {
			return c.SendStatus(fiber.StatusInternalServerError)
		}

		return c.Render("protected", fiber.Map{
			"Title": "Protected",
			"csrf":  csrfToken,
		})
	})

	// Route for processing the protected form
	app.Post("/protected", csrfMiddleware, func(c *fiber.Ctx) error {
		// Check if the user is logged in
		session, err := store.Get(c)
		if err != nil {
			return c.SendStatus(fiber.StatusInternalServerError)
		}
		loggedIn, _ := session.Get("loggedIn").(bool)

		if !loggedIn {
			// User is not authenticated, redirect to the login page
			return c.Redirect("/login")
		}

		csrfToken, ok := c.Locals("csrf").(string)
		if !ok {
			return c.SendStatus(fiber.StatusInternalServerError)
		}

		// Retrieve the submitted form data
		message := c.FormValue("message")

		return c.Render("protected", fiber.Map{
			"Title":   "Protected",
			"csrf":    csrfToken,
			"message": message,
		})
	})
}
