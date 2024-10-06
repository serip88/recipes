package router

import (
	"context"
	"fmt"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/session"

	servicev1 "rain.io/protogen/service/v1"
)

type Router struct {
	ServiceCli     servicev1.AddServiceClient
	Store          *session.Store
	CsrfMiddleware func(*fiber.Ctx) error
}

func New(client servicev1.AddServiceClient, store *session.Store, csrfMiddleware func(*fiber.Ctx) error) *Router {
	fmt.Println("Begin api main router...")
	return &Router{
		ServiceCli:     client,
		Store:          store,
		CsrfMiddleware: csrfMiddleware,
	}
}

// SetupRoutes setup router api
func (p *Router) SetupRoutes(app *fiber.App) {
	// @TODO: Check JS framework if not working csrf and HTTPOnly: true => change to jwt
	// apiGroup := app.Group("/api", p.CsrfMiddleware, logger.New())
	apiGroup := app.Group("/api", logger.New(), p.CtxCheckCsrf)
	apiGroup.Post("/login", p.Login)
	// p.AuthRoutes(app)
}
func (p *Router) CtxCheckCsrf(c *fiber.Ctx) error {
	csrfHeader := c.Get("x-csrf-token")
	csrfCookie := c.Cookies("__Host-csrf")
	//B get session
	session, err := p.Store.Get(c)
	if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}
	csrfSessionI := session.Get("fiber.csrf.token")
	csrfSessionStr := strings.ReplaceAll(fmt.Sprint(csrfSessionI), "{", "")
	csrfSessionStr = strings.ReplaceAll(csrfSessionStr, "}", "")
	csrfSessionStrs := strings.Split(csrfSessionStr, " ")
	csrfSession := csrfSessionStrs[0]
	//E get session
	if csrfSession != csrfHeader {
		return c.SendStatus(fiber.StatusInternalServerError)
	}
	fmt.Println("csrf...csrfCookie.", csrfHeader, csrfSession, csrfCookie)

	return c.Next()
}
func (p *Router) Login(c *fiber.Ctx) error {

	type LoginInput struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}
	var input LoginInput
	if err := c.BodyParser(&input); err != nil { //json value
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "error", "message": "Error on login request", "data": err})
	}
	fmt.Println("BodyParser...", input)

	req := &servicev1.Request{
		Module: servicev1.Module_MODULE_LOGIN,
		User: &servicev1.User{
			Email:    input.Username,
			Password: input.Password,
		},
	}
	user := &servicev1.User{}
	if res, err := p.ServiceCli.GetUser(context.Background(), req); err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"status": "error", "message": "User not found", "data": err})
	} else {
		user = res.User
		if user == nil {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"status": "error", "message": "User not found", "data": err})
		}
		return c.JSON(fiber.Map{"status": "success", "message": "Success login", "data": user})

	}
}
