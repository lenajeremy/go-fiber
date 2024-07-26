package database

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"learn-fibre/config"
	"learn-fibre/models"
	"log"
)

func ConnectToDB() {
	config.LoadConfig()

	dbPass, dbUser, dbName, dbPort, dbHost :=
		config.GetEnv("DB_PASS", ""),
		config.GetEnv("DB_USER", ""),
		config.GetEnv("DB_NAME", ""),
		config.GetEnv("DB_PORT", ""),
		config.GetEnv("DB_HOST", "")

	var err error
	dsn := fmt.Sprintf("host=%s user=%s port=%s password=%s dbname=%s sslmode=disable", dbHost, dbUser, dbPort, dbPass, dbName)
	pg := postgres.Open(dsn)
	DB, err = gorm.Open(pg, &gorm.Config{})

	if err != nil {
		log.Fatal(err)
	}

	log.Println("Connection to DB established... Making migrations")
	DB.AutoMigrate(&models.User{}, &models.Profile{}, &models.Todo{})
	log.Println("Migrations complete")
}
