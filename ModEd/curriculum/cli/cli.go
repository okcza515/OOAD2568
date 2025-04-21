package main

import (
	internship "ModEd/curriculum/cli/Internship"
	"ModEd/curriculum/cli/curriculum"
	instructorWorkload "ModEd/curriculum/cli/instructor_workload"
	migrationcli "ModEd/curriculum/cli/migration"
	wilproject "ModEd/curriculum/cli/wil-project"
	controller "ModEd/curriculum/controller/curriculum"
	"ModEd/curriculum/controller/migration"
	"ModEd/curriculum/utils"
	"fmt"

	"gorm.io/gorm"
)

const (
	defaultDBPath = "../../data/curriculum.db"
)

// TODO: not sure is this a good approach to do at all might need to discuss further
func main() {

	database := utils.GetInputDatabasePath(defaultDBPath)

	db, err := utils.NewGormSqlite(&utils.GormConfig{
		DBPath: database,
		Config: &gorm.Config{},
	})
	if err != nil {
		panic(err)
	}

	curriculumController := controller.NewCurriculumController(db)
	classController := controller.NewClassController(db)
	courseController := controller.NewCourseController(db)
	migrationController := migration.NewMigrationController(db)

	for {
		displayMainMenu()
		choice := getUserChoice()

		switch choice {
		case "1":
			migrationcli.RunMigrationCLI(migrationController)
		case "2":
			curriculum.RunCurriculumCLI(curriculumController)
		case "3":
			curriculum.RunClassCLI(classController)
		case "4":
			curriculum.RunCourseCLI(courseController)
		case "5":
			wilproject.RunWILModuleCLI(db, courseController, classController)
		case "6":
			instructorWorkload.RunInstructorWorkloadModuleCLI(db, courseController, classController, curriculumController)
		case "7":
			internship.RunInterShipCLI(db)
		case "0":
			fmt.Println("Exiting...")
			return
		default:
			fmt.Println("Invalid option")
		}
	}
}

func displayMainMenu() {
	fmt.Println("\nModEd CLI Menu")
	fmt.Println("1. Migration")
	fmt.Println("2. Curriculum")
	fmt.Println("3. Class")
	fmt.Println("4. Course")
	fmt.Println("5. WIL-Project")
	fmt.Println("6. Instructor Workload")
	fmt.Println("7. Internship")
	fmt.Println("0. Exit")
}

func getUserChoice() string {
	var choice string
	fmt.Print("Enter your choice: ")
	fmt.Scanln(&choice)
	return choice
}
