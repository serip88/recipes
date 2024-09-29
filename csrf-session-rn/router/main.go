package router

import (
	"csrf-session-rn/router/util"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"

	servicev1 "rain.io/protogen/service/v1"
)

type Router struct {
	ServiceCli servicev1.AddServiceClient
	Store      *session.Store
}

func New(client servicev1.AddServiceClient) *Router {
	store := util.InitSessionStore()
	return &Router{
		ServiceCli: client,
		Store:      store,
	}
}

// User represents a user in the dummy authentication system
type User struct {
	Username string
	Password string
}

// SetupRoutes setup router api
func (p *Router) SetupRoutes(app *fiber.App) {

	csrfMiddleware := util.MakeCsrf(p.Store)

	// Route for the root path
	app.Get("/", func(c *fiber.Ctx) error {
		// render the root page as HTML
		return c.Render("index", fiber.Map{
			"Title": "Index",
		})
	})
	//Set module routes
	p.AuthRoutes(app, csrfMiddleware)
	p.ProtectedRoutes(app, csrfMiddleware)

}
