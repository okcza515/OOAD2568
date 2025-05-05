package main

import (
	controller "ModEd/common/controller"

	"errors"
	"flag"
	"fmt"
	"os"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type MenusChoices int

const (
	READFILE MenusChoices = iota
	REGISTER
	RETRIEVE
	DELETE
	CLEAR_DB
	EXIT
	TEST
)

func main() {
	var (
		database string
		path     string
	)
	flag.StringVar(&database, "database", "data/ModEd.bin", "Path of SQLite Database.")
	flag.StringVar(&path, "path", "", "Path to CSV or JSON.")
	flag.Parse()

	args := flag.Args()
	fmt.Printf("args: %v\n", args)

	db := ConnectDB()

	for {
		choice := Menus()
	
		switch choice {
		case READFILE:
			if _, err := os.Stat(path); errors.Is(err, os.ErrNotExist) {
				fmt.Printf("*** Error: %s does not exist.\n", path)
				return
			} else {
				fmt.Printf("*** File %s is readable\n", path)
			}
	
		case REGISTER:
			if _, err := os.Stat(path); errors.Is(err, os.ErrNotExist) {
				fmt.Printf("*** Error: %s does not exist.\n", path)
				return
			}
	
			var registerChoice int
			fmt.Print("\nSelect model:\n0. Exit\n1. Student\n2. Instuctor\n3. Department\n4. Faculty\nchoice: ")
			fmt.Scan(&registerChoice)
	
			controller.GenericRegister(registerChoice, db, path)
	
		case RETRIEVE:
			var retrieveChoice int
			fmt.Print("\nSelect model:\n0. Exit\n1. Student\n2. Instuctor\n3. Department\n4. Faculty\nchoice: ")
			fmt.Scan(&retrieveChoice)

			controller.GenericRetrieve(retrieveChoice, db)
	
		case DELETE:
			var deleteChoice int
			fmt.Print("\nSelect model:\n0. Exit\n1. Student\n2. Instuctor\n3. Department\n4. Faculty\nchoice: ")
			fmt.Scan(&deleteChoice)
	
			controller.GenericDelete(deleteChoice, db)
	
		case CLEAR_DB:
			if confirmAction("Are you sure you want to clear all tables? This action cannot be undone (y/n): ") {
				studentController := controller.CreateStudentController(db)
				instructorController := controller.CreateInstructorController(db)
				departmentController := controller.CreateDepartmentController(db)
				facultyController := controller.CreateFacultyController(db)
	
				studentController.Truncate()
				instructorController.Truncate()
				departmentController.Truncate()
				facultyController.Truncate()
	
				fmt.Println("All tables have been cleared.")
			} else {
				fmt.Println("Operation cancelled.")
			}
	
		case EXIT:
			fmt.Println("Goodbye!")
			return
	
		case TEST:
			//fmt.Println("This function has been disabled.")
			instructorController := controller.CreateInstructorController(db)
			instructorController.ManualAddInstructor()
	
		default:
			fmt.Println("Invalid choice. Please select a valid option.")
		}
	}
}


func ConnectDB() *gorm.DB {
	connector, err := gorm.Open(sqlite.Open("data/ModEd.bin"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	return connector
}

func Menus() MenusChoices {
	var choice MenusChoices
	fmt.Print("Select an option:\n0. Read file\n1. Register\n2. Retrieve\n3. Delete\n4. Clear DB\n5. Exit\nchoice: ")
	fmt.Scan(&choice)
	return choice
}

func confirmAction(prompt string) bool {
	var response string
	fmt.Print(prompt)
	fmt.Scan(&response)
	return response == "y" || response == "Y"
}