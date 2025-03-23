package main

import (
	controllerExamination "ModEd/eval/controller/examination"
	migration_controller "ModEd/eval/controller/migration"
	"ModEd/eval/model"
	"errors"
	"flag"
	"fmt"
	"os"
	"time"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {
	var (
		database string
		path     string
	)

	flag.StringVar(&database, "database", "", "Path of SQLite Database.")
	flag.StringVar(&path, "path", "", "Path to CSV or JSON for student registration.")
	flag.Parse()

	db, err := gorm.Open(sqlite.Open(database), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	if _, err := os.Stat(path); errors.Is(err, os.ErrNotExist) {
		panic("*** Error: " + path + " does not exist.\n")
	}

	
	migrationController := migration_controller.NewMigrationController(db)
	err = migrationController.MigrateToDB()
	if err != nil {
		panic("err: migration failed")
	}

	// Create an instance of the controller
	examController := controllerExamination.NewExaminationController(db)

	// CLI Menu
	for {
		fmt.Println("\nExamination CLI")
		fmt.Println("1. Create Examination")
		fmt.Println("2. Exit")
		fmt.Print("Enter your choice: ")
		var choice int
		fmt.Scan(&choice)

		switch choice {
		case 1:
			var exam model.Examination
			fmt.Print("Enter Exam Name: ")
			fmt.Scan(&exam.Exam_name)

			fmt.Print("Enter Instructor ID: ")
			fmt.Scan(&exam.Instructor_id)

			fmt.Print("Enter Course ID: ")
			fmt.Scan(&exam.CourseId)

			fmt.Print("Enter Curriculum ID: ")
			fmt.Scan(&exam.CurriculumId)

			fmt.Print("Enter Criteria: ")
			fmt.Scan(&exam.Criteria)

			fmt.Print("Enter Description: ")
			fmt.Scan(&exam.Description)

			fmt.Print("Enter Exam Date (YYYY-MM-DD): ")
			var examDateStr string
			fmt.Scan(&examDateStr)
			exam.Exam_date, _ = time.Parse("2006-01-02", examDateStr)

			// Use the controller to create the exam
			if err := examController.CreateExam(&exam); err != nil {
				fmt.Println("Error:", err)
			} else {
				fmt.Println("Examination created successfully!")
			}

		case 2:
			fmt.Println("Exiting...")
			return

		default:
			fmt.Println("Invalid choice")
		}
	}
}
