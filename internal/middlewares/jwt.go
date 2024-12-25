package middleware

import (
	"anywhere-api/pkg/helper"
	"anywhere-api/pkg/jwt"
	"strings"

	"github.com/gofiber/fiber/v2"
)

// Protect is a middleware to protect routes with JWT token validation
func Protect(c *fiber.Ctx) error {
	// Get the token from Authorization header
	tokenString := c.Get("Authorization")
	if tokenString == "" {
		return helper.ErrorResponse(c, "Authorization token required", "Missing token")
	}

	// Strip "Bearer " prefix from token
	tokenString = strings.TrimPrefix(tokenString, "Bearer ")

	// Validate the JWT token
	claims, err := jwt.ValidateJWT(tokenString)
	if err != nil {
		return helper.ErrorResponse(c, "Invalid token", err.Error())
	}

	// Set user context from JWT claims
	c.Locals("username", claims.Username)

	return c.Next()
}
