package config

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

func LoadConfig() {
	log.Println("loading config")
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func GetEnv(key, fallback string) string {
	v, ok := os.LookupEnv(key)

	if ok {
		return v
	} else {
		return fallback
	}
}
