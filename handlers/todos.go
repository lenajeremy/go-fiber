package handlers

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"learn-fibre/database"
	"learn-fibre/models"
)

func ListTodos(c *fiber.Ctx) error {
	db := database.DB
	var todos []models.Todo

	err := db.Find(&todos).Error

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"success": false,
			"error":   err,
			"data":    todos,
		})
	}

	user, err := getUser(c)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"success": false,
			"error":   err,
			"data":    nil,
		})
	}

	err = db.Where("user_id = ?", user.ID).Find(&todos).Error

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
	user, err := getUser(c)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"success": false,
			"error":   err,
			"data":    nil,
		})
	}

	todo := models.Todo{
		UserID: user.ID.String(),
	}

	if err := c.BodyParser(&todo); err != nil {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{
			"success": false,
			"data":    nil,
			"message": err.Error(),
		})
	}

	err = db.Create(&todo).Error

	if err != nil {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{
			"success": false,
			"message": err.Error(),
			"data":    nil,
		})
	}

	return c.JSON(fiber.Map{
		"success": true,
		"data":    todo,
		"message": "Successfully created todo",
	})
}

func MarkTodoAsDone(c *fiber.Ctx) error {
	db := database.DB
	var todo models.Todo

	todoId := c.Params("id", "")

	if err := db.Where("id = ?", todoId).First(&todo).Error; err != nil {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{
			"success": false,
			"error":   fmt.Sprintf("Todo with id `%s` does not exist", todoId),
			"data":    nil,
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
	todoId := c.Params("id", "")

	if todoId == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "Invalid todo id",
			"data":    nil,
		})
	}

	err := db.Delete(&todo, "id = ?", todoId).Error

	if err != nil {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{
			"success": false,
			"message": err.Error(),
			"data":    nil,
		})
	}

	return c.JSON(fiber.Map{
		"success": true,
		"data":    nil,
		"message": "Successfully deleted todo",
	})
}

func getUser(c *fiber.Ctx) (user models.User, err error) {
	token := c.Locals("user").(*jwt.Token)
	claims := token.Claims.(jwt.MapClaims)
	uid, _ := uuid.Parse(claims["id"].(string))

	err = database.DB.First(&user, uid).Error
	return user, err
}
