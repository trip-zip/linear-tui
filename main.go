package main

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/trip-zip/linear-tui/cmd"
)

func main() {
	// Load environment variables from .env file
	err := godotenv.Load()
	if err != nil {
		log.Printf("Warning: Could not load .env file: %v", err)
	}

	// Check if we have the required API key
	if os.Getenv("LINEAR_API_KEY") == "" {
		fmt.Println("Error: LINEAR_API_KEY environment variable not set")
		fmt.Println("Please create a .env file with your Linear API key:")
		fmt.Println("LINEAR_API_KEY=your_api_key_here")
		os.Exit(1)
	}

	// Execute cobra CLI
	cmd.Execute()
}
