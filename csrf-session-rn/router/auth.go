package router

import (
	"context"
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
	"golang.org/x/crypto/bcrypt"
	servicev1 "rain.io/protogen/service/v1"
)

func (p *Router) AuthRoutes(app *fiber.App, csrfMiddleware func(*fiber.Ctx) error, store *session.Store, users map[string]User, emptyHashString string) {

	// Route for the login page
	app.Get("/login", csrfMiddleware, func(c *fiber.Ctx) error {
		//B Check if the user is logged in
		session, err := store.Get(c)
		if err != nil {
			return c.SendStatus(fiber.StatusInternalServerError)
		}
		loggedIn, _ := session.Get("loggedIn").(bool)
		if loggedIn {
			// User is authenticated, redirect to the home page
			return c.Redirect("/")
		}
		//E Check if the user is logged in
		csrfToken, ok := c.Locals("csrf").(string)
		if !ok {
			return c.SendStatus(fiber.StatusInternalServerError)
		}
		return c.Render("login", fiber.Map{
			"Title": "Login",
			"csrf":  csrfToken,
		})
	})

	// Route for logging out
	app.Get("/logout", func(c *fiber.Ctx) error {
		// Retrieve the session
		session, err := store.Get(c)
		if err != nil {
			return c.SendStatus(fiber.StatusInternalServerError)
		}

		// Revoke users authentication
		if err := session.Destroy(); err != nil {
			return c.SendStatus(fiber.StatusInternalServerError)
		}

		// Redirect to the login page
		return c.Redirect("/login")
	})
	// Route for processing the login
	app.Post("/login", csrfMiddleware, func(c *fiber.Ctx) error {
		// Retrieve the submitted form data
		username := c.FormValue("username")
		password := c.FormValue("password")
		//B get user from service
		fmt.Println("username.password...", username, password)
		fmt.Println("Start Post login...")
		req := &servicev1.Request{
			Module: servicev1.Module_MODULE_LOGIN,
			User: &servicev1.User{
				Email:    username,
				Password: password,
			},
		}
		if res, err := p.ServiceCli.GetUser(context.Background(), req); err == nil {
			fmt.Println("Res User...", res.User)
		} else {
			fmt.Println("Login fails...", err.Error())
		}
		fmt.Println("End Post login...")
		//E get user from service
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
