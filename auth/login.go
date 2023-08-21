package auth

import (
	"github.com/gofiber/fiber/v2"
)

func Login(c *fiber.Ctx) error {

	return c.SendString("Login")
}

func Signup(c *fiber.Ctx) error {
	return c.SendString("Signup" + c.Context().String())
}
