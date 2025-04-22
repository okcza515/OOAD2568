package migration

import (
	controller "ModEd/curriculum/controller"
	"ModEd/curriculum/utils"
	"fmt"
)

func RunMigrationCLI(migrationController *controller.MigrationController) {
	for {
		printCurriculumMenu()
		choice := utils.GetUserChoice()

		switch choice {
		case "1":
			if err := migrationController.MigrateToDB(); err != nil {
				fmt.Println("Error migrating to DB:", err)
			} else {
				fmt.Println("Migration completed successfully.")
			}
			return
		case "2":
			if err := migrationController.DropAllTables(); err != nil {
				fmt.Println("Error dropping tables:", err)
			} else {
				fmt.Println("All tables dropped successfully.")
			}
		default:
			fmt.Println("Invalid option")
		}
	}
}

func printCurriculumMenu() {
	fmt.Println("\nCurriculum Menu:")
	fmt.Println("1. Migrate tables")
	fmt.Println("2. Drop all tables")
	fmt.Println("0. Exit")
}
