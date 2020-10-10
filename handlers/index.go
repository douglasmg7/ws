package handlers

import (
	"github.com/gofiber/fiber/v2"
)

func Index(c *fiber.Ctx) error {
	return c.SendString("Hello World!")
}

func Ping(c *fiber.Ctx) error {
	return c.SendString("pong")
}

// NotFound returns custom 404 page
func NotFound(c *fiber.Ctx) error {
	// return c.Status(404).SendFile("./static/private/404.html")
	// return c.SendStatus(404)
	return c.Status(404).SendString("Página não encontrada")
}

// func InternalError(c *fiber.Ctx) bool {
// c.Status(500).SendString("Alguma coisa deu errado, tente mais tarde!")
// return true
// }

// Default error handler
var ErrorHandlerProd = func(c *fiber.Ctx, err error) error {
	// Default 500 statuscode
	code := fiber.StatusInternalServerError

	if e, ok := err.(*fiber.Error); ok {
		// Override status code if fiber.Error type
		code = e.Code
	}
	// Set Content-Type: text/plain; charset=utf-8
	c.Set(fiber.HeaderContentType, fiber.MIMETextPlainCharsetUTF8)

	// Return statuscode with error message
	return c.Status(code).SendString("Alguma coisa deu errado, tente mais tarde!")
}

// Default error handler
var ErrorHandler = func(c *fiber.Ctx, err error) error {
	// Default 500 statuscode
	code := fiber.StatusInternalServerError

	if e, ok := err.(*fiber.Error); ok {
		// Override status code if fiber.Error type
		code = e.Code
	}
	// Set Content-Type: text/plain; charset=utf-8
	c.Set(fiber.HeaderContentType, fiber.MIMETextPlainCharsetUTF8)

	// Return statuscode with error message
	return c.Status(code).SendString(err.Error())
}
