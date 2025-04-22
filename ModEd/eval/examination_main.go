package main

import (
	cli_examination "ModEd/eval/cli"
	migration_controller "ModEd/eval/controller"
	"fmt"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"log"
	"bufio"
	"os"
)

func clearStdin() {
    reader := bufio.NewReader(os.Stdin)
    reader.ReadString('\n') 
}

func main() {
	
	db, err := gorm.Open(sqlite.Open("database.db"), &gorm.Config{})
	if err != nil {
		log.Fatal("Connection failed:", err)
	}

	migrationController := migration_controller.NewMigrationController(db)
	err = migrationController.MigrateToDB()
	if err != nil {
		panic("err: migration failed")
	}

	fmt.Println("Starting Examination CLI...")
	for {
		fmt.Println("Select an option:")
		fmt.Println("1. Run Answer CLI")
		fmt.Println("2. Run Examination CLI")
		fmt.Println("3. Run Question CLI")
		fmt.Println("4. Run Result CLI")
		fmt.Println("5. Exit")
		var choice int
		fmt.Print("Enter your choice: ")
		fmt.Scan(&choice)
		clearStdin()

		switch choice {
		case 1:
			cli_examination.RunAnswerCLI(db)
		case 2:
			cli_examination.RunExaminationCLI(db)
		case 3:
			cli_examination.RunQuestionCLI(db)
		case 4:
			cli_examination.RunResultCLI(db)
		case 5:
			fmt.Println("Exiting...")
			return
		default:
			fmt.Println("Invalid choice. Please try again.")
		}
	}
}