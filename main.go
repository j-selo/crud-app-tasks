package main

import (
	"crud-app-tasks/api"
	"crud-app-tasks/db"
	"os"

	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/logger"
)

func getAddr() string {
	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}
	return ":" + port
}

func main() {
	// Create Fiber App instance
	app := fiber.New()

	// Middleware for logging HTTP requests
	app.Use(logger.New(logger.Config{
		Format: "[${time}] ${status} - ${latency} ${method} ${path}\n",
	}))

	// Connect to Supabase PostgreSQL database and migrate schema
	db.ConnectDB().AutoMigrate(&api.Task{})

	// Create Task
	app.Post("/tasks/create", api.CreateTask)

	// Read Task(s)
	app.Get("/tasks/:id", api.GetTask)
	app.Get("/tasks", api.GetTasks)

	// Update Task
	app.Put("/tasks/update/:id", api.UpdateTask)

	// Delete Task
	app.Delete("/tasks/delete/:id", api.DeleteTask)

	// Start server
	app.Listen(getAddr())

}
