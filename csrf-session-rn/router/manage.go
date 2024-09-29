package router

import (
	"github.com/gofiber/fiber/v2"
)

func (p *Router) ManageRoutes(app *fiber.App) {
	store := p.Store
	csrfMiddleware := p.CsrfMiddleware
	// Route for the manage content
	app.Get("/manage", csrfMiddleware, func(c *fiber.Ctx) error {
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

		return c.Render("manage", fiber.Map{
			"Title": "Manage",
			"csrf":  csrfToken,
		}, "layouts/manage")
	})

	// Route for processing the manage form
	app.Post("/manage", csrfMiddleware, func(c *fiber.Ctx) error {
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

		return c.Render("manage", fiber.Map{
			"Title":   "Manage",
			"csrf":    csrfToken,
			"message": message,
		})
	})
}
