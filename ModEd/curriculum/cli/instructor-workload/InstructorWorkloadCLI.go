package main

import (
	"ModEd/curriculum/cli/instructor-workload/handler"
	controller "ModEd/curriculum/controller/migration"
	"ModEd/curriculum/utils"
	"fmt"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {
	var (
		database string
		// path     string
	)

	database = "data/ModEd.bin"
	connector, err := gorm.Open(sqlite.Open(database), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	migrationController := controller.NewMigrationController(connector)
	err = migrationController.MigrateToDB()
	if err != nil {
		panic("err: migration failed")
	}

	input := ""
	for input != "exit" {
		printOptions()
		choice := utils.GetUserChoice()
		fmt.Println("choice: ", choice)
		switch choice {
		case "1": // Teaching Responsibility
			handler.RunAcademicWorkloadHandler()
		case "2": // Administrative Tasks
			handler.RunAdminstrativeWorkloadHandler()
		case "3": // Senior Projects
			handler.RunSeniorProjectWorkloadHandler()
		case "exit":
			fmt.Println("Exiting...")
			return
		default:
			fmt.Println("Invalid option. Please try again.")
		}
	}
}

func printOptions() {
	fmt.Println("\nInstructor Workload Module Menu:")
	fmt.Println("1. Academic")
	fmt.Println("2. Administrative Task")
	fmt.Println("3. Senior Project")
	fmt.Println("Type 'exit' to quit")
}
