package middleware

import (
	"strings"
	"yourapp/pkg/auth"

	"github.com/gofiber/fiber/v2"
)

// AuthMiddleware handles JWT authentication
func AuthMiddleware() fiber.Handler {
	return func(c *fiber.Ctx) error {
		authHeader := c.Get("Authorization")
		if authHeader == "" {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "Authorization header is required",
			})
		}

		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "Invalid authorization header format",
			})
		}

		tokenStr := parts[1]
		manager := auth.GetJWTManager()
		claims, err := manager.ValidateToken(tokenStr)
		if err != nil {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "Invalid or expired token",
			})
		}

		// Store full claims in context
		c.Locals("jwt_claims", claims)
		return c.Next()
	}
}

// GetClaims returns JWT claims from context
func GetClaims(c *fiber.Ctx) *auth.Claims {
	if claims, ok := c.Locals("jwt_claims").(*auth.Claims); ok {
		return claims
	}
	return nil
}

// GetUserID returns user ID from context
func GetUserID(c *fiber.Ctx) uint64 {
	claims := GetClaims(c)
	if claims != nil {
		return claims.UserID
	}
	return 0
}
