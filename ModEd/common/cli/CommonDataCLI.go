package main

import (
	controller "ModEd/common/controller"
	"ModEd/common/model"
	"ModEd/common/util"
	"ModEd/utils/deserializer"

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
	DELETE
	CLEAR_DB
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
		fmt.Print("\nSelect model:\n0. Student\n1. Instuctor\n2. Department\n3. Faculty\nchoice: ")
		fmt.Scan(&registerChoice)

		GenericRegister(registerChoice, db, path)

	case DELETE:
		var deleteChoice int
		fmt.Print("\nSelect model to delete from:\n0. Student\n1. Instructor\n2. Department\n3. Faculty\nchoice: ")
		fmt.Scan(&deleteChoice)

		GenericDelete(deleteChoice, db)

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

	case TEST:
		departmentController := controller.CreateDepartmentController(db)
		departmentController.Register([]*model.Department{{Name: "CS", Budget: 100}})
		fmt.Println("Test data inserted successfully.")
	default:
		fmt.Println("Invalid choice. Please select a valid option.")
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
	fmt.Print("Select an option:\n0. Read file\n1. Register\n2. Delete\n3. Clear DB\nchoice: ")
	fmt.Scan(&choice)
	return choice
}

func GenericRegister(choice int, db *gorm.DB, path string) {
	deserializer := util.CommonDeserializer(path)

	switch choice {
	case 0:
		registerModel(db, &deserializer,
			controller.CreateStudentController(db),
			[]*model.Student{},
			"student")
	case 1:
		registerModel(db, &deserializer,
			controller.CreateInstructorController(db),
			[]*model.Instructor{},
			"instructor")
	case 2:
		registerModel(db, &deserializer,
			controller.CreateDepartmentController(db),
			[]*model.Department{},
			"department")
	case 3:
		registerModel(db, &deserializer,
			controller.CreateFacultyController(db),
			[]*model.Faculty{},
			"faculty")
	default:
		fmt.Println("Invalid choice. Please select a number between 0 and 3.")
	}
}

func registerModel[T any](db *gorm.DB, deserializer *deserializer.FileDeserializer,
	controller interface{ Register(models []*T) error },
	models []*T, modelName string) {

	// Deserialize the data
	if err := deserializer.Deserialize(&models); err != nil {
		fmt.Printf("Error deserializing %s data: %v\n", modelName, err)
		return
	}

	// Register the data with the controller
	if err := controller.Register(models); err != nil {
		fmt.Printf("Error registering %s data: %v\n", modelName, err)
		return
	}

	fmt.Printf("Complete task. (on %s)\n", modelName)
}

func GenericDelete(choice int, db *gorm.DB) {
	var id string
	fmt.Print("Enter ID to delete: ")
	fmt.Scan(&id)

	if id == "" {
		fmt.Println("Error: ID cannot be empty.")
		return
	}

	switch choice {
	case 0: // Student
		if err := model.DeleteStudentByCode(db, id); err != nil {
			fmt.Printf("Error deleting student: %v\n", err)
			return
		}
		fmt.Printf("Student with ID %s deleted successfully.\n", id)
	case 1: // Instructor
		if err := model.DeleteInstructorByCode(db, id); err != nil {
			fmt.Printf("Error deleting instructor: %v\n", err)
			return
		}
		fmt.Printf("Instructor with ID %s deleted successfully.\n", id)
	case 2: // Department
		fmt.Println("Department deletion not implemented. Use CLEAR_DB to reset all departments.")
	case 3: // Faculty
		fmt.Println("Faculty deletion not implemented. Use CLEAR_DB to reset all faculties.")
	default:
		fmt.Println("Invalid choice. Please select a number between 0 and 3.")
	}
}

func confirmAction(prompt string) bool {
	var response string
	fmt.Print(prompt)
	fmt.Scan(&response)
	return response == "y" || response == "Y"
}
