package api

import (
	"crud-app-tasks/db"

	"github.com/gofiber/fiber/v3"
)

func UpdateTask(c fiber.Ctx) error {
	var task Task

	id := c.Params("id")
	db.ConnectDB().First(&task, "id = ?", id)

	if task.ID == 0 {
		return c.Status(404).JSON(fiber.Map{"error": "Task not found"})
	}

	db.ConnectDB().Save(&task)

	return c.JSON(task)
}
