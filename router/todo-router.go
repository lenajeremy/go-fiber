package router

import (
	"github.com/gofiber/fiber/v2"
	"learn-fibre/handlers"
)

func TodosRouter(group *fiber.Router) {
	(*group).Get("/", handlers.ListTodos)
	(*group).Post("/create", handlers.CreateTodo)
	(*group).Patch("/:id/update", handlers.MarkTodoAsDone)
	(*group).Delete("/:id/delete", handlers.DeleteTodo)
}
