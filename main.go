// app.go
package main

import (
	"fmt"

	"github.com/Aslamdesusa/click-tracker/app/http"
	"github.com/joho/godotenv"
)

func main() {
	// Load environment variables from .env file
	if err := godotenv.Load(); err != nil {
		fmt.Println("Error loading .env file")
	}

	// Initialize your application
	http.Init("development")
}
