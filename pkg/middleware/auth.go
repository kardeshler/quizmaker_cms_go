package middleware

import (
	"github.com/gofiber/fiber"
	jwtware "github.com/gofiber/jwt"
)

func Protected() func(c *fiber.Ctx) {
	return jwtware.New(jwtware.Config{
		SigningKey:   []byte("mysecret"),
		ErrorHandler: jwtError,
	})
}

func jwtError(c *fiber.Ctx, err error) {
	if err.Error() == "Missing or malformed JWT" {
		c.Status(fiber.StatusBadRequest)
		c.JSON(fiber.Map{"status": "error", "message": "Missing or malformed JWT", "data": nil})
	} else {
		c.Status(fiber.StatusUnauthorized)
		c.JSON(fiber.Map{"status": "error", "message": "Invalid or expired JWT", "data": nil})
	}
}
