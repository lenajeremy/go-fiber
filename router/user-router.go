package router

import (
	"github.com/gofiber/fiber/v2"
	"learn-fibre/handlers"
)

func UserRouter(router *fiber.Router) {
	(*router).Post("/register", handlers.RegisterUser)
	(*router).Post("/login", handlers.LoginUser)
}
