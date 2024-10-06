package router

import (
	"github.com/gofiber/fiber/v2"
	servicev1 "rain.io/protogen/service/v1"
)

func (p *Router) PublicRoutes(app *fiber.App) {
	csrfMiddleware := p.CsrfMiddleware
	// Route for the public page
	app.Post("/public", csrfMiddleware, GetUser)

}

func GetUser(c *fiber.Ctx) error {
	// id := c.Params("id")
	// db := database.DB
	user := &servicev1.User{
		Email:    "username",
		Password: "password",
	}
	// db.Find(&user, id)
	if user.Email == "" {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "No user found with ID", "data": nil})
	}
	return c.JSON(fiber.Map{"status": "success", "message": "User found", "data": user})
}
