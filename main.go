package main

import (
	"elyaash/auth"
	"elyaash/auth/config"
	"elyaash/auth/handlers"
	"fmt"
	"os"
	"time"

	myLog "log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
)

func main() {
	logger := getLogger("system")
	logger.Println("Starting web server...")
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

	err := app.Listen(":3000")
	if err != nil {
		log.Error(err)
	}
	logger.Println("Error occured...", err)
}

func getLogger(name string) *myLog.Logger {
	year, month, day := time.Now().Date()
	LOG_FILE_LOCATION := fmt.Sprintf("./logs/"+name+"-%d-%d-%dd.log", month, day, year)
	file, err := os.OpenFile(LOG_FILE_LOCATION, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0666)
	if err != nil {
		//log.Error("Could not create a log file error\n", err)
		panic(err)
	}
	return myLog.New(file, "", myLog.Ldate|myLog.Ltime|myLog.Lshortfile)
}
