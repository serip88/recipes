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
type Page struct {
	Title  string
	Page   string
	Layout string
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

func (p *Router) HandleErrorPage(c *fiber.Ctx, page string, err string, layout string) error {
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
		"error":     err, //"Invalid credentials"
	}
	if layout == "" {
		return c.Render(page, fMap)
	} else {
		return c.Render(page, fMap, layout) //"layouts/login/index"
	}
}
func (p *Router) HandlePage(c *fiber.Ctx, page Page, fMap fiber.Map) error {

	csrfToken, ok := c.Locals("csrf").(string)
	if !ok {
		return c.SendStatus(fiber.StatusInternalServerError)
	}
	fMap["csrf"] = csrfToken
	fMap["Title"] = page.Title
	if page.Layout == "" {
		return c.Render(page.Page, fMap)
	} else {
		return c.Render(page.Page, fMap, page.Layout) //"layouts/login/index"
	}
}
