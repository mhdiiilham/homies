package rest

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
)

func RootHandler(r fiber.Router) {
	r.Get("/", get())
}

func get() fiber.Handler {
	return func(c *fiber.Ctx) error {
		requestId := c.Locals(RequestIDKey).(string)

		resp := NewSuccessResponse(http.StatusOK, map[string]any{
			"name": "homies API",
			"request": map[string]any{
				"id": requestId,
				"ip": c.IP(),
			},
		})

		return c.Status(resp.Code).JSON(resp)
	}
}
