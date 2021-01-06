package handler

import (
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber"
)

// Login to login with username and password todo: what else can be done instead of username:password?
func Login(c *fiber.Ctx) {
	type LoginInput struct {
		Identity string `json:"identity"`
		Password string `json:"password"`
	}
	var input LoginInput
	if err := c.BodyParser(&input); err != nil {
		c.SendStatus(fiber.StatusUnauthorized)
		return
	}
	identity := input.Identity
	pass := input.Password
	if identity != "kardeshler" || pass != "1234" {
		c.SendStatus(fiber.StatusUnauthorized)
		return
	}

	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["identity"] = identity
	claims["admin"] = true
	claims["exp"] = time.Now().Add(time.Hour * 72).Unix()

	t, err := token.SignedString([]byte("mysecret"))
	if err != nil {
		c.SendStatus(fiber.StatusInternalServerError)
		return
	}
	c.JSON(fiber.Map{"status": "success", "message": "Success login", "data": t})
}
