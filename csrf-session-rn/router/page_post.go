package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
	"golang.org/x/crypto/bcrypt"
)

func SetupPostRoutes(app *fiber.App, csrfMiddleware func(*fiber.Ctx) error, store *session.Store, users map[string]User, emptyHashString string) {
	// Route for processing the login
	app.Post("/login", csrfMiddleware, func(c *fiber.Ctx) error {
		// Retrieve the submitted form data
		username := c.FormValue("username")
		password := c.FormValue("password")

		// Check if the credentials are valid
		user, exists := users[username]
		var checkPassword string
		if exists {
			checkPassword = user.Password
		} else {
			checkPassword = emptyHashString
		}

		if bcrypt.CompareHashAndPassword([]byte(checkPassword), []byte(password)) != nil {
			// Authentication failed
			csrfToken, ok := c.Locals("csrf").(string)
			if !ok {
				return c.SendStatus(fiber.StatusInternalServerError)
			}

			return c.Render("login", fiber.Map{
				"Title": "Login",
				"csrf":  csrfToken,
				"error": "Invalid credentials",
			})
		}

		// Set a session variable to mark the user as logged in
		session, err := store.Get(c)
		if err != nil {
			return c.SendStatus(fiber.StatusInternalServerError)
		}
		if err := session.Reset(); err != nil {
			return c.SendStatus(fiber.StatusInternalServerError)
		}
		session.Set("loggedIn", true)
		if err := session.Save(); err != nil {
			return c.SendStatus(fiber.StatusInternalServerError)
		}

		// Redirect to the protected route
		return c.Redirect("/protected")
	})

}
