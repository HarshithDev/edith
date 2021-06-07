package router

import (
	"github.com/gofiber/fiber/v2"
)

// User handles all user routes
var USER fiber.Router

func hello(c *fiber.Ctx) error {
	return c.SendString("Hello, World!")
}

// SetupRoutes sets up all routes
func SetupRoutes(app *fiber.App) {
	api := app.Group("/api")

	// defining user route
	USER = api.Group("/user")
	SetupUserRoutes()

	// for testing purpose
	api.Get("/", hello)
}
