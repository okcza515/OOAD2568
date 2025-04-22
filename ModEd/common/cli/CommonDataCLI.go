package main

import (
	controller "ModEd/common/controller"
	"ModEd/common/model"
	"ModEd/utils/deserializer"

	// "ModEd/core"
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

	case REGISTER: // Not generic yet
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
		fmt.Print("\nSelect model:\n0. Student\n1. Instuctor\n2. Department\n3. Faculty\nchoice: ")
		fmt.Scan(&deleteChoice)

	case CLEAR_DB:
		studentController := controller.CreateStudentController(db)
		intructorController := controller.CreateInstructorController(db)
		departmentController := controller.CreateDepartmentController(db)
		facultyController := controller.CreateFacultyController(db)

		studentController.Truncate()
		intructorController.Truncate()
		departmentController.Truncate()
		facultyController.Truncate()

		fmt.Println("All table has been cleared.")

	case TEST:
		departmentController := controller.CreateDepartmentController(db)
		departmentController.Register([]*model.Department{{Name: "CS", Budget: 100}})
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

// **TODO when BaseController is fully functional or the truncate is no longer needed**
// A unified and reusable register function
func GenericRegister(choice int, db *gorm.DB, path string) {
	if choice == 0 {
		studentController := controller.CreateStudentController(db)
		deserializer, err := deserializer.NewFileDeserializer(path)
		if err != nil {
			panic(err)
		}
		var students []*model.Student
		if err := deserializer.Deserialize(&students); err != nil {
			panic(err)
		}
		studentController.Register(students)
		fmt.Println("Student data updated successfully.")
	} else if choice == 1 {
		instructorController := controller.CreateInstructorController(db)
		deserializer, err := deserializer.NewFileDeserializer(path)
		if err != nil {
			panic(err)
		}
		var instructors []*model.Instructor
		if err := deserializer.Deserialize(&instructors); err != nil {
			panic(err)
		}
		instructorController.Register(instructors)
		fmt.Println("Instructor data updated successfully.")
	} else if choice == 2 {
		departmentController := controller.CreateDepartmentController(db)
		deserializer, err := deserializer.NewFileDeserializer(path)
		if err != nil {
			panic(err)
		}
		var departments []*model.Department
		if err := deserializer.Deserialize(&departments); err != nil {
			panic(err)
		}
		departmentController.Register(departments)
		fmt.Println("Department data updated successfully.")
	} else if choice == 3 {
		facultyController := controller.CreateFacultyController(db)
		deserializer, err := deserializer.NewFileDeserializer(path)
		if err != nil {
			panic(err)
		}
		var faculties []*model.Faculty
		if err := deserializer.Deserialize(&faculties); err != nil {
			panic(err)
		}
		facultyController.Register(faculties)
		fmt.Println("Faculty data updated successfully.")
	}
}
