package api

import (
	"crud-app-tasks/db"

	"github.com/gofiber/fiber/v3"
)

func GetTask(c fiber.Ctx) error {
	var task Task

	id := c.Params("id")
	db.ConnectDB().First(&task, "id = ?", id)

	if task.ID == 0 {
		return c.Status(404).JSON(fiber.Map{"error": "Task not found"})
	}

	return c.JSON(task)
}

func GetTasks(c fiber.Ctx) error {
	var tasks []Task

	db.ConnectDB().Find(&tasks)

	if len(tasks) == 0 {
		return c.Status(404).JSON(fiber.Map{"error": "No tasks found"})
	}

	return c.JSON(tasks)
}
