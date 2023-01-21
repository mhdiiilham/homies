package rest

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func PreRequest() fiber.Handler {
	return func(c *fiber.Ctx) error {
		requestID := uuid.New().String()
		c.Locals(RequestIDKey, requestID)
		return c.Next()
	}
}
