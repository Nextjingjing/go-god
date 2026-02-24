package main

import (
	"log"

	"github.com/gofiber/fiber/v3"
)

func main() {
	app := fiber.New()

	app.Get("/", func(c fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	err := app.Listen(":8080")
	if err != nil {
		log.Fatalf("Failed to open server at :8080: %v", err)
	}
}
