package router

import (
	"csrf-session-rn/router/util"
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"

	servicev1 "rain.io/protogen/service/v1"
)

type Router struct {
	ServiceCli     servicev1.AddServiceClient
	Store          *session.Store
	CsrfMiddleware func(*fiber.Ctx) error
}

func New(client servicev1.AddServiceClient) *Router {
	store := util.InitSessionStore()
	csrfMiddleware := util.MakeCsrf(store)
	return &Router{
		ServiceCli:     client,
		Store:          store,
		CsrfMiddleware: csrfMiddleware,
	}
}

// User represents a user in the dummy authentication system
type User struct {
	Username string
	Password string
}

// SetupRoutes setup router api
func (p *Router) SetupRoutes(app *fiber.App) {

	app.Static("/static", "./static")
	// Route for the root path
	app.Get("/", func(c *fiber.Ctx) error {
		// render the root page as HTML
		return c.Render("index", fiber.Map{
			"Title": "Index",
		})
	})
	//Set module routes
	p.AuthRoutes(app)
	p.ProtectedRoutes(app)
	p.ManageRoutes(app)

}

func (p *Router) HandleErrorPage(c *fiber.Ctx, page string, msg string, layout string) error {
	// Authentication failed
	csrfToken, ok := c.Locals("csrf").(string)
	if !ok {
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	fmt.Println("Check password fails...")
	fMap := fiber.Map{
		"Title":     "Login",
		"csrf":      csrfToken,
		"StaticURL": "/static/",
		"error":     msg, //"Invalid credentials"
	}
	if layout == "" {
		return c.Render(page, fMap)
	} else {
		return c.Render(page, fMap, layout) //"layouts/login/index"
	}
}
