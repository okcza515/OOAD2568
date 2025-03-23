package main

import (
	controllerExamination "ModEd/eval/controller/examination"
	migration_controller "ModEd/eval/controller/migration"
	"ModEd/eval/model"
	// "errors"
	// "flag"
	"fmt"
	// "os"
	"time"
	"log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

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

	// Create an instance of the controller
	examController := controllerExamination.NewExaminationController(db)

	// CLI Menu
	for {
		fmt.Println("\nExamination CLI")
		fmt.Println("1. Create Examination")
		fmt.Println("2. Display All Examination")
		fmt.Println("3. Exit")
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
			DisplayAllExams(db)
		case 3:
			fmt.Println("Exiting...")
			return

		default:
			fmt.Println("Invalid choice")
		}
	}
}

func DisplayAllExams(db *gorm.DB) {
	var exams []model.Examination
	if err := db.Find(&exams).Error; err != nil {
		fmt.Println("Error retrieving exams:", err)
		return
	}

	fmt.Println("\nList of Examinations:")
	if len(exams) == 0 {
		fmt.Println("No examinations found.")
		return
	}

	fmt.Println("-----------------------------------------------------")
	fmt.Printf("%-5s | %-20s | %-10s | %-10s | %-10s | %-15s | %-20s | %-10s\n",
		"ID", "Exam Name", "Instructor", "Course", "Curriculum", "Criteria", "Description", "Exam Date")
	fmt.Println("-----------------------------------------------------")

	for _, exam := range exams {
		fmt.Printf("%-5d | %-20s | %-10d | %-10d | %-10d | %-15s | %-20s | %-10s\n",
			exam.ID, exam.Exam_name, exam.Instructor_id, exam.CourseId, exam.CurriculumId,
			exam.Criteria, exam.Description, exam.Exam_date.Format("2006-01-02"))
	}
}
