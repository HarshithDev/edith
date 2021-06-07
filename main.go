package main

import (
	database "edith/db"
	"edith/router"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func CreateServer() *fiber.App {
	app := fiber.New()
	return app
}

func main() {

	// connect to db
	database.ConnectToDB()

	app := CreateServer()

	app.Use(cors.New())

	router.SetupRoutes(app)

	// 404 handler
	app.Use(func(c *fiber.Ctx) error {
		return c.SendStatus(404) // 404 "Not Found"
	})

	log.Fatal(app.Listen(":3000"))

}
