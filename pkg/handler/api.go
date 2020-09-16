package handler

import "github.com/gofiber/fiber"

func Hello(c *fiber.Ctx) {
	c.JSON(fiber.Map{"status": "success", "message": "Hello i'm ok!", "data": nil})
}
