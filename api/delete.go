package api

import (
	"crud-app-tasks/db"

	"github.com/gofiber/fiber/v3"
)

func DeleteTask(c fiber.Ctx) error {
	var task Task
	id := c.Params("id")

	// Check if tasks exists
	if err := db.ConnectDB().First(&task, "id = ?", id).Error; err != nil {
		return c.Status(404).JSON(fiber.Map{"error": "Task not found"})
	}

	// Delete the task from the database
	db.ConnectDB().Delete(&task)

	return c.Status(200).JSON(task)
}
