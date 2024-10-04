package router

import (
	"context"
	"fmt"

	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
	servicev1 "rain.io/protogen/service/v1"
)

func (p *Router) AuthRoutes(app *fiber.App) {

	store := p.Store
	csrfMiddleware := p.CsrfMiddleware
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
		return c.Render("embed", fiber.Map{
			"Title":     "Login",
			"csrf":      csrfToken,
			"StaticURL": "/static/",
		}, "layouts/login/index")
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

		loginPage := Page{
			Title:  "Login",
			Page:   "embed",
			Layout: "layouts/login/index",
			Error:  "Invalid credentials",
		}
		fMap := fiber.Map{}
		user := &servicev1.User{}
		if res, err := p.ServiceCli.GetUser(context.Background(), req); err != nil {
			fmt.Println("Login fails...", err.Error())
			loginPage.Error = err.Error()
			return p.HandlePage(c, loginPage, fMap)
			// return err
		} else {
			user = res.User
			fmt.Println("Res User...", res.User)
			if user == nil {
				loginPage.Error = "User Not Found."
				return p.HandlePage(c, loginPage, fMap)
			}

		}
		fmt.Println("End Post login...")
		//E get user from service
		//B get hard users
		/*users, emptyHashString := HardGetUsers()
		// Check if the credentials are valid
		user, exists := users[username]
		var checkPassword string
		if exists {
			checkPassword = user.Password
		} else {
			checkPassword = emptyHashString
		}*/

		if bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)) != nil {
			// Authentication failed
			loginPage.Error = "Wrong Password."
			return p.HandlePage(c, loginPage, fMap)

		}
		fmt.Println("Check password success...")
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
		return c.Redirect("/manage")
	})

}

func HardGetUsers() (map[string]User, string) {
	//B Hard code password
	// Never hardcode passwords in production code
	hashedPasswords := make(map[string]string)
	for username, password := range map[string]string{
		"user1": "password1",
		"user2": "password2",
	} {
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), 10)
		if err != nil {
			panic(err)
		}
		hashedPasswords[username] = string(hashedPassword)
	}

	// Used to help prevent timing attacks
	emptyHash, err := bcrypt.GenerateFromPassword([]byte(""), 10)
	if err != nil {
		panic(err)
	}
	emptyHashString := string(emptyHash)

	users := make(map[string]User)
	for username, hashedPassword := range hashedPasswords {
		users[username] = User{Username: username, Password: hashedPassword}
	}
	//E Hard code password

	return users, emptyHashString
}
