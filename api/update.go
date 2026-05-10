package api

import (
	"crud-app-tasks/db"

	"github.com/gofiber/fiber/v3"
)

func UpdateTask(c fiber.Ctx) error {
	var task Task
	id := c.Params("id")

	// Check if tasks exists
	if err := db.ConnectDB().First(&task, "id = ?", id).Error; err != nil {
		return c.Status(404).JSON(fiber.Map{"error": "Task not found"})
	}

	// Parse the request body
	if err := c.Bind().JSON(&task); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Cannot parse JSON"})
	}

	// Update the task in the database
	db.ConnectDB().Save(&task)

	return c.Status(200).JSON(task)
}
