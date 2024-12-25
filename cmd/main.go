package main

import (
	handlers "anywhere-api/internal/handler"
	middleware "anywhere-api/internal/middlewares"
	"anywhere-api/internal/repositories"
	"anywhere-api/internal/services"
	"anywhere-api/pkg/database"
	"anywhere-api/pkg/helper"
	"os"

	"github.com/gofiber/fiber/v2"
)

func main() {
	// Load .env file
	helper.LoadEnv()
	// Load the database connection (now the environment variables are loaded)
	database.DBConnect()

	// Initialize the Fiber app
	app := fiber.New()

	// Initialize user repository, service, and handler
	userRepo := repositories.NewUserRepository(database.DBConnect())
	userSvc := services.NewUserService(userRepo)
	userHandler := handlers.NewUserHandler(userSvc)

	// Define routes
	app.Post("/users", userHandler.CreateUser)
	app.Get("/users", userHandler.GetAllUsers)
	app.Get("/users/:id", userHandler.GetUserByID)
	app.Get("/users/username/:username", userHandler.GetUserByUsername)
	app.Delete("/users/:id", userHandler.DeleteUser)
	app.Post("/login", userHandler.Login)

	// Test authentication middleware with a protected route
	app.Get("/protected", middleware.Protect, func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{"message": "This route is protected"})
	})

	// Get port from environment variables, fallback to 3000 if not set
	port := os.Getenv("PORT")
	if port == "" {
		port = "3000" // Default port if not set in .env or environment
	}

	// Start the server
	app.Listen(":" + port)
}
