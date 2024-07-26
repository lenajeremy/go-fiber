package handlers

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
	"learn-fibre/database"
	"learn-fibre/models"
	"log"
)

type RegisterFormValues struct {
	Username  string `json:"username"`
	Password  string `json:"password"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
}

type LoginFormValues struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func LoginUser(ctx *fiber.Ctx) error {
	var loginValues LoginFormValues

	if err := ctx.BodyParser(&loginValues); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"data":    nil,
			"success": false,
			"message": err.Error(),
		})
	}

	db := database.DB
	user := models.User{Username: loginValues.Username}

	err := db.First(&user).Error
	if err != nil {
		log.Println(err)
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"data":    nil,
			"success": false,
			"message": "User Login failed. Please reach out to support",
		})
	}

	if user.Password == "" || user.Username == "" {
		return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
			"data":    nil,
			"success": false,
			"message": "Username or Password is empty",
		})
	}

	if ok := checkPasswordMatch(user.Password, loginValues.Password); ok {
		// the password matches, log the user in
		log.Println("User Login Success")
		return ctx.JSON(fiber.Map{
			"data":    user,
			"success": true,
			"message": "Logged in successfully",
		})
	} else {
		return ctx.JSON(fiber.Map{
			"data":    nil,
			"success": false,
			"message": "Invalid login credentials",
		})
	}
}

func RegisterUser(ctx *fiber.Ctx) error {

	var registerValues RegisterFormValues
	if err := ctx.BodyParser(&registerValues); err != nil {
		return ctx.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{
			"data":    nil,
			"success": false,
			"message": err.Error(),
		})
	}

	if registerValues.Username == "" || registerValues.Password == "" {
		return ctx.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{
			"data":    nil,
			"success": false,
			"message": "Username and Password are required",
		})
	}

	fmt.Println(registerValues)

	db := database.DB

	hashedPassword, err := hashPassword(registerValues.Password)

	if err != nil {
		log.Println(err.Error())
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"data":    nil,
			"success": false,
			"message": "Failed to register user",
		})
	}

	user := models.User{
		Username:  registerValues.Username,
		Password:  hashedPassword,
		FirstName: registerValues.FirstName,
		LastName:  registerValues.LastName,
		Profile:   models.Profile{},
	}

	err = db.Create(&user).Error

	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"data":    nil,
			"success": false,
			"message": err.Error(),
		})
	}

	return ctx.JSON(fiber.Map{
		"data":    user,
		"success": true,
		"message": "User registered successfully",
	})
}

func checkPasswordMatch(password string, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func hashPassword(password string) (string, error) {
	b, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(b), err
}
