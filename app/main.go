package main

import (
	"fmt"
	"log"

	"github.com/joho/godotenv"
)

func main() {
	loadenv()

	fmt.Println("hello world")
}

func loadenv() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}
