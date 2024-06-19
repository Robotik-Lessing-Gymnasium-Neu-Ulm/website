package web

import (
	"github.com/gofiber/fiber/v2"
)

func checkAuth(c *fiber.Ctx) bool {
	return c.OriginalURL() != "/monitor"
}
