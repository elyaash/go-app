package main

import (
	"elyaash/auth"
	"elyaash/auth/config"
	"elyaash/auth/handlers"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
)

func main() {

	app := fiber.New()

	authApi := app.Group("/auth")
	authApi.Use(auth.Auth())
	authApi.Get("/login", auth.Login)
	authApi.Get("/sign-up", auth.Signup)

	// Note: This is just an example, please use a secure secret key
	jwt := auth.NewAuthMiddleware(config.Secret)
	app.Post("/login", handlers.Login)
	app.Get("/protected", jwt, handlers.Protected)

	//Open route
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Welcome to Go App")
	})

	err := app.Listen(":80")
	if err != nil {
		log.Error(err)
	}
}
