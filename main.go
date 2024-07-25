package main

import (
	"encoding/json"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"log"
)

type Mails struct {
	From    string `json:"from"`
	Body    string `json:"body"`
	Subject string `json:"subject"`
	To      string `json:"to"`
}

func main() {
	app := fiber.New()
	app.Use(logger.New())
	app.Use(cors.New())

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	app.Get("/ping", func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"message":     "o=ping",
			"allowed":     "true",
			"personality": "extrovert",
		})
	})

	app.Get("/mails", func(c *fiber.Ctx) error {
		mails := []Mails{
			{"jeremiahlena13@gmail.com", "How are you doing sir?", "Saying Hello", "winnerokere@gmail.com"},
			{"marvelouslena13@gmail.com", "I'm doing well, there's something that I've been meaning to tell you", "Story time", "jeremiahlena13@gmail.com"},
		}

		bytes, err := json.Marshal(mails)

		if err != nil {
			return fiber.NewError(fiber.StatusInternalServerError, err.Error())
		}

		mailsArray := make([]map[string]interface{}, 0)

		err = json.Unmarshal(bytes, &mailsArray)

		return c.JSON(mailsArray)
	})

	app.Static("", "./public")

	log.Fatal(app.Listen(":3000"))
}
