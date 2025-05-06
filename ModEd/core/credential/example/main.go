package main

import (
	"fmt"
	"log"

	"ModEd/core/cli"
	"ModEd/core/credential"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {
	// Initialize database
	db, err := gorm.Open(sqlite.Open("auth.db"), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	// Create menu manager
	menuManager := cli.NewCLIMenuManager()

	// Create and add authentication menu
	authMenu := credential.NewAuthMenuState(db)
	menuManager.AddMenu("auth", authMenu)

	// Set initial state to auth menu
	err = menuManager.GoToMenu("auth")
	if err != nil {
		log.Fatalf("Failed to set initial menu state: %v", err)
	}

	// Main loop
	for {
		// Render current menu
		menuManager.Render()

		// Get user input
		var input string
		fmt.Scanln(&input)
		menuManager.UserInput = input

		// Handle user input
		err := menuManager.HandleUserInput()
		if err != nil {
			fmt.Printf("Error: %v\n", err)
		}

		// Exit if user selects option 5
		if input == "5" {
			break
		}
	}
}
