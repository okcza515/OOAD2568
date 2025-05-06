package main

import (
	"ModEd/eval/cli"
	"ModEd/eval/controller"
	"fmt"
	"log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {
	db, err := gorm.Open(sqlite.Open("database.db"), &gorm.Config{})
	if err != nil {
		log.Fatal("Connection failed:", err)
	}

	// Run migrations if you have a migration controller
	migrationController := controller.NewMigrationController(db)
	err = migrationController.MigrateToDB()
	if err != nil {
		panic("err: migration failed")
	}

	assessmentCLI := cli.NewAssessmentCLI(db)

	fmt.Println("Starting Assessment CLI...")
	for {
		fmt.Println("\nEnter assessment command (or 'exit' to quit):")
		var input string
		fmt.Print("> ")
		// Read the whole line as input
		if _, err := fmt.Scanln(&input); err != nil {
			if err.Error() == "unexpected newline" {
				continue
			}
			fmt.Println("Error reading input:", err)
			continue
		}
		if input == "exit" {
			fmt.Println("Exiting Assessment CLI.")
			break
		}
		// Split input into args (simulate os.Args)
		args := append([]string{"assessment"}, input)
		assessmentCLI.Run(args)
	}
}
