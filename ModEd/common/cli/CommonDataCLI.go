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

	connector, err := gorm.Open(sqlite.Open(database), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	if _, err := os.Stat(path); errors.Is(err, os.ErrNotExist) {
		fmt.Printf("*** Error: %s does not exist.\n", path)
		return
	}

	studentController := controller.CreateStudentController(connector)
	instructorController := controller.CreateInstructorController(connector)
	// controller := core.NewBaseController[*model.Student](connector)

	// mapper, err := util.CreateMapper(path)
	// if err != nil {
	// 	panic(err)
	// }

	// students := mapper.Map()
	// controller.Register(students)

	deserializer, err := deserializer.NewFileDeserializer(path)
	if err != nil {
		panic(err)
	}
	var students []*model.Student
	if err := deserializer.Deserialize(&students); err != nil {
		panic(err)
	}

	studentController.Register(students)

	//InstuctorRegister(&[]model.Instructor{}, deserializer, instructorController)
	// for _, student := range students {
	// 	controller.Insert(student)
	// }

	//retrieved, err := controller.GetAll()
	// retrieved, err := controller.List(nil)

	// if err != nil {
	// 	panic(err)
	// }
	// for _, student := range retrieved {
	// 	fmt.Printf("%s\n", student.StudentCode)
	// }

	retrievedI, err := instructorController.GetAll()
	if err != nil {
		panic(err)
	}
	for _, i := range retrievedI {
		fmt.Printf("%s\n", i.FirstName)
	}

	fmt.Println("✔️ Student data updated successfully.")
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
