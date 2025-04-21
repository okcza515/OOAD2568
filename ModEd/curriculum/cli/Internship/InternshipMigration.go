package Internship

import (
	controller "ModEd/curriculum/controller/Internship"
	"ModEd/curriculum/utils"
	"fmt"
)

func InternshipMigration(migrationController *controller.MigrationController) {
	for {
		SelectOption()
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

func SelectOption() {
	fmt.Println("\nSelect Option:")
	fmt.Println("1. Migrate tables")
	fmt.Println("2. Drop all tables")
	fmt.Println("0. Exit")
}
