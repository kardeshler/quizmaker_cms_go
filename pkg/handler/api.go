package handler

import "github.com/gofiber/fiber"

// Hello api health check
func Hello(c *fiber.Ctx) {
	c.JSON(fiber.Map{"status": "success", "message": "Hello i'm ok!", "data": nil})
}
