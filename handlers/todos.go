package handlers

import (
	"github.com/gofiber/fiber/v2"
	"learn-fibre/database"
	"learn-fibre/models"
)

func ListTodos(c *fiber.Ctx) error {
	db := database.DB
	var todos []models.Todo
	err := db.Find(&todos).Error

	if err != nil {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{
			"success": false,
			"data":    nil,
			"message": err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"success": true,
		"data":    todos,
		"message": "Successfully retrieved todos",
	})
}

func CreateTodo(c *fiber.Ctx) error {
	db := database.DB
	var todo models.Todo
	if err := c.BodyParser(&todo); err != nil {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{
			"success": false,
			"data":    nil,
			"message": err.Error(),
		})
	}

	db.Create(&todo)

	return c.JSON(fiber.Map{
		"success": true,
		"data":    todo,
		"message": "Successfully created todo",
	})
}

func MarkTodoAsDone(c *fiber.Ctx) error {
	db := database.DB
	var todo models.Todo
	if err := c.BodyParser(&todo); err != nil {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{
			"success": false,
			"data":    nil,
			"message": err.Error(),
		})
	}

	todo.IsCompleted = true
	db.Save(&todo)

	return c.JSON(fiber.Map{
		"success": true,
		"data":    todo,
		"message": "Successfully marked as done",
	})
}

func DeleteTodo(c *fiber.Ctx) error {
	db := database.DB
	var todo models.Todo
	if err := c.BodyParser(&todo); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"data":    err,
			"message": err.Error(),
		})
	}

	db.Delete(&todo)

	return c.JSON(fiber.Map{
		"success": true,
		"data":    todo,
		"message": "Successfully deleted todo",
	})
}
