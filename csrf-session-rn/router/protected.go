package router

import (
	"github.com/gofiber/fiber/v2"
)

func (p *Router) ProtectedRoutes(app *fiber.App) {
	store := p.Store
	csrfMiddleware := p.CsrfMiddleware
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
		page := Page{
			Title:  "Protected",
			Page:   "protected",
			Layout: "",
		}
		fMap := fiber.Map{}
		return p.HandlePage(c, page, fMap)
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

		// Retrieve the submitted form data
		message := c.FormValue("message")
		page := Page{
			Title:  "Protected",
			Page:   "protected",
			Layout: "",
		}
		fMap := fiber.Map{
			"Title":   "Protected",
			"message": message,
		}
		return p.HandlePage(c, page, fMap)
	})
}
