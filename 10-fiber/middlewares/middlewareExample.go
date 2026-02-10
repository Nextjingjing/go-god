package middlewares

import (
	"log"

	"github.com/gofiber/fiber/v3"
)

func MiddlewareExample(c fiber.Ctx) error {
	// Example middleware function
	log.Println("Middleware executed")
	return c.Next()
}
