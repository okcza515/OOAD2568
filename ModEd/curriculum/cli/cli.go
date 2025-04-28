package main

import (
	"ModEd/core"
	"ModEd/core/migration"
	internship "ModEd/curriculum/cli/Internship"
	"ModEd/curriculum/cli/curriculum"
	instructorWorkload "ModEd/curriculum/cli/instructor_workload"
	wilproject "ModEd/curriculum/cli/wil-project"
	controller "ModEd/curriculum/controller"
	"fmt"
)

const (
	defaultDBPath = "../../data/curriculum.db"
)

// TODO: not sure is this a good approach to do at all might need to discuss further
func main() {

	db, err := migration.
		GetInstance().
		SetPathDB(defaultDBPath).
		MigrateModule(core.MODULE_CURRICULUM).
		MigrateModule(core.MODULE_INSTRUCTOR).
		MigrateModule(core.MODULE_INTERNSHIP).
		MigrateModule(core.MODULE_WILPROJECT).
		BuildDB()

	if err != nil {
		panic(err)
	}

	curriculumController := controller.NewCurriculumController(db)
	classController := controller.NewClassController(db)
	courseController := controller.NewCourseController(db)

	for {
		displayMainMenu()
		choice := getUserChoice()

		switch choice {
		case "2":
			curriculum.RunCurriculumModuleCLI(db, courseController, classController, curriculumController)
		case "3":
			wilproject.RunWILModuleCLI(db, courseController, classController)
		case "4":
			instructorWorkload.RunInstructorWorkloadModuleCLI(db, courseController, classController, curriculumController)
		case "5":
			internship.RunInterShipCLI(db)
		case "resetdb":
			err := resetDB()
			if err != nil {
				panic(err)
			}
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
	fmt.Println("1. Curriculum")
	fmt.Println("2. WIL-Project")
	fmt.Println("3. Instructor Workload")
	fmt.Println("4. Internship")
	fmt.Println("'resetdb' to re-initialize database")
}

func getUserChoice() string {
	var choice string
	fmt.Print("Enter your choice: ")
	fmt.Scanln(&choice)
	return choice
}

func resetDB() error {
	err := migration.GetInstance().DropAllTables()
	if err != nil {
		return err
	}

	_, err = migration.GetInstance().
		SetPathDB(defaultDBPath).
		MigrateModule(core.MODULE_CURRICULUM).
		MigrateModule(core.MODULE_INSTRUCTOR).
		MigrateModule(core.MODULE_INTERNSHIP).
		MigrateModule(core.MODULE_WILPROJECT).
		BuildDB()

	if err != nil {
		return err
	}

	return nil
}
