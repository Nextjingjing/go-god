package routes

import "github.com/gofiber/fiber/v3"

func SomeRoute(route fiber.Router) {
	route.Get("/", func(c fiber.Ctx) error {
		return c.SendString("Hello from some route!")
	})

	route.Get("/2", func(c fiber.Ctx) error {
		return c.SendString("Hello2 from some route!")
	})

	route.Get("/3", func(c fiber.Ctx) error {
		return c.SendString("Hello3 from some route!")
	})
}
