package main

import (
	"log"

	"github.com/Nextjingjing/go-god/10-fiber/middlewares"
	"github.com/Nextjingjing/go-god/10-fiber/routes"
	"github.com/gofiber/fiber/v3"
)

func main() {
	app := fiber.New()

	// Middleware example 1
	// Logs each request method and path
	// Useful for debugging, monitoring, Error handling, etc.
	app.Use(func(c fiber.Ctx) error {
		// Log each request
		log.Println("Request received: " + c.Method() + " " + c.Path())
		return c.Next()
	})

	app.Get("/", func(c fiber.Ctx) error {
		return c.SendString("Hello, World!")
		// => Response: "Hello, World!"
	})

	app.Post("/api", func(c fiber.Ctx) error {
		return c.SendString("POST request received")
		// => Response: "POST request received"
	})

	// Path parameter example
	app.Get("/api/user/:id", func(c fiber.Ctx) error {
		id := c.Params("id")
		return c.SendString("User ID: " + id)
		// => Response: "User ID: <id>"
	})

	// Query parameter example
	app.Get("/api/search", func(c fiber.Ctx) error {
		query := c.Query("q")
		return c.SendString("Search query: " + query)
		// => Response: "Search query: <q>"
	})

	// Request body parsing example
	type User struct {
		Name  string `json:"name"`
		Email string `json:"email"`
	}
	app.Post("/api/user", func(c fiber.Ctx) error {
		var user User
		err := c.Bind().Body(&user)
		if err != nil {
			return err
		}
		return c.JSON(user)
	})

	// Response example with JSON Using fiber.Map
	app.Get("/api/json", func(c fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"message": "Hello from JSON response",
			"status":  "success",
		})
	})

	// Response example with JSON Using struct
	type Response struct {
		Message string `json:"message"`
		Status  string `json:"status"`
	}

	app.Get("/api/struct", func(c fiber.Ctx) error {
		return c.JSON(Response{
			Message: "Hello from struct response",
			Status:  "success",
		})
	})

	// Middleware example 2
	// Useful for authentication, logging, etc.
	app.Get("/api/middleware", middlewares.MiddlewareExample, func(c fiber.Ctx) error {
		return c.SendString("Middleware example route")
	})

	// Error handling example
	app.Get("/api/error", func(c fiber.Ctx) error {
		return fiber.NewError(fiber.StatusBadRequest, "This is a bad request example")
	})

	// --------------------------------------------------------------------------------------------------
	// routes
	api := app.Group("/api")
	routes.SomeRoute(api.Group("/some-route"))

	// 404 Not Found handler
	app.Use(func(c fiber.Ctx) error {
		return c.Status(404).SendString("404 - Not Found")
	})

	// Start the server
	log.Println("Server started on port 3000...")
	err := app.Listen(":3000")
	if err != nil {
		log.Fatal(err)
	}
}
