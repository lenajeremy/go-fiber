package router

import (
	"encoding/json"
	"github.com/gofiber/fiber/v2"
	"learn-fibre/middleware"
)

type Mails struct {
	From    string `json:"from"`
	Body    string `json:"body"`
	Subject string `json:"subject"`
	To      string `json:"to"`
}

func SetupRouter(app *fiber.App) {
	todoRouterGroup := app.Group("/todos", middleware.Protected())
	TodosRouter(&todoRouterGroup)

	authGroup := app.Group("/auth")
	UserRouter(&authGroup)

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	app.Get("/ping", middleware.Protected(), func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"message":     "o=ping",
			"allowed":     "true",
			"personality": "extrovert",
			"fromLocal":   c.IsFromLocal(),
			"ipAddress":   c.IP(),
			"baseUrl":     c.BaseURL(),
		})
	})

	app.Get("/not-found", func(c *fiber.Ctx) error {
		return fiber.ErrNotFound
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

}
