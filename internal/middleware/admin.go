package middleware

import (
	"github.com/gofiber/fiber/v2"
)

// AdminMiddleware checks if the user has admin role
func AdminMiddleware() fiber.Handler {
	return func(c *fiber.Ctx) error {
		user := c.Locals("user")
		if user == nil {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "Unauthorized",
			})
		}

		// TODO: Check if user has admin role
		// For now, just check if user exists
		return c.Next()
	}
}
