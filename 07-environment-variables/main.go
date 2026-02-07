package main

import (
	"os"

	"github.com/joho/godotenv"
)

func main() {
	// Access environment variables
	data := os.Getenv("PATH")
	println(data[:100])

	// Load .env file
	godotenv.Load(".env.example")

	// Access .env variable
	data2 := os.Getenv("SECRET_DATA")
	println(data2)

}
