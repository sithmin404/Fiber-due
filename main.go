package main

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	app.Get("/user", func(c *fiber.Ctx) error {
		return c.SendString("Hi, user")
	})

	app.Get("/:name", func(c *fiber.Ctx) error {
		name := c.Params("name")
		txt := fmt.Sprintf("Hi, %v", name)
		return c.SendString(txt)
	})

	app.Listen(":3000")
}
