package handlers

import (
	"github.com/gofiber/fiber/v2"
)

func Index(c *fiber.Ctx) error {
	return c.SendString("Hello World!")
	// return c.Render("index", fiber.Map{
	// "Title": "Hello, World!",
	// })
}

func Ping(c *fiber.Ctx) error {
	return c.SendString("pong")
}

// NotFound returns custom 404 page
func NotFound(c *fiber.Ctx) error {
	return c.Status(404).SendFile("./static/private/404.html")
}
