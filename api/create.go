package api

import (
	"crud-app-tasks/db"

	"github.com/gofiber/fiber/v3"
	"gorm.io/gorm"
)

type Task struct {
	gorm.Model
	Title       string `json:"title"`
	Description string `json:"description"`
	Task        string `json:"task"`
}

func CreateTask(c fiber.Ctx) error {
	task := new(Task)

	if err := c.Bind().Body(task); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": err.Error()})
	}

	db.ConnectDB().Create(&task)

	return c.Status(201).JSON(task)
}
