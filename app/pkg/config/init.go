package config

import (
	"log"

	"github.com/joho/godotenv"
)

func init() {
	loadEnv()
	if !GetEnv().IsTest() {
		err := godotenv.Load()
		if err != nil {
			log.Fatal("Error loading .env file")
		}
	}
	loadGCP()
}
