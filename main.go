package main

import (
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()
	app.Get("/v1/users", func(c *fiber.Ctx) error {
		return c.SendString("Get all users")
	})
	app.Get("/v1/users/:id", func(c *fiber.Ctx) error {
		return c.SendString("Get user " + c.Params("id"))
	})
	app.Post("/v1/users/:id", func(c *fiber.Ctx) error {
		return c.SendString("Create user " + c.Params("id"))
	})
	app.Delete("/v1/users/:id", func(c *fiber.Ctx) error {
		return c.SendString("Delete user " + c.Params("id"))
	})
	app.Patch("/v1/users/:id", func(c *fiber.Ctx) error {
		return c.SendString("Update a user " + c.Params("id"))
	})

	app.Listen(":3000")
}
