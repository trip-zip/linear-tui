package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/joho/godotenv"
)

func main() {
	// Load environment variables from .env file
	// Try multiple locations: current dir, parent dirs, etc.
	envPaths := []string{
		".env",
		"../.env", 
		"../../.env",
		filepath.Join(os.Getenv("HOME"), ".config", "linear", ".env"),
	}
	
	var err error
	for _, path := range envPaths {
		err = godotenv.Load(path)
		if err == nil {
			break
		}
	}
	if err != nil {
		log.Printf("Warning: Could not load .env file from any location: %v", err)
	}

	// Check if we have the required API key
	if os.Getenv("LINEAR_API_KEY") == "" {
		fmt.Println("Error: LINEAR_API_KEY environment variable not set")
		fmt.Println("Please create a .env file with your Linear API key:")
		fmt.Println("LINEAR_API_KEY=your_api_key_here")
		os.Exit(1)
	}

	// Initialize and execute CLI
	initCLI()
	executeCLI()
}