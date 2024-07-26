package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"learn-fibre/database"
	"learn-fibre/router"
	"log"
)

func main() {

	app := fiber.New()
	app.Use(logger.New())
	app.Use(cors.New())

	// serve public files
	app.Static("", "./public")

	database.ConnectToDB()

	router.SetupRouter(app)

	log.Fatal(app.Listen(":3000"))
}
