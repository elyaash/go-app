package auth

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/basicauth"
	jwtware "github.com/gofiber/jwt/v3"
)

//Basic Auth
func Auth() fiber.Handler {
	return basicauth.New(basicauth.Config{
		Users: map[string]string{
			"john":  "doe",
			"admin": "123456",
		},
	})
}

// Middleware JWT function
func NewAuthMiddleware(secret string) fiber.Handler {
	return jwtware.New(jwtware.Config{
		SigningKey: []byte(secret),
	})
}
