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

	connector := ConnectDB()

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

		studentController := controller.CreateStudentController(connector)
		deserializer, err := deserializer.NewFileDeserializer(path)
		if err != nil {
			panic(err)
		}
		//var students []*model.Student
		var students *model.Student
		if err := deserializer.Deserialize(&students); err != nil {
			panic(err)
		}

		studentController.Create(students)
		fmt.Println("Student data updated successfully.")

	case DELETE:
		
	case CLEAR_DB: 
		studentController := controller.CreateStudentController(connector)
		studentController.TruncateStudents()

		fmt.Println("Student table has been cleared.")
	}

	
	instructorController := controller.CreateInstructorController(connector)
	// controller := core.NewBaseController[*model.Student](connector)

	//InstuctorRegister(&[]model.Instructor{}, deserializer, instructorController)
	// for _, student := range students {
	// 	controller.Insert(student)
	// }

	// retrieved, err := controller.List(nil)

	retrievedI, err := instructorController.GetAll()
	if err != nil {
		panic(err)
	}
	for _, i := range retrievedI {
		fmt.Printf("%s\n", i.FirstName)
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
func GenericRegister[T any]() {

}

func InstuctorRegister(instructors *[]model.Instructor, d *deserializer.FileDeserializer, iController *controller.InstructorController) {
	if err := d.Deserialize(&instructors); err != nil {
		panic(err)
	}

	iController.Register(instructors)
}
