package router

import (
	"csrf-session-rn/router/util"
	"fmt"

	router_api "csrf-session-rn/router/api"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"

	servicev1 "rain.io/protogen/service/v1"
)

type Router struct {
	ServiceCli     servicev1.CommonServiceClient
	Store          *session.Store
	CsrfMiddleware func(*fiber.Ctx) error
}

func New(client servicev1.CommonServiceClient) *Router {
	fmt.Println("Begin new main router...")
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
	Error  string
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
	api := router_api.New(p.ServiceCli, p.Store, p.CsrfMiddleware)
	api.SetupRoutes(app)
	//Set module routes
	p.PublicRoutes(app)
	p.AuthRoutes(app)
	p.ProtectedRoutes(app)
	p.ManageRoutes(app)

}

func (p *Router) HandlePage(c *fiber.Ctx, page Page, fMap fiber.Map) error {

	csrfToken, ok := c.Locals("csrf").(string)
	if !ok {
		return c.SendStatus(fiber.StatusInternalServerError)
	}
	fMap["csrf"] = csrfToken
	fMap["StaticURL"] = "/static/"
	fMap["Title"] = page.Title
	fMap["error"] = page.Error

	if page.Layout == "" {
		return c.Render(page.Page, fMap)
	} else {
		return c.Render(page.Page, fMap, page.Layout) //"layouts/login/index"
	}
}
