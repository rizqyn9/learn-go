package server

import "github.com/gofiber/fiber/v2"

var name = "server-service"

func NewServer() *fiber.App {
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hi Fiber")
	})
	return app
}
