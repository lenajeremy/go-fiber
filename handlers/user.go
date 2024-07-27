package handlers

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"learn-fibre/database"
	"learn-fibre/models"
	"learn-fibre/utils"
	"log"
	"strings"
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
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"data":    nil,
			"success": false,
			"message": "Username or Password is empty",
		})
	}

	log.Println(user.Password, loginValues.Password)
	if ok := checkPasswordMatch(user.Password, loginValues.Password); ok {
		// the password matches, log the user in
		log.Println(err)
		log.Println("User Login Success")
		return ctx.JSON(fiber.Map{
			"data":    user,
			"success": true,
			"message": "Logged in successfully",
		})
	} else {
		return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
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

	if registerValues.Username == "" ||
		registerValues.Password == "" ||
		registerValues.FirstName == "" ||
		registerValues.LastName == "" {

		var missingValues []string

		if registerValues.Username == "" {
			missingValues = append(missingValues, "username")
		}
		if registerValues.Password == "" {
			missingValues = append(missingValues, "password")
		}
		if registerValues.FirstName == "" {
			missingValues = append(missingValues, "first name")
		}
		if registerValues.LastName == "" {
			missingValues = append(missingValues, "last name")
		}

		return ctx.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{
			"data":    nil,
			"success": false,
			"message": getMessageFromMissingFields(missingValues),
			"testing": false,
		})
	}

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

	var user models.User

	err = db.Transaction(func(tx *gorm.DB) error {
		userProfile := models.Profile{}

		user = models.User{
			Username:  registerValues.Username,
			Password:  hashedPassword,
			FirstName: registerValues.FirstName,
			LastName:  registerValues.LastName,
			Profile:   userProfile,
		}

		//if err := tx.Create(&userProfile).Error; err != nil {
		//	return err
		//}

		if err := tx.Create(&user).Error; err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		log.Println(err.Error())
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

func checkPasswordMatch(hash string, password string) bool {
	return bcrypt.CompareHashAndPassword([]byte(hash), []byte(password)) == nil
}

func hashPassword(password string) (string, error) {
	b, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(b), err
}

func getMessageFromMissingFields(missingFields []string) string {
	if len(missingFields) == 1 {
		return fmt.Sprintf("%s is missing", utils.Capitalize(missingFields[0]))
	} else {
		return fmt.Sprintf("%s and %s are missing",
			utils.Capitalize(strings.Join(missingFields[:len(missingFields)-1], ", ")),
			missingFields[len(missingFields)-1],
		)
	}
}
