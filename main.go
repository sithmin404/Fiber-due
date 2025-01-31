package main

import (
	"fmt"

	"github.com/gofiber/fiber/v2"

	"github.com/gofiber/template/html"
)

func main() {
	app := fiber.New(fiber.Config{
		Views: html.New("./views", ".html"),
	})

	app.Static("/", "./public")

	app.Get("/", func(c *fiber.Ctx) error {
		return c.Render("index", fiber.Map{})
	})

	app.Get("/user", func(c *fiber.Ctx) error {
		return c.SendString("Hi, user")
	})

	app.Get("/car", func(c *fiber.Ctx) error {
		model := c.Query("model")
		brand := c.Query("brand")
		year := c.Query("year")
		return c.JSON(&fiber.Map{
			"brand": brand,
			"model": model,
			"year":  year,
		})
	})

	app.Get("/:name", func(c *fiber.Ctx) error {
		name := c.Params("name")
		txt := fmt.Sprintf("Hi, %v", name)
		return c.SendString(txt)
	})

	app.Listen(":3000")
}
