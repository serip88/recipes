package router

import (
	"csrf-session-rn/router/util"

	"github.com/gofiber/fiber/v2"

	"golang.org/x/crypto/bcrypt"
)

// User represents a user in the dummy authentication system
type User struct {
	Username string
	Password string
}

// SetupRoutes setup router api
func SetupRoutes(app *fiber.App) {
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

	store := util.InitSessionStore()
	csrfMiddleware := util.MakeCsrf(store)

	// Route for the root path
	app.Get("/", func(c *fiber.Ctx) error {
		// render the root page as HTML
		return c.Render("index", fiber.Map{
			"Title": "Index",
		})
	})
	//Set module routes
	AuthRoutes(app, csrfMiddleware, store, users, emptyHashString)
	ProtectedRoutes(app, csrfMiddleware, store)

}
