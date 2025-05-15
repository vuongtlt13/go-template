package middleware

import (
	"github.com/gofiber/fiber/v2"
)

// RequestMiddleware stores the request in the context
var RequestMiddleware = func(c *fiber.Ctx) error {
	// Store request in context
	c.Locals("request", c.Request())

	return c.Next()
}
