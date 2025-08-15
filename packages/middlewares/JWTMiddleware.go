package middlewares

import (
	"strings"

	"github.com/gofiber/fiber/v2"
)

func AuthMiddleware() fiber.Handler {
	return func(c *fiber.Ctx) error {
		// Get Authorization header
		auth := c.Get("Authorization")
		
		// Check if Authorization header exists
		if auth == "" {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"status":  fiber.StatusUnauthorized,
				"message": "Missing authorization header",
			})
		}

		// Split Bearer token
		parts := strings.Split(auth, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"status":  fiber.StatusUnauthorized,
				"message": "Invalid authorization format",
			})
		}

		// Get the token
		token := parts[1]

		// Verify token against your API key or JWT secret
		if token == "" {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"status":  fiber.StatusUnauthorized,
				"message": "Invalid token",
			})
		}

		// Store token in context for later use if needed
		c.Locals("token", token)
		
		return c.Next()
	}
}