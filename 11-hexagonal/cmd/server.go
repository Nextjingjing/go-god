package main

import (
	"log"

	handler "github.com/Nextjingjing/go-god/11-hexagonal/internal/adapters/handler/http"
	repo "github.com/Nextjingjing/go-god/11-hexagonal/internal/adapters/sqlite"
	"github.com/Nextjingjing/go-god/11-hexagonal/internal/core/services"
	"github.com/gofiber/fiber/v3"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {
	db, err := gorm.Open(sqlite.Open("user.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// Migrate the schema
	err = db.AutoMigrate(&repo.UserModel{})
	if err != nil {
		log.Fatal("Failed to migrate database:", err)
	}

	// Initialize repository, service, and handler
	userRepo := repo.NewUserRepositoryImpl(db)
	userService := services.NewUserServiceImpl(userRepo)
	userHandler := handler.NewUserHandler(userService)

	// Initialize Fiber app
	app := fiber.New()

	// Logging middleware
	app.Use(func(c fiber.Ctx) error {
		log.Println("Request received: " + c.Method() + " " + c.Path())
		return c.Next()
	})

	// Set up routes
	api := app.Group("/api")
	userHandler.UserRoute(api.Group("/users"))

	// 404 Handler
	app.Use(func(c fiber.Ctx) error {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "Route not found",
		})
	})

	// Start the server
	log.Println("Server started on port 3000...")
	err = app.Listen(":3000")
	if err != nil {
		log.Fatal(err)
	}
}
